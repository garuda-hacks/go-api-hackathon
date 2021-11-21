package request

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type RegLoginEmail struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

func (rg RegLoginEmail) Validate() error {
	return validation.ValidateStruct(&rg,
		validation.Field(&rg.Username, validation.Required, validation.Length(6, 50)),
		validation.Field(&rg.Password, validation.Required, validation.Length(6, 50)),
	)
}

// type RegLoginPhone struct {
// 	Handphone      string `json:"handphone" form:"handphone"`
// 	Password   string `json:"password" form:"password"`
// }

// func (rg RegLoginPhone) Validate() error {
// 	return validation.ValidateStruct(&rg,
// 		validation.Field(&rg.Handphone, validation.Required, validation.Length(11, 14), is.UTFNumeric),
// 		validation.Field(&rg.Password, validation.Required, validation.Length(6, 50)),
// 	)
// }
