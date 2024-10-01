package models

import (
	"time"

	"example.com/car-rental/db"
	"example.com/car-rental/utils"
)

type Rental struct {
	ID         		uint     		`gorm:"primaryKey"`
	StartDate 		time.Time		`binding:"required"`
	EndDate 			*time.Time
	StartMile			int					`binding:"required"`
	EndMile				*int
	CustomerNote	string			
	Detail				string
	Expense				float64
	TotalAmount		float64
	TotalNet  		*float64
	DailyRate			float64			`binding:"required"`
	CarDelivery1	float64
	CarDelivery2	float64
	Status				string			`binding:"required"`
	Name					string			
	Nid						string			
	Phone					string			
	UserID				uint				`binding:"required"`
	CarID  				uint				`binding:"required"`
	CreatedAt 		time.Time
	UpdatedAt 		time.Time
}

func (r *Rental) Save() error {
	result := db.DB.Create(&r)
	if result.Error != nil {
		return result.Error
	}
	//check updated date ล่าสุดก่อน update lastest mileage
	updateCarLatestMileage := db.DB.Model(&Car{}).Where("id = ?", r.CarID).Update("latest_mileage", r.StartMile)
	if updateCarLatestMileage.Error != nil {
		return updateCarLatestMileage.Error
	}
	return nil
}

func FindAllRentals(userId uint, carId uint) ([]Rental, error) {
	var rentals []Rental
	result := db.DB.Where(&Rental{UserID: userId, CarID: carId}).Find(&rentals)
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
	updateCarLatestMileage := db.DB.Model(&Car{}).Where("id = ?", c.CarID).Update("latest_mileage", c.EndMile)
	if updateCarLatestMileage.Error != nil {
		return updateCarLatestMileage.Error
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

func FindRentalByCondition(uid uint, name string, nid string, start string, end string, cid string) ([]Rental, error) {
	var rentals []Rental
	startDate := utils.ConvertUnixToDateTimeFormat(start)
	endDate := utils.ConvertUnixToDateTimeFormat(end)
	where := "(car_id = ? AND user_id = ?) AND (name = ? OR start_date = ? OR end_date = ?)"
	result := db.DB.Where(where, cid, uid, name, startDate, endDate).Find(&rentals)
	if result.Error != nil {
		return rentals, result.Error
	}
	return rentals, nil
}