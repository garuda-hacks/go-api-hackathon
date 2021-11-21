package entity

import (
	"time"

	"github.com/garuda-hacks/go-api-hackathon/config"
)

type Notification struct {
	ID             int       `json:"id"`
	IDUser         int       `json:"id_user"`
	IDNotification int       `json:"id_notification"`
	Description    string    `json:"description"`
	Date           time.Time `json:"date"`
}

func (Notification) TableName() string {
	return config.C.Database.Schema.GH + ".notification"
}
