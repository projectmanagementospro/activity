package repository

import (
	"activity/models/domain"
	"errors"

	"gorm.io/gorm"
)

type SubActivityRepository interface {
	All() []domain.SubActivity
	Create(subactivity domain.SubActivity) domain.SubActivity
	Update(subactivity domain.SubActivity) domain.SubActivity
	Delete(subactivity domain.SubActivity)
	FindById(id uint) (domain.SubActivity, error)
	IsDReportExist(id uint) (domain.Activity, error)
}

type SubActivityConnectDB struct {
	dbConnect *gorm.DB //connect to database
}

func NewSubActivityRepository(db *gorm.DB) SubActivityRepository {
	return &SubActivityConnectDB{dbConnect: db} //connect database to interface
}

func (conn *SubActivityConnectDB) All() []domain.SubActivity {
	var subactivity []domain.SubActivity
	conn.dbConnect.Find(&subactivity)
	return subactivity
}

func (conn *SubActivityConnectDB) Create(subactivity domain.SubActivity) domain.SubActivity {
	conn.dbConnect.Save(&subactivity)
	return subactivity
}

func (conn *SubActivityConnectDB) Update(subactivity domain.SubActivity) domain.SubActivity {
	conn.dbConnect.Omit("created_at").Save(&subactivity)
	return subactivity
}

func (conn *SubActivityConnectDB) Delete(subactivity domain.SubActivity) {
	conn.dbConnect.Delete(&subactivity)
}

func (conn *SubActivityConnectDB) FindById(id uint) (domain.SubActivity, error) {
	var subactivity domain.SubActivity
	conn.dbConnect.Preload("Activity").Find(&subactivity, "id = ?", id)
	if subactivity.ID == 0 {
		return subactivity, errors.New("id not found")
	}
	return subactivity, nil
}
func (conn *SubActivityConnectDB) IsDReportExist(id uint) (domain.Activity, error) {
	var subactivity domain.Activity
	conn.dbConnect.Preload("Activity").Find(&subactivity, "id = ?", id)
	if subactivity.ID == 0 {
		return subactivity, errors.New("DailyReport id haven't been created yet")
	}
	return subactivity, nil
}
