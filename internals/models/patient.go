package models

type Patient struct {
	ID          int64
	Name        string
	LastName    string
	Email       string
	PhoneNumber uint
}

type PatientHistorySummary struct {
	Patient         Patient
	Stats           AppointmentStats
	LastAppointment Appointment
	LastVisit       Visit
}
