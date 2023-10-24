package database

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/hacktiv8-ks07-g04/final-project-4/config"
	"github.com/hacktiv8-ks07-g04/final-project-4/domain/entity"
)

var (
	db  *gorm.DB
	err error
)

func Init() {
	Connect()
	Migration()
}

func Connect() {
	config := config.Get().Database
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		config.Host,
		config.User,
		config.Password,
		config.Name,
		config.Port,
	)

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("Failed to connect to database", err)
	}
}

func Migration() {
	// create enum type for role
	db.Exec("CREATE TYPE role AS ENUM ('admin', 'customer');")
	db.AutoMigrate(
		&entity.User{},
		&entity.Category{},
		&entity.Product{},
		&entity.TransactionHistory{},
	)
}

func GetInstance() *gorm.DB {
	if db == nil {
		log.Fatal("Database instance is not initialized")
	}

	return db
}
