package postgres

import (
	"log"

	resp "web-service/internal/lib/api/response"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := "host=localhost user=postgres password=pasword dbname=postgres port=5432 sslmode=disable"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Не удалось подключится к базе данных: %v", err)
	}

	DB.AutoMigrate(&resp.Message{})
}
