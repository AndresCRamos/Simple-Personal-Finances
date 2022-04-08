package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	incomesource "github.com/AndresCRamos/Simple-Personal-Finances/models/income_source"
	"github.com/AndresCRamos/Simple-Personal-Finances/models/user"
	"github.com/AndresCRamos/Simple-Personal-Finances/utils"
)

var err error

func Connect(connectionString string) {
	log.Println("Connecting to Database...")
	utils.Instance, err = gorm.Open(postgres.New(postgres.Config{
		DSN: connectionString,
	}), &gorm.Config{},
	)
	if err != nil {
		log.Fatal(err)
		panic("Cannot connect to DB")
	}
	log.Println("Connected to Database")
}

func Migrate() {
	log.Println("Database Migration Started...")
	utils.Instance.AutoMigrate(&user.User{})
	utils.Instance.AutoMigrate(&incomesource.IncomeSource{})
	log.Println("Database Migration Completed")
}
