package entity

import "github.com/garuda-hacks/go-api-hackathon/config"

type Internship struct {
	ID          int    `json:"id"`
	IDIntern    int    `json:"id_intern"`
	Title       string `gorm:"column:title"`
	Provider    string `json:"provider"`
	City        string `json:"city"`
	Country     string `json:"country"`
	Description string `json:"description"`
	Requirement string `json:"requirement"`
	StartDate   string `gorm:"column:startdate"`
	EndDate     string `gorm:"column:enddate"`
	IsDone      string `json:"isdone"`
	Avatar      string `json:"avatar"`
	Target      string `gorm:"column:target"`
	Duration    int    `gorm:"column:duration"`
}

func (Internship) TableName() string {
	return config.C.Database.Schema.GH + ".internship"
}
