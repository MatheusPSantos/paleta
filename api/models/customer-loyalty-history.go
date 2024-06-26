package models

import (
	"paleta-api/enums"

	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

type CustomerLoyaltyHistory struct {
	gorm.Model
	CustomerID      uint                      `json:"customer_id"`
	LoyaltyCardID   uint                      `json:"loyalty_card_id"`
	Value           uint                      `json:"value"`
	Scale           uint                      `json:"scale" gorm:"comment:'If 2 for example, 1234 means 12.34, if 3 so 1234 means 1.234'"`
	TransactionType enums.TransactionTypeEnum `json:"transaction_type" gorm:"type:VARCHAR(20)"`
	LoyaltyCard     LoyaltyCard               `gorm:"foreignkey:LoyaltyCardID"`
	Customer        Customer                  `gorm:"foreignkey:CustomerID"`
}

func ValidateCustomerLoyaltyHistoryData(l *CustomerLoyaltyHistory) error {
	if err := validator.Validate(l); err != nil {
		return err
	}
	return nil
}
