package models

import (
	"database/sql"
	"fmt"
)

type Patient struct {
	Id          uint
	DocumentId  uint
	Name        string
	LastName    string
	Email       string
	PhoneNumber uint
}

func CreatePatient(p Patient, db *sql.DB) (Patient, error) {

	query := `
	INSERT INTO patients (document_id, name, last_name, email, phone_number) 
	VALUES (?,?,?,?,?)
	RETURNING id;`

	res, err := db.Exec(
		query,
		p.DocumentId,
		p.Name,
		p.LastName,
		p.Email,
		p.PhoneNumber,
	)
	if err != nil {
		return p, fmt.Errorf("insert error: %w", err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return p, fmt.Errorf("result error: %w", err)
	}

	p.Id = uint(id)

	return p, nil
}

func (p Patient) GetPatientAppointments() {}
