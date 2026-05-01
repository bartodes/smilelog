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

var (
	ErrScheduleOlderThanCurrentDate  = errors.New("the appointment date cannot be earlier than the current date")
	ErrScheduleNotWithInWorkingHours = errors.New("the appointment must be scheduled with a date that matches the defined range")
)

func (a Appointment) IsValid(ws WorkingSchedule) error {
	start, err := time.Parse("2006-01-02 15:04", a.ScheduledFor)
	if err != nil {
		return err
	}

	end := start.Add(time.Minute * time.Duration(a.DurationMinutes))

	if time.Now().Compare(start) > 0 {
		return ErrScheduleOlderThanCurrentDate
	}

	if !ws.Days[start.Weekday()] || !ws.Days[end.Weekday()] {
		return ErrScheduleNotWithInWorkingHours
	}

	if start.Hour() < ws.Start || start.Hour() >= ws.End {
		return ErrScheduleNotWithInWorkingHours
	}

	if end.Add(-time.Second).Hour() < ws.Start || end.Add(-time.Second).Hour() >= ws.End {
		return ErrScheduleNotWithInWorkingHours
	}

	return nil
}
