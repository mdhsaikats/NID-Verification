# Bangladesh NID Verification System

This project is a simple web-based National ID (NID) verification system for Bangladesh. It allows users to enter a Bangladeshi NID number and view the associated information if it exists in the database.

## Features
- Frontend form for NID input and result display (HTML, Tailwind CSS, JavaScript)
- Backend API built with Go (Chi router, MySQL)
- MySQL database for storing NID records
- Example NID data in JSON and SQL formats
- CORS enabled for local development

## Project Structure
```
index.html           # Main frontend page
script.js            # Frontend logic for form and API calls
style.css            # (Optional) Custom styles
main.go              # Go backend server (API)
nid-database.json    # Example NID data in JSON
nid-database.sql     # SQL script to create and populate the database
```

## How It Works
1. User enters a NID number in the web form.
2. The frontend sends a request to the Go backend API (`/verify-nid`).
3. The backend checks the MySQL database for the NID and returns the details as JSON.
4. The frontend displays the NID information or an error message.

## Setup Instructions

### 1. Database Setup
- Install MySQL and create a database named `nid_info`.
- Run the SQL script `nid-database.sql` to create the `bangladesh_nid` table and insert sample data:
  ```sql
  SOURCE nid-database.sql;
  ```
- Update the database credentials in `main.go` if needed.

### 2. Backend Setup (Go)
- Install Go and required packages:
  ```sh
  go get github.com/go-chi/chi/v5
  go get github.com/go-sql-driver/mysql
  ```
- Run the Go server:
  ```sh
  go run main.go
  ```
- The API will be available at `http://localhost:8080/verify-nid`.

### 3. Frontend Setup
- Open `index.html` in your browser.
- Enter a NID number and click Verify.

## API Endpoint
- `GET /verify-nid?nid=YOUR_NID`
  - Returns NID details as JSON if found, or an error if not found.

## Example NID Numbers
- 1234567890
- 9876543210123
- 20202020202020202

## License
This project is for educational/demo purposes only.
