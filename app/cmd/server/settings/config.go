package settings

import (
	"fmt"
	"log"
	"os"
)

type Config struct {
	db   string
	port string
}

func (c *Config) GetDB() string {
	return c.db
}

func (c *Config) GetPort() string {
	return ":" + c.port
}

func LoadConfig() *Config {
	log.Println("Loading Config...")
	appConfig := Config{
		port: "8080",
	}

	db_host := os.Getenv("DB_HOST")
	db_user := os.Getenv("DB_USER")
	db_pass := os.Getenv("DB_PASS")
	db_name := os.Getenv("DB_NAME")
	db_port := os.Getenv("DB_PORT")

	if db_user == "" || db_name == "" {
		log.Fatal("Can't get database connection params")
	}

	if db_host == "" {
		db_host = "localhost"
	}

	if db_port == "" {
		db_port = "5432"
	}

	connString := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		db_host,
		db_user,
		db_pass,
		db_name,
		db_port,
	)

	if port := os.Getenv("PORT"); port != "" {
		appConfig.port = port
	}

	appConfig.db = connString

	return &appConfig
}
