package models

import (
	"gorm.io/gorm"
)

type CreateArgInput struct {
	Name        string `json:"name" binding:"required"`
	Required    bool   `json:"required" binding:"required"`
	Identifier  string `json:"identifier" binding:"required"`
	Description string `json:"description"`
}

type UpdateArgInput struct {
	Name        string `json:"name"`
	Required    bool   `json:"required"`
	Identifier  string `json:"identifier"`
	Description string `json:"description"`
}

type Arg struct {
	gorm.Model
	Name        string `json:"name" gorm:"type:varchar(255);not null"`
	Required    bool   `json:"required" gorm:"not null"`
	Identifier  string `json:"identifier" gorm:"type:varchar(1);unique"`
	Description string `json:"description" gorm:"type:varchar(255)"`
}
