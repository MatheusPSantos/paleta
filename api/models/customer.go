package models

import (
	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model
	Name    string     `json:"name" validate:"nonzero"`
	Email   string     `json:"email" validate:"nonzero" gorm:"unique"`
	Phone   string     `json:"phone" validate:"nonzero,regexp=^[0-9]+$"`
	CPF     string     `json:"cpf" validate:"nonzero,regexp=^[0-9]+$" gorm:"unique"`
	Address []*Address `json:"addresses" validate:"nonzero" gorm:"foreignkey:CustomerID;association_foreignkey:ID"`
}

func ValidateDataCustomer(customer *Customer) error {
	if err := validator.Validate(customer); err != nil {
		return err
	}
	return nil
}
