package models

import (
	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string     `json:"name" validate:"nonzero"`
	Email    string     `json:"email" validate:"nonzero" gorm:"unique"`
	Phone    string     `json:"phone" validate:"nonzero,regexp=^[0-9]+$"`
	Cpf_Cnpj string     `json:"cpf_cnpj" validate:"nonzero,regexp=^[0-9]+$" gorm:"unique"`
	IsSeller bool       `json:"is_seller" validate:"nonzero"`
	address  []*Address `gorm:"foreignkey:UserID;association_foreignkey:ID"`
}

func ValidateDataUser(user *User) error {
	if err := validator.Validate(user); err != nil {
		return err
	}
	return nil
}
