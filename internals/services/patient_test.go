package services

import (
	"testing"

	"github.com/bartodes/smilelog/internals/models"
)

func TestCreatePatient(t *testing.T) {
	db := SetupTestDB(t)

	tests := []struct {
		name    string
		patient models.Patient
	}{
		{"Create patient with all fields", models.Patient{Name: "Juan", LastName: "Doe", Email: "jhonnydoe@mail.com", PhoneNumber: 123456790}},
		{"Create patient missing lastname and phone number", models.Patient{Name: "Juan", LastName: "", Email: "jhonny@mail.com", PhoneNumber: 0}},
		{"Create patient with all fields 2", models.Patient{Name: "Sandra", LastName: "Fernandez", Email: "sanfer@mail.com", PhoneNumber: 123456777}},
		{"Create patient missing lastname and email", models.Patient{Name: "Sandra", LastName: "", Email: "", PhoneNumber: 123456778}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := CreatePatient(tt.patient, db)
			if err != nil {
				t.Error(err)
			}
		})
	}
}
