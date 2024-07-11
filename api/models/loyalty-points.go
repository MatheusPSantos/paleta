package models

import (
	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

type LoyaltyPoints struct {
	gorm.Model
	Points            float64         `json:"quantity" validate:"nonzero"`
	LoyaltyCampaingID uint            `json:"loyalty_card_id"` // foreing key
	LoyaltyCampaing   LoyaltyCampaing `gorm:"foreignkey:LoyaltyCampaingID"`
}

func ValidateDataLoyaltyPoints(loyaltyPoints *LoyaltyPoints) error {
	if err := validator.Validate(loyaltyPoints); err != nil {
		return err
	}
	return nil
}
