package entity

import (
	"github.com/garuda-hacks/go-api-hackathon/config"
)

type Ugc struct {
	ID           int    `json:"id"`
	IDUgc        int    `json:"id_ugc"`
	IDUser       int    `json:"id_user"`
	IDInterest   int    `json:"id_interest"`
	Posting      string `json:"posting"`
	PostingDate  string `gorm:"column:postingdate"`
	PhotoPosting string `gorm:"column:photopost"`
	VideoPosting string `gorm:"column:videopost"`
}

func (Ugc) TableName() string {
	return config.C.Database.Schema.GH + ".ugc"
}
