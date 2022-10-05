package domain

import (
	"gorm.io/gorm"
)

type Activity struct {
	gorm.Model
	Name        string `json:="name" gorm:"type:varchar(255);not null"`
	Description string `json:="description" gorm:"type:varchar(255);not null"`
	CreatedBy   string `json:="createdby" `
	UpdateBy    string `json:="updatedby" `
}
