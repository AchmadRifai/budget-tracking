package db

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func dbConnect() (*gorm.DB, error) {
	host, user, dbname, port, password := os.Getenv("PSQL_HOST"), os.Getenv("PSQL_USER"), os.Getenv("PSQL_DB"), os.Getenv("PSQL_PORT"), os.Getenv("PSQL_PASSWORD")
	database, err := gorm.Open(postgres.Open(fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", host, user, password, dbname, port)), &gorm.Config{})
	return database, err
}
