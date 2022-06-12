package models

type Program struct {
	Id          uint   `json:"id" gorm:"primary_key"`
	Name        string `json:"name" gorm:"type:varchar(16);not null;unique"`
	Description string `json:"description"`
	Url         string `json:"url"`
	Role        Role   `json:"role"`
	RoleId      uint
}

type CreateProgram struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	Url         string `json:"url" binding:"required"`
	Role        Role   `json:"role" binding:"required"`
}

type UpdateProgram struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Url         string `json:"url"`
	Role        Role   `json:"role"`
}
