package models

type Program struct {
	ID            uint   `json:"id" gorm:"primary_key"`
	Name          string `json:"name" gorm:"type:varchar(16);not null;unique"`
	Description   string `json:"description"`
	Url           string `json:"url"`
	RequiredLevel uint   `json:"required_level" gorm:"not null"`
}

type CreateProgram struct {
	Name          string `json:"name" binding:"required"`
	Description   string `json:"description" binding:"required"`
	Url           string `json:"url" binding:"required"`
	RequiredLevel uint   `json:"required_level" binding:"required"`
}

type UpdateProgram struct {
	Name          string `json:"name"`
	Description   string `json:"description"`
	Url           string `json:"url"`
	RequiredLevel uint   `json:"required_level"`
}
