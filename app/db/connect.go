package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	auth_user "github.com/AndresCRamos/Simple-Personal-Finances/auth/models/user"
	bill "github.com/AndresCRamos/Simple-Personal-Finances/models/bills"
	earning "github.com/AndresCRamos/Simple-Personal-Finances/models/earning"
	incomesource "github.com/AndresCRamos/Simple-Personal-Finances/models/income_source"
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
	utils.Instance.AutoMigrate(&incomesource.IncomeSource{})
	utils.Instance.AutoMigrate(&bill.Bill{})
	utils.Instance.AutoMigrate(&earning.Earning{})
	utils.Instance.AutoMigrate(&auth_user.User{})
	log.Println("Database Migration Completed")
}
