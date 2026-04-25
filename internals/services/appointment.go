package services

import (
	"database/sql"
	"time"

	. "github.com/bartodes/smilelog/internals/models"
)

func CreateAppointment(a Appointment, db *sql.DB) (Appointment, error) {
	query := `INSERT INTO appointments (patient_id, status, scheduled_for, duration_minutes)
		VALUES(?,?,?,?)
		RETURNING id
	;`

	err := db.QueryRow(
		query,
		a.PatientID,
		CREATED,
		a.ScheduledFor,
		a.DurationMinutes,
	).Scan(&a.ID)

	if err != nil {
		return Appointment{}, err
	}

	return a, nil
}

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

	if err == sql.ErrNoRows {
		return Appointment{}, ErrAppointmentNotFound
	} else if err != nil {
		// returning if err is still not nil
		return Appointment{}, err
	}

	return a, nil
}

func ListAppointments(patientId int64, s Status, db *sql.DB) ([]Appointment, error) {
	query := `SELECT id, patient_id, status, scheduled_for, duration_minutes FROM appointments 
		WHERE 
			(NULLIF(:patient_id, 0) IS NULL OR patient_id = :patient_id) 
		AND 
			(NULLIF(:status, "") IS NULL OR status = :status)
	;`

	var appointments []Appointment
	rows, err := db.Query(
		query,
		sql.Named("patient_id", patientId),
		sql.Named("status", s),
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

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

/*
Checks overlap of appointment schedule time
*/
func CheckAppointmentOverlap(scheduleFor string, durationMinutes int, db *sql.DB) (bool, error) {
	t, err := time.Parse(time.RFC3339, scheduleFor)

	if err != nil {
		return false, err
	}

	scheduleEnd := t.Add(time.Duration(durationMinutes) * time.Minute).Format(time.RFC3339)
	appointments, err := ListAppointmentsByScheduleRange(scheduleFor, scheduleEnd, db)
	if err != nil {
		return false, err
	}

	if len(appointments) == 0 {
		return false, nil
	}

	return true, nil
}

/*
Gets avaibale schedule time for appointments
Should get from config the working hours and get when the next appointment can be scheduled (default 30 min)
*/
func GetAvailableScheduleForAppointment() {}

/*
Lists appointments by schedule datetime range (time.RFC3339 Fromat)
*/
func ListAppointmentsByScheduleRange(start string, end string, db *sql.DB) ([]Appointment, error) {
	query := `SELECT id, patient_id, status, scheduled_for, duration_minutes FROM appointments 
		WHERE scheduled_for > ? OR scheduled_for < ?
	;`

	rows, err := db.Query(query, start, end)
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

func UpdateAppointmentStatus(id int64, s Status, db *sql.DB) error {
	query := `UPDATE appointments SET status = ? WHERE id = ?;`

	_, err := db.Exec(query, s, id)
	if err != nil {
		return err
	}

	return nil
}
