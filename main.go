package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func verifyNIDHandler(w http.ResponseWriter, r *http.Request) {
	// Allow CORS for local development
	w.Header().Set("Access-Control-Allow-Origin", "*")
	nid := r.URL.Query().Get("nid")
	if nid == "" {
		http.Error(w, "NID number is required", http.StatusBadRequest)
		return
	}

	var name, dateOfBirth, fatherName, motherName, address, gender string
	err := db.QueryRow("SELECT name, date_of_birth, father_name, mother_name, address, gender FROM bangladesh_nid WHERE nid = ?", nid).Scan(&name, &dateOfBirth, &fatherName, &motherName, &address, &gender)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "NID not found", http.StatusNotFound)
		} else {
			fmt.Println("Database error:", err)
			http.Error(w, "Database error: "+err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{"nid": "%s", "name": "%s", "date_of_birth": "%s", "father_name": "%s", "mother_name": "%s", "address": "%s", "gender": "%s"}`,
		nid, name, dateOfBirth, fatherName, motherName, address, gender)
}

func main() {
	var err error
	db, err = sql.Open("mysql", "root:29112003@tcp(localhost:3306)/nid_info")
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		return
	}
	defer db.Close()

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/verify-nid", verifyNIDHandler)
	http.ListenAndServe(":8080", r)
}
