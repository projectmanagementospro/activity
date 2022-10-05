package repository

import (
	"activity/models/domain"
	"errors"

	"gorm.io/gorm"
)

type ActivityRepository interface {
	All() []domain.Activity
	Create(a domain.Activity) domain.Activity
	Update(a domain.Activity) domain.Activity
	Delete(a domain.Activity)
	FindById(id uint) (domain.Activity, error)
}

type ActivityConnection struct {
	connection *gorm.DB
}

func NewActivityRepository(connection *gorm.DB) ActivityRepository {
	return &ActivityConnection{connection: connection}
}

func (c *ActivityConnection) All() []domain.Activity {
	var activity []domain.Activity
	c.connection.Find(&activity)
	return activity
}

func (c *ActivityConnection) Create(a domain.Activity) domain.Activity {
	c.connection.Save(&a)
	return a
}

func (c *ActivityConnection) Update(a domain.Activity) domain.Activity {
	c.connection.Omit("created_at").Save(&a)
	return a
}

func (c *ActivityConnection) Delete(a domain.Activity) {
	c.connection.Delete(&a)
}

func (c *ActivityConnection) FindById(id uint) (domain.Activity, error) {
	var activity domain.Activity
	c.connection.Find(&activity, "id = ?", id)
	if activity.ID == 0 {
		return activity, errors.New("id not found")
	}
	return activity, nil
}
