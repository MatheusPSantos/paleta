package models

import (
	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

type LoyaltyQuantity struct {
	gorm.Model
	Quantity    int          `json:"quantity" validate:"nonzero"`
	TypeLoyalty *TypeLoyalty `json:"type_loyalty" validate:"nonzero" gorm:"foreignkey:LoyaltyQuantityID;association_foreignkey:ID"`
}

func ValidateDataLoyaltyQuantity(loyaltyQuantity *LoyaltyQuantity) error {
	if err := validator.Validate(loyaltyQuantity); err != nil {
		return err
	}
	return nil
}
