package services

import (
	"go-mysql-phpmyadmin/db"
	"go-mysql-phpmyadmin/models"
)

func CreateEvent(event *models.Event) error {
	return db.DB.Create(event).Error
}

func GetAllEvents(page, pageSize int) ([]models.Event, error) {
	var events []models.Event
	offset := (page - 1) * pageSize

	err := db.DB.
		Order("created_at desc").
		Limit(pageSize).
		Offset(offset).
		Find(&events).
		Error

	return events, err
}
