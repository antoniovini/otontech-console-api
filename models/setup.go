package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDatabase() *gorm.DB {
	dsn := "root:antoniovini123@tcp(127.0.0.1:3306)/otontech?charset=utf8mb4&parseTime=True&loc=Local"
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database! " + err.Error())
	}

	database.AutoMigrate(&Command{})

	return database
}
