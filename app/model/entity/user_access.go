package entity

import (
	"time"

	"github.com/garuda-hacks/go-api-hackathon/config"
)

type UserAccess struct {
	ID         int       `json:"id"`
	IDUser     int       `json:"id_user"`
	Token      string    `json:"token"`
	LastUpdate time.Time `json:"last_update"`
}

func (UserAccess) TableName() string {
	return config.C.Database.Schema.GH + ".user_access"
}
