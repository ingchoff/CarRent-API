package models

import (
	"errors"
	"math"
	"strconv"
	"time"

	"example.com/car-rental/db"
)

type Inspection struct {
	ID                uint
	InspectionDate    time.Time		`binding:"required"`
	Mileage           int
	Amount						float64
	Service						string			`binding:"required"`
	Description 			string			`binding:"required"`
	Name             	string			`binding:"required"`
	PercentDuration		float64
	PercentMileage		float64
	CarID          		uint				`binding:"required"`
	UserID						uint
	CreatedAt 				time.Time
	UpdatedAt 				time.Time
}

type LatestInspection struct {
	InspectionDate		time.Time
	PercentDuration		float64
	Mileage						int
	PercentMileage		float64
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
	cidsd, _ := strconv.ParseInt(cid, 10, 32)
	car, err := FindCarById(uint(cidsd))
	if err != nil {
		return nil, err
	}
	for _, service := range distincServices {
		var serviceInfo Service
		var latestInspection LatestInspection
		services := db.DB.Table("services").Where("name = ? AND car_id = ?", service, cid).Order("name asc").Find(&serviceInfo)
		if services.Error != nil {
			return nil, result.Error
		}
		inspection := db.DB.Table("inspections").Where("service = ? AND car_id = ?", service, cid).Order("inspection_date desc").Find(&latestInspection)
		if inspection.Error != nil {
			return nil, result.Error
		}
		now := time.Now()
		elapsed := now.Sub(latestInspection.InspectionDate).Hours()/24
		duration := serviceInfo.Duration*365
		mileage := serviceInfo.Mileage
		var latestMileage int
		if (car.LatestMileage == 0) {
			latestMileage = latestInspection.Mileage
		} else {
			latestMileage = car.LatestMileage
		}
		elapsedMileage := latestMileage-latestInspection.Mileage
		percentMileage := (float64(elapsedMileage)/float64(mileage))
		percentDuration := (elapsed/float64(duration))*100
		latestInspection.PercentDuration = math.Round(percentDuration*10) / 10
		latestInspection.PercentMileage = math.Round(percentMileage*100)
		inspections[service] = latestInspection
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
