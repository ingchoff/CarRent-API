package db

import (
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// User represents a User renting Cars.
type User struct {
	ID        uint      	`gorm:"primaryKey"`
	FirstName string
	LastName  string
	Email     string
	Password	string
	Phone     string
	CreatedAt time.Time
	Role			string
	Rentals   []Rental 		`gorm:"foreignKey:UserID"`
}

type RefreshToken struct {
	ID				uint
	UserID		uint
	User   		User    		`gorm:"foreignKey:UserID"`
	Token			string
	Revorked	bool				`gorm:"default:false"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Car represents a Car available for rent.
type Car struct {
	ID        uint      		`gorm:"primaryKey"`
	Model     string
	Make      string
	Year      int
	Color     string
	DailyRate float64
	Available bool					`gorm:"default:false"`
	UserID		uint
	User			User					`gorm:"foreignKey:UserID"`
	Rentals   []Rental			`gorm:"foreignKey:CarID"`
}

// Rental represents the rental of a Car by a User.
type Rental struct {
	ID         uint       `gorm:"primaryKey"`
	RentalDate time.Time
	ReturnDate *time.Time
	TotalCost  *float64
	UserID uint
	CarID  uint
	User   User    `gorm:"foreignKey:UserID"`
	Car    Car     `gorm:"foreignKey:CarID"`
}

// Inspection represents an inspection of a rented Car.
type Inspection struct {
	ID                uint      `gorm:"primaryKey"`
	InspectionDate    time.Time
	Mileage           int
	FuelLevel         float64
	DamageDescription string
	Notes             string
	CarID          		uint
	Car            		Car 		`gorm:"foreignKey:CarID"`
}

var DB *gorm.DB

func InitDb() {
	dsn := "Ing:Cr@753951@tcp(127.0.0.1:3306)/car_rent?charset=utf8mb4&parseTime=True&loc=Local"
  var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
    panic("failed to connect database")
  }
	DB.AutoMigrate(&User{}, &Car{}, &Rental{}, &Inspection{}, &RefreshToken{})
}