package config

import (
	"prakerja2/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := "root:123@tcp(127.0.0.1:3306)/prakerja_2?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("connection DB")
	}
	migration()
}

func migration() {
	DB.AutoMigrate(&model.User{})
}
