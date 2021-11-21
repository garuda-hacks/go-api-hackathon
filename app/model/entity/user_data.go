package entity

import (
	"github.com/garuda-hacks/go-api-hackathon/config"
)

type UserData struct {
	ID             int    `json:"id"`
	IDUser         int    `json:"id_user"`
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	Handphone      string `json:"handphone"`
	Avatar         string `json:"avatar"`
	Following      int    `json:"following"`
	Followers      int    `json:"followers"`
	Community      int    `json:"community"`
	Instagram      string `json:"instagram"`
	Facebook       string `json:"facebook"`
	Twitter        string `json:"twitter"`
	Linkedin       string `json:"linkedin"`
	IDExperiences  int    `json:"id_experiences"`
	IDEducation    int    `json:"id_education"`
	IDOrganization int    `json:"id_organization"`
	IDNotification int    `json:"id_notification"`
	IDInterest     int    `json:"id_interest"`
	// Experiences    *Experiences  `gorm:"ForeignKey:IDExperiences;save_associations:false" json:"lv_sistemkuliah,omitempty"`
	// Education      *Education    `gorm:"ForeignKey:IDEducation;save_associations:false" json:"lv_sistemkuliah,omitempty"`
	// Organization   *Organization `gorm:"ForeignKey:IDOrganizaztion;save_associations:false" json:"lv_sistemkuliah,omitempty"`
	// Interest       *InterestMap  `gorm:"ForeignKey:IDInterest;save_associations:false" json:"lv_sistemkuliah,omitempty"`
	// Notification   *Notification `gorm:"ForeignKey:IDNotification;save_associations:false" json:"lv_sistemkuliah,omitempty"`
}

func (UserData) TableName() string {
	return config.C.Database.Schema.GH + ".user_data"
}
