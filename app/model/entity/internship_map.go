package entity

import (
	"time"

	"github.com/garuda-hacks/go-api-hackathon/config"
)

type InternshipMap struct {
	ID               int       `json:"id"`
	IDInternship     int       `json:"id_internship"`
	IDUser           int       `json:"id_user"`
	IsApply          int       `json:"isapply"`
	RegistrationDate time.Time `json:"registrationdate"`
	Process          string    `json:"process"`
}

func (InternshipMap) TableName() string {
	return config.C.Database.Schema.GH + ".internship_map"
}
