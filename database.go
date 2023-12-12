package main

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbInstance struct {
	Db *gorm.DB
}

var Database DbInstance

func ConnectDb() {
	dsn := "host=localhost user=postgres dbname=postgres password=kanaka port=5432 sslmode=disable" // Update with your PostgreSQL connection parameters
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{}) // Use postgres.Open

	if err != nil {
		log.Fatal("Failed to connect to the database! \n", err)
		os.Exit(2)
	}

	log.Println("Connected Successfully to Database")
	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("Running Migrations")

	db.AutoMigrate(&User{}, &Product{}, &Order{})

	Database = DbInstance{
		Db: db,
	}
}
