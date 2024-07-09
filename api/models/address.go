package models

import (
	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

type Address struct {
	gorm.Model
	UserID uint `json:"seller_id"` // foreign key

	Street     string `json:"street" validate:"nonzero"`
	City       string `json:"city" validate:"nonzero"`
	ZipCode    string `json:"zip_code" validate:"nonzero"`
	Complement string `json:"complement"`
}

func ValidateDataAddress(address *Address) error {
	if err := validator.Validate(address); err != nil {
		return err
	}
	return nil
}
