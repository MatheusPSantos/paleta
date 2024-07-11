package models

import (
	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

type LoyaltyCampaing struct {
	gorm.Model
	Title       string `json:"title" validate:"nonzero"`
	Description string `json:"description"`
	Layout      string `json:"layout" validate:"nonzero"`
	IsActive    bool   `json:"is_active"`
	OwnerId     uint   `json:"owner_id" validate:"nonzero"`
	user        User   `gorm:"foreignkey:OwnerId"`
}

func ValidateDataLoyaltyCampaing(loyaltyCard *LoyaltyCampaing) error {
	if err := validator.Validate(loyaltyCard); err != nil {
		return err
	}
	return nil
}
