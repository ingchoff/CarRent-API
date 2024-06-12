package models

import (
	"time"

	"example.com/car-rental/db"
)

type Car struct {
	ID        uint
	Model     string	`binding:"required"`
	Make      string	`binding:"required"`
	Year      int			`binding:"required"`
	Color     string	`binding:"required"`
	DailyRate float64	`binding:"required"`
	Available bool
	UserID		uint
}

type Rental struct {
	ID         uint `gorm:"primaryKey"`
	RentalDate time.Time
	ReturnDate *time.Time
	TotalCost  *float64
	UserID     uint
	User       User `gorm:"foreignKey:UserID"`
}

func (c *Car) Save() error {
	result := db.DB.Create(&c)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func FindAllCars(userId uint) ([]Car, error) {
	var cars []Car
	result := db.DB.Where(&Car{UserID: userId}).Find(&cars)
	if result.Error != nil {
		return nil, result.Error
	}
	return cars, nil
}

func FindCarById(cid uint) (Car, error) {
	var car Car
	result := db.DB.First(&car, cid)
	if result.Error != nil {
		return Car{}, result.Error
	}
	return car, nil
}

func (c *Car) UpdateCar() error {
	result := db.DB.Save(&c)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func DeleteCarById(cid uint) (error) {
	result := db.DB.Delete(&Car{}, cid)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
