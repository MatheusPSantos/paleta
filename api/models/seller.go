package models

import (
	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

type Seller struct {
	gorm.Model
	Name      string     `json:"name" validate:"nonzero"`
	Email     string     `json:"email" validate:"nonzero" gorm:"unique"`
	Phone     string     `json:"phone" validate:"nonzero,regexp=^[0-9]+$"`
	CNPJ      string     `json:"cnpj" validate:"nonzero,regexp=^[0-9]+$" gorm:"unique"`
}

func ValidateDataSeller(seller *Seller) error {
	if err := validator.Validate(seller); err != nil {
		return err
	}
	return nil
}
