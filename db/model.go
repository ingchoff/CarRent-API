package db

import (
	"fmt"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// User represents a User renting Cars.
type User struct {
	ID        uint        `gorm:"primaryKey"`
	FirstName string
	LastName  string
	Email     string
	Password	string
	Phone     string
	CreatedAt time.Time
	Role			string
	Rentals   []Rental    `gorm:"foreignKey:UserID"`
}

type RefreshToken struct {
	ID				uint
	UserID		uint
	User   		User        `gorm:"foreignKey:UserID"`
	Token			string
	Revorked	bool        `gorm:"default:false"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Car represents a Car available for rent.
type Car struct {
	ID        	    					uint      	`gorm:"primaryKey"`
	Model     								string
	SubModel									string
	Make      								string
	Year      								int
	Color     								string
	Gear											string
	Fuel											string
	Engine										string
	Image											string
	DailyRate 								float64
	License										string
	CarName										string
	Door											int
	LatestMileage							int					`gorm:"default:0"`
	LatestInspectionDate			time.Time
	Available 								bool				`gorm:"default:false"`
	UserID										uint
	User											User				`gorm:"foreignKey:UserID"`
	Rentals   								[]Rental		`gorm:"foreignKey:CarID"`
	CreatedAt 								time.Time
	UpdatedAt 								time.Time
}

// Rental represents the rental of a Car by a User.
type Rental struct {
	ID         		uint     		`gorm:"primaryKey"`
	Name					string
	Nid						string
	Phone					string
	StartDate 		time.Time
	EndDate 			time.Time
	StartMile			int
	EndMile				*int
	CustomerNote	string
	Detail				string
	Expense				float64
	TotalAmount		float64
	TotalNet  		*float64
	DailyRate			float64
	CarDelivery1	float64
	CarDelivery2	float64
	Status				string
	UserID				uint
	CarID  				uint
	User   				User    		`gorm:"foreignKey:UserID"`
	Car    				Car     		`gorm:"foreignKey:CarID"`
	CreatedAt 		time.Time
	UpdatedAt 		time.Time
}

// Inspection represents an inspection of a rented Car.
type Inspection struct {
	ID                uint      `gorm:"primaryKey"`
	InspectionDate    time.Time
	Mileage           int
	Amount						float64
	Service						string
	Description 			string
	Name             	string
	PercentDuration		float64
	PercentMileage		float64
	UserID						uint
	CarID          		uint
	Car            		Car 			`gorm:"foreignKey:CarID"`
	User   						User    	`gorm:"foreignKey:UserID"`
	CreatedAt 				time.Time
	UpdatedAt 				time.Time
}

type Service struct {
	ID					uint				`gorm:"primaryKey"`
	Name				string
	Duration		int
	Mileage			int
	CarID				uint
	UserID			uint
	Car					Car					`gorm:"foreignKey:CarID"`
	User				User				`gorm:"foreignKey:UserID"`
	CreatedAt 	time.Time
	UpdatedAt 	time.Time
}

var DB *gorm.DB

func InitDb() *gorm.DB {
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbPort, dbName)
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	DB.AutoMigrate(&User{}, &Car{}, &Rental{}, &Inspection{}, &RefreshToken{}, &Service{})
	return DB
}