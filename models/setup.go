package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() *gorm.DB {
	dsn := "root:loko2455@tcp(192.168.0.31:3307)/otontech?charset=utf8mb4&parseTime=True&loc=Local"
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database! " + err.Error())
	}

	database.AutoMigrate(
		&Command{},
		&User{},
		&Role{},
		&Step{},
		&Param{},
		&Arg{},
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
