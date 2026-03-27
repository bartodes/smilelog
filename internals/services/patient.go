package services

import (
	"database/sql"
)

type Patient struct {
	ID          int64
	Name        string
	LastName    string
	Email       string
	PhoneNumber uint
}

func CreatePatient(p Patient, db *sql.DB) (Patient, error) {
	query := `
	INSERT INTO patients (name, last_name, email, phone_number) 
	VALUES (?,?,?,?)
	RETURNING id;`

	err := db.QueryRow(
		query,
		p.Name,
		p.LastName,
		p.Email,
		p.PhoneNumber,
	).Scan(&p.ID)

	if err != nil {
		return Patient{}, err
	}

	return p, nil
}

func GetPatient(id int64, db *sql.DB) (Patient, error) {
	query := `SELECT id, name, last_name, email, phone_number FROM patients WHERE id = ?`

	var p Patient

	err := db.QueryRow(query, id).Scan(
		&p.ID,
		&p.Name,
		&p.LastName,
		&p.Email,
		&p.PhoneNumber,
	)

	if err != nil {
		return Patient{}, err
	}

	return p, nil
}
