package models

import "gorm.io/gorm"

type Auth struct {
	ID       int    `gorm:"primaryKey" json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// CheckAuth checks if authentication information exists
func CheckAuth(username, password string) (bool, error) {
	var auth Auth

	err := db.Where(&Auth{Username: username, Password: password}).First(&auth).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if auth.ID > 0 {
		return true, nil
	}

	return false, nil
}

// ExistAuthByUsername check if a user exists based on username
func ExistAuthByUsername(username string) (bool, error) {
	var auth Auth
	err := db.Where(&Auth{Username: username}).First(&auth).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}
	if auth.ID > 0 {
		return true, nil
	}
	return false, nil
}

// GetUserByID get user infomation by id
func GetUserByID(id int) (*Auth, error) {
	var auth Auth
	err := db.Where(&Auth{ID: id}).First(&auth).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &auth, nil
}

// GetUserByUsername get user infomathion by username
func GetUserByUsername(username string) (*Auth, error) {
	var auth Auth
	err := db.Where(&Auth{Username: username}).First(&auth).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &auth, nil
}

// AddUser add a user record
func AddUser(username, password string) error {
	user := Auth{
		Username: username,
		Password: password,
	}

	if err := db.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

// EditUser modify password
func EditUser(id int, password string) error {
	if err := db.Model(&Auth{}).Where("id = ?", id).Update("password", password).Error; err != nil {
		return err
	}
	return nil
}
