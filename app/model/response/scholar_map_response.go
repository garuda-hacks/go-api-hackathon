package response

import "time"

type ScholarshipMap struct {
	ID               int       `json:"id"`
	IDScholarship    int       `json:"id_scholarship"`
	IDUser           int       `json:"id_user"`
	IsApply          int       `json:"isapply"`
	RegistrationDate time.Time `json:"registrationdate"`
	Process          string    `json:"process"`
}
