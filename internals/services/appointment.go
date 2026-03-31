package services

import (
	"database/sql"
	"time"
)

type Appointment struct {
	ID              int64
	PatientID       int64
	Status          string
	ScheduledFor    time.Time // ISO8601 datetime
	DurationMinutes int
}

func CreateAppointment(a Appointment, patientId int64, db *sql.DB) (Appointment, error) {
	query := `INSERT INTO appointments (patient_id, status, scheduled_for, durations_minutes)
	VALUES(?,?,?,?)
	RETURNING id;`

	err := db.QueryRow(
		query,
		a.PatientID,
		a.Status,
		a.ScheduledFor,
		a.DurationMinutes,
	).Scan(&a.ID)

	if err != nil {
		return Appointment{}, err
	}

	return a, nil
}

// Implement with optional parameter filter by schedule time (making use of the composite index)
func GetAppointment(id int64, db *sql.DB) (Appointment, error) {
	query := `SELECT id, patient_id, status, scheduled_for, duration_minutes FROM appointments WHERE id = ?;`

	var a Appointment
	err := db.QueryRow(query, id).Scan(
		&a.ID,
		&a.PatientID,
		&a.Status,
		&a.ScheduledFor,
		&a.DurationMinutes,
	)

	if err != nil {
		return Appointment{}, err
	}

	return a, nil
}

// status -> 'CREATED','CONFIRMED','CANCELED','NO_SHOW'
// Create new idx!
func ListAppointmentsByStatus(status string, db *sql.DB) ([]Appointment, error) {
	query := `SELECT id, patient_id, status, scheduled_for, duration_minutes FROM appointments WHERE status = ?;`

	rows, err := db.Query(query, status)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var appointments []Appointment

	for rows.Next() {
		var a Appointment

		rows.Scan(
			&a.ID,
			&a.PatientID,
			&a.Status,
			&a.ScheduledFor,
			&a.DurationMinutes,
		)

		appointments = append(appointments, a)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return appointments, nil
}

// schedule
// make default range 1 month: if only start is provided the rage will be = start + 1 month in time.Time
func ListAppointmentsByScheduleRange(start time.Time, end time.Time, db *sql.DB) ([]Appointment, error) {
	var appointments []Appointment
	return appointments, nil
}
