package config

import (
	"cell-pos/server/helper"
	"cell-pos/server/models"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connection() {
	err := godotenv.Load()
	helper.PanicIfError(err)
	
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	name := os.Getenv("DB_NAME")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")

	dsn := user + ":" + pass + "@tcp(" + host + ":" + port + ")/" + name + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	helper.PanicIfError(err)

	err = db.AutoMigrate(
		&models.User{},
		&models.Brand{},
		&models.Category{},
	)
	helper.PanicIfError(err)

	DB = db
}