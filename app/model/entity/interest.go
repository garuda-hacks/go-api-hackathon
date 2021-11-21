package entity

import (
	"github.com/garuda-hacks/go-api-hackathon/config"
)

type Interest struct {
	ID          int    `json:"id"`
	IDInterest  int    `json:"id_interest"`
	Description string `json:"description"`
}

func (Interest) TableName() string {
	return config.C.Database.Schema.GH + ".interest"
}
