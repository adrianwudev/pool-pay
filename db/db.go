package db

import (
	"fmt"
	"log"

	"pool-pay/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewConnection(config *config.Config) (*gorm.DB, error) {
	connString := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.Host, config.Port, config.User, config.Password, config.DBName, config.SSLMode,
	)
	db, err := gorm.Open(postgres.Open(connString), &gorm.Config{})
	if err != nil {
		return db, err
	}
	return db, nil
}

func ConnectDb(cfg config.Config) *gorm.DB {
	myDb, err := NewConnection(&cfg)
	if err != nil {
		log.Println(err)
	}
	err = pingDatabase(myDb)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Database connection successful")
	}

	return myDb
}

func pingDatabase(db *gorm.DB) error {
	sqlDb, err := db.DB()
	if err != nil {
		return err
	}

	err = sqlDb.Ping()
	if err != nil {
		return err
	}

	return nil
}
