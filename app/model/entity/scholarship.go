package entity

import (
	"github.com/garuda-hacks/go-api-hackathon/config"
)

type Scholarship struct {
	ID            int    `json:"id"`
	IDScholarship int    `json:"id_scholarship"`
	Title         string `gorm:"column:judul"`
	Provider      string `json:"provider"`
	Location      string `json:"location"`
	Description   string `json:"description"`
	Requirement   string `json:"requirement"`
	StartDate     string `gorm:"column:startdate"`
	EndDate       string `gorm:"column:enddate"`
	IsDone        bool   `json:"isdone"`
	Avatar        string `json:"avatar"`
	Target        string `gorm:"column:target"`
	Duration      int    `gorm:"column:duration"`
}

func (Scholarship) TableName() string {
	return config.C.Database.Schema.GH + ".scholarship"
}
