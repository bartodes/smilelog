package services

import (
	"testing"
	"time"

	"github.com/bartodes/smilelog/internals/models"
	"github.com/bartodes/smilelog/internals/testutils"
)

func TestCreateVisit(t *testing.T) {
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
		name    string
		notes   string
		wantErr bool
	}{
		{"visit with notes", "The procedure was performed correctly on the patient", false},
	}

	for _, tt := range tests {
		a, err := CreateAppointment(models.Appointment{
			PatientID:       p.ID,
			ScheduledFor:    time.Now().Add(time.Minute).Format("2006-01-02 15:04"),
			DurationMinutes: 30,
		}, db)

		if err != nil {
			t.Fatal(err)
		}
		t.Run(tt.name, func(t *testing.T) {
			_, err := CreateVisit(a.ID, tt.notes, db)
			if (err != nil) != tt.wantErr {
				t.Errorf(
					"CreateVisit() error: %v\nwantErr = %v",
					err,
					tt.wantErr,
				)
			}
		})
	}
}
