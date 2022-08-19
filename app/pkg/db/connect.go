package db

import (
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/AndresCRamos/Simple-Personal-Finances/models/class"
	incomesource "github.com/AndresCRamos/Simple-Personal-Finances/models/income_source"
	token "github.com/AndresCRamos/Simple-Personal-Finances/pkg/auth/models/token"
	auth_user "github.com/AndresCRamos/Simple-Personal-Finances/pkg/auth/models/user"
	"github.com/AndresCRamos/Simple-Personal-Finances/pkg/utils"
)

func tryConnection(connectionString string) (*gorm.DB, error) {
	var dbConn *gorm.DB
	var err error
	for i := 0; i < 3; i++ {
		dbConn, err = gorm.Open(postgres.New(postgres.Config{
			DSN: connectionString,
		}), &gorm.Config{},
		)
		if err == nil {
			break
		}
		log.Println("Can't connect to database, reconnecting...")
		time.Sleep(time.Minute)
	}
	return dbConn, err
}

func Connect(connectionString string) {
	log.Println("Connecting to Database...")

	dbConn, err := tryConnection(connectionString)
	if err != nil {
		log.Fatal(err)
		panic("Failed to connect to DB")
	}
	utils.Instance = dbConn
	log.Println("Connected to Database")
}

func Migrate() {
	log.Println("Database Migration Started....")
	utils.Instance.AutoMigrate(&incomesource.IncomeSource{})
	utils.Instance.AutoMigrate(&class.Bill{})
	utils.Instance.AutoMigrate(&class.Earning{})
	utils.Instance.AutoMigrate(&auth_user.User{})
	utils.Instance.AutoMigrate(&token.Token{})
	log.Println("Database Migration Completed")
}
