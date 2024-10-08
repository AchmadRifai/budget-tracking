package db

import (
	"fmt"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DbConnect() *gorm.DB {
	host, user, dbname, port, password := os.Getenv("PSQL_HOST"), os.Getenv("PSQL_USER"), os.Getenv("PSQL_DB"), os.Getenv("PSQL_PORT"), os.Getenv("PSQL_PASSWORD")
	psqlOpen := postgres.Open(fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", host, user, password, dbname, port))

	database, err := gorm.Open(psqlOpen, &gorm.Config{})
	if err != nil {
		panic(err)
	}
	sqlDb, err := database.DB()
	if err != nil {
		panic(err)
	}
	sqlDb.SetMaxIdleConns(10)
	sqlDb.SetMaxOpenConns(50)
	sqlDb.SetConnMaxLifetime(time.Minute)
	return database
}
