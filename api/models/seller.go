package models

import "gorm.io/gorm"

type Seller struct {
	gorm.Model
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	CNPJ  string `json:"cnpj"`
}
