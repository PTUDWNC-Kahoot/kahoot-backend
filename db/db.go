package db

import (
	"examples/identity/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDatabase(s *config.Specification) *gorm.DB {
	sqlDB := postgres.Open(s.DBConnection)
	db, _ := gorm.Open(sqlDB, &gorm.Config{})
	return db
}
