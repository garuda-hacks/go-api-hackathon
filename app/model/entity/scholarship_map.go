package entity

import (
	"time"

	"github.com/garuda-hacks/go-api-hackathon/config"
)

type ScholarshipMap struct {
	ID               int       `json:"id"`
	IDScholarship    int       `json:"id_scholarship"`
	IDUser           int       `json:"id_user"`
	IsApply          int       `json:"isapply"`
	RegistrationDate time.Time `json:"registrationdate"`
	Process          string    `json:"process"`
}

func (ScholarshipMap) TableName() string {
	return config.C.Database.Schema.GH + ".scholarship_map"
}
