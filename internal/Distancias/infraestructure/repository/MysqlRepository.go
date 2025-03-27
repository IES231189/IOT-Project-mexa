package repository

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"iot-project/internal/Distancias/domain"
)

type DBRepository struct {
	db *gorm.DB
}

func NewDBRepository() *DBRepository {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser, dbPassword, dbHost, dbPort, dbName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	db.AutoMigrate(&domain.Distance{})

	return &DBRepository{
		db: db,
	}
}

func (r *DBRepository) Save(distance domain.Distance) error {
	
	distance.IsProximity()

	result := r.db.Create(&distance)
	return result.Error
}
