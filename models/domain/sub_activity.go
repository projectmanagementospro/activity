package domain

import (
	"time"

	"gorm.io/gorm"
)

type SubActivity struct {
	gorm.Model
	ActivityId  uint      `json:"activity_id"`
	Activity    Activity  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Name        string    `json:"name" gorm:"type:varchar(255);not null"`
	Description string    `json:"description" gorm:"type:varchar(255);not null"`
	User_id     uint64    `json:"user_id" gorm:"type:uint;not null"`
	StartTime   time.Time `json:"start_time"`
	EndTime     time.Time `json:"end_time"`
	UpdatedBy   string    `json:"updated_by"`
	DeletedBy   string    `json:"deleted_by"`
}
