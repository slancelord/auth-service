package database

import (
	"database/sql"
	"fmt"
	"log"
	"sync"

	"auth-service/internal/config"

	_ "github.com/lib/pq"
)

var (
	db   *sql.DB
	once sync.Once
)

func Init() {
	once.Do(func() {
		config := config.GetConfig()

		dsn := fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			config.DBHost, config.DBPort, config.DBUser, config.DBPassword, config.DBName,
		)

		var err error
		db, err = sql.Open("postgres", dsn)
		if err != nil {
			log.Fatalf("[ERROR] Failed to open database connection: %v", err)
		}

		err = db.Ping()
		if err != nil {
			log.Fatalf("[ERROR] Database ping failed: %v", err)
		}
	})
}

func GetDB() *sql.DB {
	if db == nil {
		log.Fatal("[ERROR] Database not initialized")
	}
	return db
}
