package models

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Authentication struct {
	ID       uint   `gorm:"primaryKey"`
	UserID   uint   `gorm:"not null;uniqueIndex"`
	Password string `gorm:"not null"`
}

// HashPassword criptografa a senha utilizando bcrypt
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// CheckPassword compara uma senha fornecida com a senha criptografada armazenada
func (auth *Authentication) CheckPassword(password string) bool {
	fmt.Println("entrada >> ", password)
	fmt.Println("cripto >>> ", auth.Password)
	err := bcrypt.CompareHashAndPassword([]byte(auth.Password), []byte(password))
	fmt.Println(err.Error())
	return err == nil
}

// Método para atualizar a senha criptografada
func (auth *Authentication) UpdatePassword(newPassword string, db *gorm.DB) error {
	hashedPassword, err := HashPassword(newPassword)
	if err != nil {
		return err
	}
	auth.Password = hashedPassword
	return db.Save(auth).Error
}
