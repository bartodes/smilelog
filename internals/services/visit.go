package services

import (
	"database/sql"

	. "github.com/bartodes/smilelog/internals/models"
)

func CreateVisit(appointmentId int64, notes string, db *sql.DB) (Visit, error) {
	query := `INSERT INTO visits (appointment_id, notes) VALUES(?,?) RETURNING *;`

	var v Visit

	err := db.QueryRow(query, appointmentId, notes).Scan(
		&v.ID,
		&v.AppointmentId,
		&v.Notes,
	)

	if err != nil {
		return Visit{}, err
	}

	return v, nil
}

func GetVisit(appointmentId int64, db *sql.DB) (Visit, error) {
	query := `SELECT id, appointment_id, notes FROM visits WHERE appointment_id = ?;`

	var v Visit
	err := db.QueryRow(query, appointmentId).Scan(
		&v.ID,
		&v.AppointmentId,
		&v.Notes,
	)

	if err != nil {
		return Visit{}, err
	}

	return v, nil
}

/*
Lists all visits from a patient
*/
func ListVisits(patientId int64, db *sql.DB) ([]Visit, error) {
	query := `SELECT 
		v.id,
		v.appointment_id,
			v.notes
		FROM visits v
		INNER JOIN appointments a 
		ON v.appointment_id = a.id
		WHERE a.patient_id = ?
	;`

	rows, err := db.Query(query, patientId)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var visits []Visit

	for rows.Next() {
		var v Visit
		rows.Scan(
			&v.ID,
			&v.AppointmentId,
			&v.Notes,
		)

		visits = append(visits, v)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return visits, nil
}
