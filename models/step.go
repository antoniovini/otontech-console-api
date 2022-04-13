package models

import (
	"gorm.io/gorm"
)

type CreateStepInput struct {
	Name   string  `json:"name" binding:"required"`
	Params []Param `json:"params" binding:"required"`
}

type UpdateStepInput struct {
	Name   string  `json:"activator"`
	Params []Param `json:"params"`
}

type Step struct {
	gorm.Model
	Name   string  `json:"name" gorm:"type:varchar(255);not null;unique"`
	Params []Param `gorm:"many2many:step_params;"`
}
