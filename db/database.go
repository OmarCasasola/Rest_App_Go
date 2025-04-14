package db

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dsn = "host=localhost user=root password=root dbname=rest_api port=5432 sslmode=disable TimeZone=America/Lima"
var Database *gorm.DB

func ConnectDatabase() {
	var err error
	Database, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Successfully connected to database")
}
