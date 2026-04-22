package services

import "errors"

var (
	ErrAppointmentOverlap = errors.New("detected schedule overlap")
	ErrPatientNotFound    = errors.New("patient not found")
)
