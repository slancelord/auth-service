package database

import (
	"auth-service/internal/config"
	"fmt"
	"log"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db   *gorm.DB
	once sync.Once
)

func Init() {
	once.Do(func() {
		cfg := config.GetConfig()

		dsn := fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName,
		)

		var err error
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatalf("[ERROR] Failed to open GORM connection: %v", err)
		}

		sqlDB, err := db.DB()
		if err != nil {
			log.Fatalf("[ERROR] Failed to get sql.DB from GORM: %v", err)
		}

		if err := sqlDB.Ping(); err != nil {
			log.Fatalf("[ERROR] Ping failed: %v", err)
		}
	})
}

func Get() *gorm.DB {
	if db == nil {
		log.Fatal("[ERROR] GORM database not initialized")
	}
	return db
}
