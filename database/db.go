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

var PgConn *gorm.DB

func register() []interface{} {
	return []interface{}{
		&models.Users{},
		&models.UsersMeta{},
	}
}
func registerSeedData() []interface{} {
	return []interface{}{
		&models.Users{ID: "admin", Account: "admin", Password: "5a1d689fabfeefb613fbf4399f8795e9b54102bdc2ce85d13483dc3e2b97c003",
			PasswordSalt: `()#"(#!%+%`, ErrorCount: 0, Activated: true, Role: 1},
		&models.UsersMeta{UsersID: "admin", Name: "administrator", Email: "test@yahoo.com.tw", Avatar: nil, SendEmail: true},
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

	PgConn, err := gorm.Open(postgres.Open(fmt.Sprintf("%s dbname=%s", conn, dbName)), &gorm.Config{})

	if err != nil {
		return fmt.Errorf("faild to connect to database: %v", err)
	}

	allModels := register()

	err = PgConn.AutoMigrate(allModels...)
	if err != nil {
		return fmt.Errorf("failed to migrate database: %v", err)
	}

	seedDatas := registerSeedData()

	for _, seedData := range seedDatas {
		if err := PgConn.FirstOrCreate(seedData, seedData).Error; err != nil {
			fmt.Println(err)
			return err
		}
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
