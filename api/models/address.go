package models

import "gorm.io/gorm"

type Address struct {
	gorm.Model
	SellerID   uint   `json:"seller_id"` // foreign key
	Street     string `json:"street" validate:"nonzero"`
	City       string `json:"city" validate:"nonzero"`
	CEP        string `json:"cep" validate:"nonzero"`
	Complement string `json:"complement"`
}
