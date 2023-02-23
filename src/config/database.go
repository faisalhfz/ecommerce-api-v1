package config

import (
	"fmt"
	"os"

	"ecommerce-api/src/entity"

	"gorm.io/gorm"

	"gorm.io/driver/postgres"

	_ "github.com/joho/godotenv/autoload"
)

var DB *gorm.DB
var err error

func Database() {
	DB, err = gorm.Open(postgres.Open(fmt.Sprintf("%v://%v:%v@%v:%v/%v", os.Getenv("DB_DRIVER"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}

func AutoMigrate() {
	DB.AutoMigrate(&entity.Product{})
	DB.AutoMigrate(&entity.Cart{})
	DB.AutoMigrate(&entity.Order{})
}
