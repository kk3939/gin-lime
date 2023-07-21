package userModel

import (
	"github.com/kk3939/gin-lime/db"
	"github.com/kk3939/gin-lime/entity"
)

func GetUsers() (entity.Users, error) {
	var users entity.Users
	db := db.GetDB()
	result := db.Find(&users)
	if err := result.Error; err != nil {
		return nil, err
	}
	return users, nil
}

func GetUserByEmailPassword(email string, password string) (*entity.User, error) {
	var user entity.User
	db := db.GetDB()
	result := db.Where(&entity.User{Email: email, Password: password}).First(&user)
	if err := result.Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUser(id string) (*entity.User, error) {
	var user entity.User
	db := db.GetDB()
	result := db.First(&user, "Id = ?", id)
	if err := result.Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func CreateUser(user *entity.User) error {
	db := db.GetDB()
	result := db.Select("Email", "Password").Create(&user)
	if err := result.Error; err != nil {
		return err
	}
	return nil
}

func UpdateUser(user *entity.User) error {
	db := db.GetDB()
	result := db.Model(&user).Updates(entity.User{Email: user.Email, Password: user.Password})
	if err := result.Error; err != nil {
		return err
	}
	return nil
}

func DeleteUser(user *entity.User) error {
	db := db.GetDB()
	result := db.Delete(&user)
	if err := result.Error; err != nil {
		return err
	}
	return nil
}
