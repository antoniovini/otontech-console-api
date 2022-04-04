package models

import (
	"errors"
	"html"
	"otontech/console-api/utils/token"
	"sort"
	"strings"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type CreateUserInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type AuthInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UpdateUserInput struct {
}

type User struct {
	gorm.Model
	Username string `gorm:"size:255;not null;unique" json:"username"`
	Password string `gorm:"size:255;not null;" json:"password"`
	Roles    []Role `gorm:"many2many:user_roles;"`
}

func VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func GetUserById(uid uint) (User, error) {
	var u User
	if err := DB.Preload("Roles").First(&u, uid).Error; err != nil {
		return u, errors.New("User not found!")
	}
	u.PrepareGive()
	return u, nil
}

func GetUserByUsername(username string) (User, error) {
	var u User
	if err := DB.Preload("Roles").First(&u, User{Username: username}).Error; err != nil {
		return u, errors.New("User not found!")
	}
	u.PrepareGive()
	return u, nil
}

func (u *User) GetRolesString() []string {
	var roles []string
	for i := range u.Roles {
		roles = append(roles, u.Roles[i].Name)
	}
	return roles
}

func (u *User) PrepareGive() {
	u.Password = ""
}

func (u *User) MaxRoleLevel() uint {
	var levels []int = make([]int, len(u.Roles))
	for i := range levels {
		levels[i] = int(u.Roles[i].Level)
	}
	sort.Ints(levels)
	return uint(levels[len(levels)-1])
}

func (u *User) SaveUser() (*User, error) {
	var err error
	err = DB.Create(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}

func (u *User) CheckUser() (string, error) {
	var err error
	dUser := User{}
	err = DB.Model(User{}).Where("username = ?", u.Username).Take(&dUser).Error
	if err != nil {
		return "", err
	}
	err = VerifyPassword(u.Password, dUser.Password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}
	token, err := token.GenerateToken(dUser.ID)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}
	return token, nil
}

func (u *User) BeforeCreate(db *gorm.DB) (err error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	u.Password = string(hashedPassword)
	u.Username = html.EscapeString(strings.TrimSpace(u.Username))

	defaultRole := Role{}

	if err := db.Model(&Role{}).Where("name = ?", "default").Take(&defaultRole).Error; err != nil {
		return err
	}

	u.Roles = append(u.Roles, defaultRole)

	return nil
}
