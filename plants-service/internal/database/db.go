package database

import (
	"database/sql"
	"fmt"
	"log"
	"plant-care-app/plants-service/config"
	"plant-care-app/plants-service/internal/models"

	_ "github.com/lib/pq" // PostgreSQL driver
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	cfg := config.GetInstance()

	// First connect to postgres database to check/create our database
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=postgres sslmode=disable",
		cfg.GetDBHost(), cfg.GetDBPort(), cfg.GetDBUser(), cfg.GetDBPassword())

	fmt.Printf("Connecting to postgres with: %s\n", psqlInfo)
	adminDB, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalf("❌ Không kết nối được Postgres hệ thống: %v", err)
	}
	defer adminDB.Close()

	// Check if database exists
	exists := isDatabaseExist(adminDB, cfg)
	if !exists {
		createNewDatabase(adminDB, cfg)
	} else {
		fmt.Printf("✅ Database %s đã tồn tại.\n", cfg.GetDBName())
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.GetDBHost(), cfg.GetDBUser(), cfg.GetDBPassword(), cfg.GetDBName())
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(&models.Plant{}, &models.Schedule{}, &models.Species{})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Database migrated successfully")

	DB = db
}

func isDatabaseExist(adminDB *sql.DB, cfg *config.Config) bool {
	var exists bool
	err := adminDB.QueryRow("SELECT EXISTS(SELECT 1 FROM pg_database WHERE datname = $1)", cfg.GetDBName()).Scan(&exists)
	if err != nil {
		log.Fatalf("❌ Lỗi kiểm tra database: %v", err)
	}

	return exists
}

func createNewDatabase(adminDB *sql.DB, cfg *config.Config) {
	// Create the database with owner
	createDBQuery := fmt.Sprintf("CREATE DATABASE %s WITH OWNER = %s", cfg.GetDBName(), cfg.GetDBUser())
	_, err := adminDB.Exec(createDBQuery)
	if err != nil {
		log.Fatalf("❌ Tạo database thất bại: %v", err)
	}
	fmt.Printf("✅ Tạo database %s thành công.\n", cfg.GetDBName())
}
