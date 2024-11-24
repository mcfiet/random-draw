package model

import (
	"github.com/google/uuid"
	"github.com/mcfiet/goDo/utils"
)

type DrawResult struct {
	utils.Base
	GiverId    uuid.UUID `gorm:"not null;unique"`
	ReceiverId uuid.UUID `gorm:"not null;unique"`
}
