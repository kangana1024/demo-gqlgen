package database

import (
	"log"
	"os"

	"database/sql"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq" // The database driver in use.
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func OpenGormOutput() (*gorm.DB, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error Env")
	}
	sqlDB, err := sql.Open("postgres", os.Getenv("PG_URL_OUTPUT"))

	// defer sqlDB.Close()
	db, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{})

	db.AutoMigrate(Users{}, Links{})

	return db, err
}
