package request

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type RegProfileUser struct {
	IDUser string `json:"id_user" form:"id_user"`
	Token  string `json:"token" form:"token"`
}

func (p RegProfileUser) Validate() error {
	return validation.ValidateStruct(&p,
		validation.Field(&p.IDUser, validation.Required),
		validation.Field(&p.Token, validation.Required),
	)
}
