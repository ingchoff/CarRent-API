package models

import (
	"errors"
	"fmt"
	"time"

	"example.com/car-rental/db"
	"example.com/car-rental/utils"
)

type User struct {
	ID        uint      
	FirstName string
	LastName  string
	Email     string	`binding:"required"`
	Password	string	`binding:"required"`
	Phone     string
	Role			string
	CreatedAt time.Time
}

type RefreshToken struct {
	ID				uint
	UserID		uint
	Token			string
	User   		User
	Revorked	bool			`gorm:"default:false"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u *User) Save() (uint, error) {
	hashPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return 0, err
	}
	u.Password = hashPassword
	result := db.DB.Table("users").Create(&u)
	if result.Error != nil {
		return 0, result.Error
	}
	return u.ID, nil
}

func GetUserById(id uint) (User, error) {
	var user User
	result := db.DB.Table("users").Select([]string{"id", "FirstName", "LastName", "Email", "Phone", "Role"}).Where("id = ?", id).Find(&user)
	if result.Error != nil {
		return user, result.Error
	}
	return user, nil
} 

func (u *User) ValidateCredentials() (User, error) {
	var user User
	result := db.DB.Table("users").Where("email = ?", u.Email).First(&user)
	if result.Error != nil {
		return user, result.Error
	}
	passwordIsValid := utils.CheckPasswordHash(u.Password, user.Password)
	if !passwordIsValid {
		return user, errors.New("credentials invalid")
	}
	return user, nil
}

func AddRefreshTokenToWhitelist(user User, token string) error {
	var obj RefreshToken
	var listObj []RefreshToken
	var err error
	obj.Token = token
	obj.User = user
	fmt.Println(obj)
	query := db.DB.Where("user_id = ? AND Revorked = ?", user.ID, false).Find(&listObj)
	if query.RowsAffected >= 1 {
		revork := db.DB.Where("user_id = ? AND Revorked = ?", user.ID, false).Updates(RefreshToken{Revorked: true})
		db.DB.Create(&obj)
		if revork.Error != nil {
			err = revork.Error
		}
	} else {
		result := db.DB.Create(&obj)
		if result.Error != nil {
			err = result.Error
		}
	}
	return err
}

func RevorkToken(token string) error {
	query := db.DB.Model(&RefreshToken{}).Where("token = ?", token).Update("Revorked", true)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func GetUserByToken(token string) (User, error) {
	var obj RefreshToken
	query := db.DB.Where("token = ? AND Revorked = ?", token, false).First(&obj)
	if query.Error != nil {
		return User{}, query.Error
	}
	user, err := GetUserById(obj.UserID)
	if err != nil {
		return User{}, err
	}
	return user, nil
}