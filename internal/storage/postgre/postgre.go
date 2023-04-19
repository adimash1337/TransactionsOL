package postgre

import (
	"awesomeProject/internal/config"
	"awesomeProject/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Dial(cfg config.Config) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(cfg.Database), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	err = db.AutoMigrate(&models.Transaction{})
	return db, nil
}
