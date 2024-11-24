package model

import "github.com/mcfiet/goDo/utils"

type User struct {
	utils.Base
	Username string `json:"username" gorm:"unique,not null"`
	Password string `json:"password" gorm:"not null"`
}
