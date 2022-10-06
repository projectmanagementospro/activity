package repository

import (
	"activity/models/domain"
	"errors"

	"gorm.io/gorm"
)

type ActivityRepository interface {
	All() []domain.Activity
	Create(activity domain.Activity) domain.Activity
	Update(activity domain.Activity) domain.Activity
	Delete(activity domain.Activity)
	FindById(id uint) (domain.Activity, error)
}

type ActivityConnectDB struct {
	dbConnect *gorm.DB //connect to database
}

func NewActivityRepository(db *gorm.DB) ActivityRepository {
	return &ActivityConnectDB{dbConnect: db} //connect database to interface
}

func (conn *ActivityConnectDB) All() []domain.Activity {
	var activity []domain.Activity
	conn.dbConnect.Find(&activity)
	return activity
}

func (conn *ActivityConnectDB) Create(activity domain.Activity) domain.Activity {
	conn.dbConnect.Save(&activity)
	return activity
}

func (conn *ActivityConnectDB) Update(activity domain.Activity) domain.Activity {
	conn.dbConnect.Omit("created_at").Save(&activity)
	return activity
}

func (conn *ActivityConnectDB) Delete(activity domain.Activity) {
	conn.dbConnect.Delete(&activity)
}

func (conn *ActivityConnectDB) FindById(id uint) (domain.Activity, error) {
	var activity domain.Activity
	conn.dbConnect.Find(&activity, "id = ?", id)
	if activity.ID == 0 {
		return activity, errors.New("id not found")
	}
	return activity, nil
}
