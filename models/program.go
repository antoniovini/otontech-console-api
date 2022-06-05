package models

type Program struct {
	Id          uint   `json:"id" gorm:"primary_key"`
	Name        string `json:"name" gorm:"type:varchar(16);not null;unique"`
	Description string `json:"description"`
	Url         string `json:"url"`
	Roles       []Role `json:"roles" gorm:"many2many:program_roles;"`
}

type CreateProgram struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	Url         string `json:"url" binding:"required"`
	Roles       []Role `json:"roles" binding:"required"`
}

type UpdateProgram struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Url         string `json:"url"`
	Roles       []Role `json:"roles"`
}
