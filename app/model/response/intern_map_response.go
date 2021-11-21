package response

import "time"

type InternshipMap struct {
	ID               int       `json:"id"`
	IDInternship     int       `json:"id_internship"`
	IDUser           int       `json:"id_user"`
	IsApply          int       `json:"isapply"`
	RegistrationDate time.Time `json:"registrationdate"`
	Process          string    `json:"process"`
}
