package entity

import (
	"github.com/garuda-hacks/go-api-hackathon/config"
)

type Bookmark struct {
	ID            int    `gorm:"column:id"`
	IDUser        int    `gorm:"column:id_user"`
	IDIntern      int    `gorm:"column:id_intern"`
	IDScholarship int    `gorm:"column:id_scholarship"`
	IDBookmark    int    `gorm:"column:id_bookmark"`
	Name          string `gorm:"column:name"`
}

func (Bookmark) TableName() string {
	return config.C.Database.Schema.GH + ".bookmark"
}
