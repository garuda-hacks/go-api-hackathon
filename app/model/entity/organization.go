package entity

import (
	"github.com/garuda-hacks/go-api-hackathon/config"
)

type Organization struct {
	ID          int    `json:"id"`
	IDUser      int    `json:"id_user"`
	Institute   string `json:"institute"`
	Location    string `json:"location"`
	Role        string `json:"role"`
	Description string `json:"description"`
	StartDate   string `json:"startdate"`
	EndDate     string `json:"enddate"`
	Photo       string `json:"photo"`
	// UserData    []UserData `gorm:"ForeignKey:IDInternship;save_associations:false" json:"ak_mahasiswa,omitempty"`
}

func (Organization) TableName() string {
	return config.C.Database.Schema.GH + ".organization"
}
