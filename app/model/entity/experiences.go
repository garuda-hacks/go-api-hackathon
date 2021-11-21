package entity

import (
	"github.com/garuda-hacks/go-api-hackathon/config"
)

type Experiences struct {
	ID              int    `json:"id"`
	IDUser          int    `json:"id_user"`
	Title           string `json:"title"`
	EmployementType string `json:"employementtype"`
	CompanyName     string `json:"companyname"`
	City            string `json:"city"`
	Country         string `json:"country"`
	StartDate       string `json:"startdate"`
	EndDate         string `json:"enddate"`
	Description     string `json:"description"`
	// UserData        []UserData `gorm:"ForeignKey:IDExperiences;save_associations:false" json:"ak_mahasiswa,omitempty"`
}

func (Experiences) TableName() string {
	return config.C.Database.Schema.GH + ".experiences"
}
