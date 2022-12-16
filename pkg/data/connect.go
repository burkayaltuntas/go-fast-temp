package data

import (
	"fmt"
	"log"

	"github.com/burkayaltuntas/go-fast-temp/pkg/config"
	"github.com/burkayaltuntas/go-fast-temp/pkg/data/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect(config config.DbConfig) *gorm.DB {
	user := config.User
	password := config.Password
	host := config.Host
	port := config.Port
	database := config.Database
	anyMigration := config.AnyMigration

	dbinfo := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		user,
		password,
		host,
		port,
		database,
	)

	db, err := gorm.Open(postgres.Open(dbinfo), &gorm.Config{})
	if err != nil {
		log.Println("Failed to connect to database")
		panic(err)
	}
	log.Println("connected to db...")

	if anyMigration {
		log.Println("Migrating...")
		err = migrate(db)
		if err != nil {
			log.Println("Failed to migrate")
			panic(err)
		}
	}

	return db
}

func migrate(db *gorm.DB) error {
	db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";")
	return db.AutoMigrate(
		&models.City{},
		&models.User{},
	)
}
