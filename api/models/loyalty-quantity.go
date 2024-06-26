package models

import (
	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

type LoyaltyQuantity struct {
	gorm.Model
	Quantity      uint        `json:"quantity" validate:"nonzero"`
	LoyaltyCardID uint        `json:"loyalty_card_id"` // foreing key
	LoyaltyCard   LoyaltyCard `gorm:"foreignkey:LoyaltyCardID"`
}

func ValidateDataLoyaltyQuantity(loyaltyQuantity *LoyaltyQuantity) error {
	if err := validator.Validate(loyaltyQuantity); err != nil {
		return err
	}
	return nil
}
