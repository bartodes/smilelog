package services

import "errors"

var (
	ErrAppointmentNotCreated = errors.New("appointment status is not 'CREATED'")
	ErrAppointmentNotFound   = errors.New("appointment not found")
	ErrAppointmentOverlap    = errors.New("appointment schedule overlap")
	ErrPatientNotFound       = errors.New("patient not found")
)
