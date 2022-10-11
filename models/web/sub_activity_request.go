package web

import "time"

type SubActivityRequest struct {
	ActivityId  uint      `json:"activity_id" binding:"required,numeric"`
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description" binding:"required"`
	User_id     uint64    `json:"user_id" binding:"required"`
	StartTime   time.Time `json:"start_time"`
	EndTime     time.Time `json:"end_time"`
	UpdatedBy   string    `json:"updated_by"`
	DeletedBy   string    `json:"deleted_by"`
}

type SubActivityUpdateRequest struct {
	ID          uint
	ActivityId  uint      `json:"activity_id" binding:"required,numeric"`
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description" binding:"required"`
	User_id     uint64    `json:"user_id" binding:"required"`
	StartTime   time.Time `json:"start_time"`
	EndTime     time.Time `json:"end_time"`
	UpdatedBy   string    `json:"updated_by" binding:"required"`
	DeletedBy   string    `json:"deleted_by"`
}
