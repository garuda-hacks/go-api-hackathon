package entity

import (
	"github.com/garuda-hacks/go-api-hackathon/config"
)

type InterestMap struct {
	ID         int `json:"id"`
	IDInterest int `json:"id_interest"`
	IDUser     int `json:"id_user"`
}

func (InterestMap) TableName() string {
	return config.C.Database.Schema.GH + ".interest_map"
}
