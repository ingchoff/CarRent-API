package models

import (
	"time"

	"example.com/car-rental/db"
)

type Inspection struct {
	ID                uint
	InspectionDate    time.Time		`binding:"required"`
	Mileage           int					`binding:"required"`
	FuelLevel         float64
	DamageDescription string
	Notes             string
	CarID             uint				`binding:"required"`
}

type CarIdGetRequestBody struct {
	carId		uint	`binding:"required"`
}

func (i *Inspection) Save() error {
	result := db.DB.Create(&i)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func FindAllInspections(carId uint) ([]Inspection, error) {
	var ins []Inspection
	result := db.DB.Where(&Inspection{CarID: carId}).Find(&ins)
	if result.Error != nil {
		return nil, result.Error
	}
	return ins, nil
}

func FindInsById(id uint) (Inspection, error) {
	var inspection Inspection
	result := db.DB.First(&inspection, id)
	if result.Error != nil {
		return Inspection{}, result.Error
	}
	return inspection, nil
}

func (i *Inspection) UpdateIns() error {
	result := db.DB.Save(&i)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func DeleteInsById(id uint) (error) {
	result := db.DB.Delete(&Inspection{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
