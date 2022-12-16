package config

import (
	"os"
)

type DbConfig struct {
	User         string
	Password     string
	Host         string
	Port         string
	Database     string
	AnyMigration bool
}

func GetDbConfig() DbConfig {
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	database := os.Getenv("DB_NAME")
	anyMigration := os.Getenv("ANY_MIGRATION")

	return DbConfig{
		User:         user,
		Password:     password,
		Host:         host,
		Port:         port,
		Database:     database,
		AnyMigration: anyMigration == "true",
	}
}
