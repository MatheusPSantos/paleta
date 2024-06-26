package models

import (
	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

type LoyaltyPoints struct {
	gorm.Model
	Points      float64      `json:"quantity" validate:"nonzero"`
	TypeLoyalty *TypeLoyalty `json:"type_loyalty" validate:"nonzero" gorm:"foreignkey:LoyaltyPointsID;association_foreignkey:ID"`
}

func ValidateDataLoyaltyPoints(loyaltyPoints *LoyaltyPoints) error {
	if err := validator.Validate(loyaltyPoints); err != nil {
		return err
	}
	return nil
}
