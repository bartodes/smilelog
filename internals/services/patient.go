package services

import (
	"database/sql"

	. "github.com/bartodes/smilelog/internals/models"
)

func CreatePatient(p Patient, db *sql.DB) (Patient, error) {
	query := `INSERT INTO patients (name, last_name, email, phone_number) 
	VALUES (
		?,
		?,
		NULLIF(?,''),
		NULLIF(?,0)
	)
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

// maybe return only error and scan only to map error if patient doesnt exists
func GetPatient(id int64, db *sql.DB) (Patient, error) {
	query := `SELECT id, name, last_name, email, IFNULL(phone_number,0) FROM patients WHERE id = ?;`

	var p Patient
	err := db.QueryRow(query, id).Scan(
		&p.ID,
		&p.Name,
		&p.LastName,
		&p.Email,
		&p.PhoneNumber,
	)

	if err == sql.ErrNoRows {
		return Patient{}, ErrPatientNotFound
	}

	return p, nil
}

func ListPatients(db *sql.DB) ([]Patient, error) {
	query := `SELECT id, name, last_name, email, IFNULL(phone_number,0) FROM patients;`

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var patients []Patient
	for rows.Next() {
		var p Patient

		err := rows.Scan(
			&p.ID,
			&p.Name,
			&p.LastName,
			&p.Email,
			&p.PhoneNumber,
		)

		if err != nil {
			return nil, err
		}

		patients = append(patients, p)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return patients, nil
}
