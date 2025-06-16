package database

import (
	"database/sql"
	"fmt"
	"log"
	"plant-care-app/user-service/config"
	"plant-care-app/user-service/internal/models"

	_ "github.com/lib/pq" // PostgreSQL driver
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect(cfg *config.Config) {
	// First connect to postgres database to check/create our database
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=postgres sslmode=disable",
		cfg.DbHost, cfg.DbPort, cfg.DbUser, cfg.DbPassword)

	fmt.Printf("Connecting to postgres with: %s\n", psqlInfo)
	adminDB, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalf("❌ Không kết nối được Postgres hệ thống: %v", err)
	}
	defer adminDB.Close()

	// Check if database exists
	var exists bool
	err = adminDB.QueryRow("SELECT EXISTS(SELECT 1 FROM pg_database WHERE datname = $1)", cfg.DbName).Scan(&exists)
	if err != nil {
		log.Fatalf("❌ Lỗi kiểm tra database: %v", err)
	}

	if !exists {
		// Create the database with owner
		createDBQuery := fmt.Sprintf("CREATE DATABASE %s WITH OWNER = %s", cfg.DbName, cfg.DbUser)
		_, err = adminDB.Exec(createDBQuery)
		if err != nil {
			log.Fatalf("❌ Tạo database thất bại: %v", err)
		}
		fmt.Printf("✅ Tạo database %s thành công.\n", cfg.DbName)
	} else {
		fmt.Printf("✅ Database %s đã tồn tại.\n", cfg.DbName)
	}

	// Now connect to our actual database
	dbInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DbHost, cfg.DbPort, cfg.DbUser, cfg.DbPassword, cfg.DbName)

	db, err := gorm.Open(postgres.Open(dbInfo), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Database migrated successfully")

	DB = db
}
