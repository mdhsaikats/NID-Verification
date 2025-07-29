-- SQL script to create a table for Bangladeshi NID information
CREATE TABLE bangladesh_nid (
    nid VARCHAR(20) PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    date_of_birth DATE NOT NULL,
    father_name VARCHAR(100) NOT NULL,
    mother_name VARCHAR(100) NOT NULL,
    address VARCHAR(255) NOT NULL,
    gender VARCHAR(10) NOT NULL
);

-- Example insert statements
INSERT INTO bangladesh_nid (nid, name, date_of_birth, father_name, mother_name, address, gender) VALUES
('1234567890', 'Rahim Uddin', '1990-01-15', 'Karim Uddin', 'Ayesha Begum', 'Dhaka, Bangladesh', 'Male'),
('9876543210123', 'Fatema Khatun', '1985-05-20', 'Jalal Uddin', 'Rokeya Begum', 'Chittagong, Bangladesh', 'Female'),
('20202020202020202', 'Sabbir Hossain', '2000-12-01', 'Habib Hossain', 'Salma Akter', 'Khulna, Bangladesh', 'Male');
