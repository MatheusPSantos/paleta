package models

import (
	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

type LoyaltyCard struct {
	gorm.Model
	Title       string       `json:"title" validate:"nonzero"`
	Layout      string       `json:"layout" validate:"nonzero"`
	TypeLoyalty *TypeLoyalty `json:"type_loyalty" validate:"nonzero" gorm:"foreignkey:LoyaltyCardID;association_foreignkey:ID"`
}

func ValidateDataLoyaltyCard(loyaltyCard *LoyaltyCard) error {
	if err := validator.Validate(loyaltyCard); err != nil {
		return err
	}
	return nil
}
