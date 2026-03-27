package services

type PatientHistory struct {
	Patient      Patient
	Appointments []Appointment
	Visits       []Visit
}

func GetPatientHistory(patientID int64) (PatientHistory, error)
