package models

import "time"

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
	ScheduledFor    time.Time // ISO8601 datetime
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
