package services

import (
	"testing"

	"github.com/bartodes/smilelog/internals/models"
	"github.com/bartodes/smilelog/internals/testutils"
)

func TestCreatePatient(t *testing.T) {
	db := testutils.SetupTestDB(t)

	tests := []struct {
		name    string
		patient models.Patient
		wantErr bool
	}{
		{"patient with all fields", models.Patient{Name: "Juan", LastName: "Doe", Email: "jhonnydoe@mail.com", PhoneNumber: 123456790}, false},
		{"patient missing lastname and phone number", models.Patient{Name: "Juan", LastName: "", Email: "jhonny@mail.com", PhoneNumber: 0}, false},
		{"patient with all fields 2", models.Patient{Name: "Sandra", LastName: "Fernandez", Email: "sanfer@mail.com", PhoneNumber: 123456777}, false},
		{"patient missing lastname and email", models.Patient{Name: "Sandra", LastName: "", Email: "", PhoneNumber: 123456778}, false},
		{"patient missing name", models.Patient{Name: "", LastName: "", Email: "", PhoneNumber: 12345678}, true},
		{"duplicated unique fields", models.Patient{Name: "Juan", LastName: "Doe", Email: "jhonnydoe@mail.com", PhoneNumber: 123456790}, true},
		{"missing unique fields", models.Patient{Name: "Juan", LastName: "Doe", Email: "", PhoneNumber: 0}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := CreatePatient(tt.patient, db)
			if (err != nil) != tt.wantErr {
				t.Errorf(
					"CreatePatient() error: %v\nwantErr = %v",
					err,
					tt.wantErr,
				)
			}
		})
	}
}
