package services

import "database/sql"

type PatientHistory struct {
	Patient      Patient
	Appointments []Appointment
	Visits       []Visit
}

func GetPatientHistory(ph PatientHistory, patientId int64, db *sql.DB) (PatientHistory, error) {
	// get patient
	// get all the appointments from the patient
	// get all the visists from the appointments

	return ph, nil
}
