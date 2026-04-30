package models

import (
	"errors"
	"time"
)

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
	ScheduledFor    string // SQLite datetime => "2006-01-02 15:04" Fromat
	DurationMinutes int
}

type AppointmentStats struct {
	Total     int
	Completed int
	Canceled  int
	NoShow    int
}

type WorkingSchedule struct {
	Days  map[time.Weekday]bool
	Start int //0-23
	End   int //0-23
}

func (a Appointment) IsCreated() bool {
	return a.Status == CREATED
}

var ErrInvalidAppointment = errors.New("invalid appointment")

func (a Appointment) IsValid(ws WorkingSchedule) (bool, error) {
	start, err := time.Parse("2006-01-02 15:04", a.ScheduledFor)
	if err != nil {
		return false, err
	}

	end := start.Add(time.Minute * time.Duration(a.DurationMinutes))

	if start.Compare(time.Now()) < 0 {
		return false, nil
	}

	if start.Weekday() != end.Weekday() {
		return false, nil
	}

	if !ws.Days[start.Weekday()] || !ws.Days[end.Weekday()] {
		return false, nil
	}

	if start.Hour() < ws.Start || start.Hour() >= ws.End {
		return false, nil
	}

	if end.Add(-time.Second).Hour() < ws.Start || end.Add(-time.Second).Hour() >= ws.End {
		return false, nil
	}

	return true, nil
}
