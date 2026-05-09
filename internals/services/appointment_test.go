package services

import (
	"testing"
	"time"

	"github.com/bartodes/smilelog/internals/models"
	"github.com/bartodes/smilelog/internals/testutils"
)

func TestCreateAppointment(t *testing.T) {
	db := testutils.SetupTestDB(t)

	p, err := CreatePatient(models.Patient{
		Name:        "Juan",
		LastName:    "Doe",
		Email:       "jhonnydoe@mail.com",
		PhoneNumber: 1234567890,
	}, db)

	if err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		name        string
		appointment models.Appointment
		wantErr     bool
	}{
		{"appointment with all fields", models.Appointment{PatientID: p.ID, ScheduledFor: time.Now().Add(time.Minute).Format("2006-01-02 15:04"), DurationMinutes: 30}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := CreateAppointment(tt.appointment, db)
			if (err != nil) != tt.wantErr {
				t.Errorf(
					"CreateAppointment() error: %v\nwantErr = %v",
					err,
					tt.wantErr,
				)
			}
		})
	}
}
