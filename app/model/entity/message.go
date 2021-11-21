package entity

import "github.com/garuda-hacks/go-api-hackathon/config"

type Message struct {
	ID        int    `json:"id"`
	IDMessage int    `json:"id_message"`
	IDUser    int    `json:"id_user"`
	Name      string `json:"name"`
	Message   string `json:"message"`
	Avatar    string `json:"avatar"`
}

func (Message) TableName() string {
	return config.C.Database.Schema.GH + ".message"
}
