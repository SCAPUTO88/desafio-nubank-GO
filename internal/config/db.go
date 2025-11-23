package config

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB(DBURL string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(DBURL), &gorm.Config{})
  if err != nil {
		log.Fatalf("❌ Error connecting to database: %v", err)
	}
	log.Println("✅ Connected to database")

	return db
	}