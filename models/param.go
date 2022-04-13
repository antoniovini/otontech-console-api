package models

import (
	"gorm.io/gorm"
)

type CreateParamInput struct {
	Key   string `json:"key" binding:"required"`
	Value string `json:"value" binding:"required"`
	Type  string `json:"type" binding:"required"`
}

type UpdateParamInput struct {
	Key   string `json:"key"`
	Value string `json:"value"`
	Type  string `json:"type"`
}

type Param struct {
	gorm.Model
	Key   string `json:"key" gorm:"type:varchar(255);not null;unique"`
	Value string `json:"value" gorm:"type:varchar(255);not null"`
	Type  string `json:"type" gorm:"type:varchar(255);not null"`
}
