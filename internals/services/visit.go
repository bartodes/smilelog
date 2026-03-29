package services

import "database/sql"

type Visit struct {
	ID            int64
	AppointmentId int64
	Notes         string
}

func CreateVisit(v Visit, db *sql.DB) (Visit, error) {
	query := `INSERT INTO visits (appointment_id, notes)
	VALUES(?,?)
	RETURNING id;
	`

	err := db.QueryRow(
		query,
		v.AppointmentId,
		v.Notes,
	).Scan(&v.ID)

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

func ListVisits(db *sql.DB) ([]Visit, error) {
	var visits []Visit
	return visits, nil
}
