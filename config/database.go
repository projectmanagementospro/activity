package config

import (
	"activity/helper"
	"activity/models/domain"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB() *gorm.DB {
	err := godotenv.Load()
	helper.PanicIfError(err)

	dbURL := "postgres://root:root@172.27.1.3:5432/activity?sslmode=disable"
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	helper.PanicIfError(err)
	db.AutoMigrate(&domain.Activity{}, &domain.SubActivity{})
	return db
}

func CloseDB(db *gorm.DB) {
	dbSQL, err := db.DB()
	helper.PanicIfError(err)
	dbSQL.Close()
}
