package models

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type CreateCommandInput struct {
	Description string   `json:"description" binding:"required"`
	Activator   string   `json:"activator" binding:"required"`
	Action      string   `json:"action" binding:"required"`
	Roles       []string `json:"roles" binding:"required"`
	Steps       []Step   `json:"steps" binding:"required"`
	Args        []Arg    `json:"args" binding:"required"`
}

type UpdateCommandInput struct {
	Description string   `json:"description"`
	Activator   string   `json:"activator"`
	Action      string   `json:"action"`
	Roles       []string `json:"roles"`
	Steps       []Step   `json:"steps"`
	Args        []Arg    `json:"args"`
}

type Command struct {
	Id          uint      `json:"id" gorm:"primary_key"`
	UniqueId    uuid.UUID `json:"uniqueId" gorm:"type:char(36);not null;unique"`
	Description string    `json:"description"`
	Activator   string    `json:"activator" gorm:"type:varchar(16);not null;unique"`
	Action      string    `json:"action" gorm:"type:varchar(16)"`
	Roles       []Role    `json:"roles" gorm:"many2many:command_roles;"`
	Steps       []Step    `json:"steps" gorm:"many2many:command_steps;"`
	Args        []Arg     `json:"args" gorm:"many2many:command_args;"`
}

func (command *Command) BeforeCreate(db *gorm.DB) (err error) {
	command.UniqueId = uuid.NewV4()
	return
}
