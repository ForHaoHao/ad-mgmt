package database

import (
	"ADMgmtSystem/models"
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func register() []interface{} {
	return []interface{}{
		&models.Users{},
		&models.UsersMeta{},
	}
}

func InitDatabase() error {
	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB_NAME")
	sslMode := os.Getenv("POSTGRES_SSLMODE")

	conn := fmt.Sprintf("host=%s user=%s password=%s port=%s sslmode=%s", host, user, password, port, sslMode)

	// 檢查並創建資料庫
	err := createDatabaseIfNotExists(conn, dbName)

	if err != nil {
		log.Fatalf("Failed to ensure database exists: %v", err)
	}

	db, err := gorm.Open(postgres.Open(fmt.Sprintf("%s dbname=%s", conn, dbName)), &gorm.Config{})

	if err != nil {
		return fmt.Errorf("faild to connect to database: %v", err)
	}

	allModels := register()

	err = db.AutoMigrate(allModels...)

	if err != nil {
		return fmt.Errorf("failed to migrate database: %v", err)
	}

	return nil
}

func createDatabaseIfNotExists(conn, dbName string) error {
	db, err := sql.Open("postgres", conn)

	if err != nil {
		return fmt.Errorf("failed to connect to database server: %w", err)
	}

	defer db.Close()

	var exists bool

	query := fmt.Sprintf("SELECT EXISTS (SELECT 1 FROM pg_database WHERE datname = '%s')", dbName)
	err = db.QueryRow(query).Scan((&exists))

	if err != nil {
		return fmt.Errorf("failed to connect to database server: %w", err)
	}

	if !exists {
		_, err = db.Exec(fmt.Sprintf("CREATE DATABASE %s", dbName))

		if err != nil {
			return fmt.Errorf("faild to create database: %w", err)
		}
		log.Printf("Database %s created successfully!", dbName)
	} else {
		log.Printf("Database %s already exists!", dbName)
	}
	return nil
}
