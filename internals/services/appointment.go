package services

import "time"

type Appointment struct {
	ID              int64
	PatientID       int64
	Status          string
	ScheduledFor    time.Time // ISO8601 datetime
	DurationMinutes int
}
