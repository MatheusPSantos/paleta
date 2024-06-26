package models

import (
	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

type LoyaltyCard struct {
	gorm.Model
	Title    string `json:"title" validate:"nonzero"`
	Layout   string `json:"layout" validate:"nonzero"`
	IsActive bool   `json:"is_active" validate:"nonzero"`
}

func ValidateDataLoyaltyCard(loyaltyCard *LoyaltyCard) error {
	if err := validator.Validate(loyaltyCard); err != nil {
		return err
	}
	return nil
}
