package entity

import "github.com/garuda-hacks/go-api-hackathon/config"

type User struct {
	ID       int    `json:"id"`
	IDUser   int    `json:"id_user"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (User) TableName() string {
	return config.C.Database.Schema.GH + ".user"
}
