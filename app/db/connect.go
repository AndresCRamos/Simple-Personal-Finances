package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/AndresCRamos/Simple-Personal-Finances/models/user"
)

func Connect(connectionString string) *gorm.DB {
	log.Println("Connecting to Database...")
	Instance, err := gorm.Open(postgres.New(postgres.Config{
		DSN: connectionString,
	}), &gorm.Config{},
	)
	if err != nil {
		log.Fatal(err)
		panic("Cannot connect to DB")
	}
	log.Println("Connected to Database")
	return Instance
}

func Migrate(Instance *gorm.DB) {
	log.Println("Database Migration Started...")
	Instance.AutoMigrate(&user.User{})
	log.Println("Database Migration Completed")
}
