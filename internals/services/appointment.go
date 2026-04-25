package services

import (
	"database/sql"

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
func CheckAppointmentOverlap(scheduledFor string, durationMinutes int, db *sql.DB) (bool, error) {
	query := `SELECT count(*) FROM appointments 
		WHERE datetime(scheduled_for, '+' || :duration_minutes || ' minutes') > :scheduled_for
		AND scheduled_for < datetime(scheduled_for, '+' || :duration_minutes || ' minutes')
	;`

	var overlappedAppointments int

	err := db.QueryRow(
		query,
		sql.Named("scheduled_for", scheduledFor),
		sql.Named("duration_minutes", durationMinutes),
	).Scan(&overlappedAppointments)

	if err != nil {
		return false, err
	}

	if overlappedAppointments > 0 {
		return true, nil
	}

	return false, nil
}

/*
Gets avaibale schedule time for appointments
Should get from config the working hours and get when the next appointment can be scheduled (default 30 min)
*/
func GetAvailableScheduleForAppointment() {}

/*
Lists appointments by schedule datetime range ("2006-01-02 15:04:05" Fromat)
*/
func ListAppointmentsByScheduleRange(start string, end string, db *sql.DB) ([]Appointment, error) {
	query := `SELECT id, patient_id, status, scheduled_for, duration_minutes FROM appointments 
		WHERE scheduled_for > ? AND scheduled_for < ?
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
