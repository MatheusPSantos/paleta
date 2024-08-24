package models

import "gopkg.in/validator.v2"

type Login struct {
	User     string `json:"user" validate:"nonzero"`
	Password string `json:"password" validate:"nonzero"`
}

func ValidateDataLogin(login *Login) error {
	if err := validator.Validate(login); err != nil {
		return err
	}
	return nil
}
