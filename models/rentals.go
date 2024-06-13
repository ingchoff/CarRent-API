package models

import (
	"time"

	"example.com/car-rental/db"
)

type Rental struct {
	ID         uint `gorm:"primaryKey"`
	RentalDate time.Time
	ReturnDate *time.Time
	TotalCost  *float64
	UserID     uint
	CarID      uint
}

func (r *Rental) Save() error {
	result := db.DB.Create(&r)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func FindAllRentals(userId uint) ([]Rental, error) {
	var rentals []Rental
	result := db.DB.Where(&Rental{UserID: userId}).Find(&rentals)
	if result.Error != nil {
		return nil, result.Error
	}
	return rentals, nil
}

func FindRentalById(rid uint) (Rental, error) {
	var rental Rental
	result := db.DB.First(&rental, rid)
	if result.Error != nil {
		return Rental{}, result.Error
	}
	return rental, nil
}

func (c *Rental) UpdateRental() error {
	result := db.DB.Save(&c)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func DeleteRentalById(cid uint) (error) {
	result := db.DB.Delete(&Rental{}, cid)
	if result.Error != nil {
		return result.Error
	}
	return nil
}