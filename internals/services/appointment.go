package services

import "time"

type Appointment struct {
	ID        int64
	PatientID int64
	Status    string
	//not sure what time format
	ScheduledFor time.Time
	Duration     time.Duration
}
