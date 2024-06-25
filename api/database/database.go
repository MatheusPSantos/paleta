package database

import (
	"log"
	"paleta-api/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectDatabase() {
	log.Print("Connecting database")
	strConn := "host=localhost user=paleta-api password=paleta-api-password dbname=paleta port=5433 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(strConn))
	if err != nil {
		log.Printf("[Database]: %+v", err.Error())
	}
	DB.AutoMigrate(&models.Seller{}, models.Address{})
}
