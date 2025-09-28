package models

import (
	"gorm.io/gorm"
)

type RoleManagmentInput struct {
	Username string `json:"username" binding:"required"`
	Role     string `json:"role" binding:"required"`
}

type Role struct {
	gorm.Model
	Name        string `json:"name" gorm:"type:varchar(32);unique;not null"`
	Description string `json:"description" gorm:"size:255"`
	Level       uint   `json:"level" gorm:"not null"`
}
