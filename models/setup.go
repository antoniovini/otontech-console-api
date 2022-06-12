package models

import (
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() *gorm.DB {
	dsn := os.Getenv("CONNECTION_STRING")
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database! " + err.Error())
	}

	database.AutoMigrate(
		&Command{},
		&User{},
		&Role{},
		&Arg{},
		&Program{},
	)

	defaultRoles := []Role{
		{
			Name:        "default",
			Description: "Default role",
			Level:       0,
		},
		{
			Name:        "admin",
			Description: "Admin role",
			Level:       100,
		},
	}

	for i := range defaultRoles {
		database.Model(&Role{}).Create(&defaultRoles[i])
	}

	DB = database
	return database
}
