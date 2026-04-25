package models

type Status string

const (
	CREATED   Status = "CREATED"
	COMPLETED Status = "COMPLETED"
	CANCELED  Status = "CANCELED"
	NO_SHOW   Status = "NO_SHOW"
)

type Appointment struct {
	ID              int64
	PatientID       int64
	Status          Status
	ScheduledFor    string // SQLite datetime => "2006-01-02 15:04:05" Fromat
	DurationMinutes int
}

type AppointmentStats struct {
	Total     int
	Completed int
	Canceled  int
	NoShow    int
}

func (a Appointment) IsCreated() bool {
	return a.Status == CREATED
}
