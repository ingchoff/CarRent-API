package models

import (
	"fmt"
	"time"

	"example.com/car-rental/db"
)

type Car struct {
	ID        uint
	Model     string		`binding:"required"`
	SubModel	string		`binding:"required"`
	Make      string		`binding:"required"`
	Year      int				`binding:"required"`
	Color     string		`binding:"required"`
	Gear			string		`binding:"required"`
	Fuel			string		`binding:"required"`
	Engine		string		`binding:"required"`
	Image			string		`binding:"required"`
	DailyRate float64		`binding:"required"`
	License		string		`binding:"required"`
	Door			int				`binding:"required"`
	Available bool
	UserID		uint
	CreatedAt time.Time
	UpdatedAt time.Time
}

type ModelsData struct {
	Make	string
	Model	string
}

type objModels struct {
	Model	map[string][]string
}

type ObjMakes struct {
	Make	map[string][]string
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

func DistinctModelNames(uid uint) (ObjMakes, error) {
	var results []ModelsData
	objMakes := ObjMakes{Make: make(map[string][]string)}
	result := db.DB.Table("cars").Where("user_id = ?", uid).Distinct("Make", "Model").Order("Make, Model asc").Find(&results)
	if result.Error != nil {
		return objMakes, result.Error
	}
	for _, value := range results {
		objMakes.Make[value.Make] = append(objMakes.Make[value.Make], value.Model)
		// objModels.Model[value.Model] = append(objModels.Model[value.Model], value.Model)
	}
	return objMakes, nil
}

func FindCarByCondition(uid uint, condition string, value string) ([]Car, error) {
	var car []Car
	where := fmt.Sprintf("user_id = ? AND %s = ?", condition)
	result := db.DB.Where(where, uid, value).Find(&car)
	if result.Error != nil {
		return car, result.Error
	}
	return car, nil
}
