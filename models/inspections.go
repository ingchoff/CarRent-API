package models

import (
	"errors"
	"math"
	"time"

	"example.com/car-rental/db"
)

type Inspection struct {
	ID                uint
	InspectionDate    time.Time		`binding:"required"`
	Mileage           int					`binding:"required"`
	Amount						float64			`binding:"required"`
	Service						string			`binding:"required"`
	Description 			string			`binding:"required"`
	Name             	string			`binding:"required"`
	Duration					float64
	CarID          		uint				`binding:"required"`
	UserID						uint
	CreatedAt 				time.Time
	UpdatedAt 				time.Time
}

type LatestInspection struct {
	InspectionDate		time.Time
	Amount						float64
	Mileage						int
	Duration					float64
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
	result := db.DB.Where(&Inspection{CarID: carId}).Order("inspection_date asc").Find(&ins)
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

func FindInsByType(service string, cid string) ([]Inspection, error) {
	var inspection []Inspection
		if (service == "") {
			result := db.DB.Where("car_id = ?", cid).Order("inspection_date asc").Find(&inspection)
			if result.Error != nil {
				return inspection, result.Error
			}
			return inspection, nil
		} else {
			result := db.DB.Where("service = ? AND car_id = ?", service, cid).Order("inspection_date asc").Find(&inspection)
			if result.Error != nil {
				return inspection, result.Error
			}
			return inspection, nil
		}
		
}

func LatestInsByCar(cid string) (map[string]LatestInspection, error) {
	// var summary SummaryInspections
	var distincServices []string
	inspections := make(map[string]LatestInspection)
	result := db.DB.Table("inspections").Where("car_id = ?", cid).Distinct("Service").Order("Service asc").Find(&distincServices)
	if result.Error != nil {
		return nil, result.Error
	}
	if (len(distincServices) == 0) {
		return nil, errors.New("not found")
	}
	for _, service := range distincServices {
		var inspection LatestInspection
		result := db.DB.Table("inspections").Where("service = ?", service).Order("inspection_date desc").Find(&inspection)
		if result.Error != nil {
			return nil, result.Error
		}
		now := time.Now()
		elapsed := now.Sub(inspection.InspectionDate).Hours()/24
		duration := inspection.Duration*365
		percent := (elapsed/float64(duration))*100
		inspection.Duration = math.Round(percent*10) / 10
		inspections[service] = inspection
	}
	return inspections, nil
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
