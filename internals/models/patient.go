package models

type Patient struct {
	ID          int64
	Name        string
	LastName    string
	Email       string
	PhoneNumber uint
}

func (p Patient) FullName() string {
	if p.LastName != "" {
		return p.Name + " " + p.LastName
	}

	return p.Name
}
