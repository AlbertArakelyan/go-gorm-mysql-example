package services

import (
	"go-mysql-phpmyadmin/db"
	"go-mysql-phpmyadmin/models"
)

func CreateRegistration(eventID int64) error {
	return db.DB.Create(&models.Registration{EventID: eventID}).Error
}

func GetAllRegistrations(page, pageSize int) ([]models.Registration, error) {
	var registrations []models.Registration
	offset := (page - 1) * pageSize

	// db.Preload("Orders", "state NOT IN (?)", "cancelled").Find(&users) // we can have condition for preloading

	err := db.DB.
		Preload("Event"). // when not preloading it still brings the object but without values in it all IDs and dates are 0
		// solution just make a pointer the relation type in the model(struct) of Registration e.g. Event -> *Event
		Order("created_at desc").
		Limit(pageSize).
		Offset(offset).
		Find(&registrations).
		Error

	return registrations, err
}
