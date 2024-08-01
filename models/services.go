package models

import (
	"time"

	"example.com/car-rental/db"
)

type Service struct {
	ID        uint
	Name      string			`binding:"required"`
	Duration  int					`binding:"required"`
	Mileage   int					`binding:"required"`
	CarID     uint				`binding:"required"`
	UserID		uint
	CreatedAt time.Time
	UpdatedAt time.Time
}

func CreateServices(listServices []Service) error {
	result := db.DB.Model(&Service{}).Create(listServices)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func FindAllServices(carId uint) ([]Service, error) {
	var services []Service
	result := db.DB.Where(&Service{CarID: carId}).Order("name asc").Find(&services)
	if result.Error != nil {
		return nil, result.Error
	}
	return services, nil
}

func (s *Service) UpdateService() error {
	result := db.DB.Save(&s)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

