package repository

import (
	"auth/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Open() error {
	gormDB, err := gorm.Open(postgres.Open("host=localhost user=root password=root dbname=authDB sslmode=disable"), &gorm.Config{})
	if err != nil {
		panic(err)
		return err
	}

	DB = gormDB
	gormDB.AutoMigrate(&model.User{})

	return nil
}
