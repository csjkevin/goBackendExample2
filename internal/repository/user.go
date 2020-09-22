package repository

import (
	"log"
	"goBackendExample/internal/model"
)

func CreateUser(user model.User) error {
	query := db.Create(&user)
	err := query.Error
	if err != nil {
		log.Panic(err)
		return err
	}
	return nil
}

func GetAllUser() ([]model.User, error) {
	var users []model.User
	query := db.Order("id").Find(&users)
	err := query.Error
	if err != nil {
		log.Panic(err)
		return nil, err
	}
	return users, nil
}

func GetUserById(id uint) (*model.User, error) {
	var user model.User
	query := db.First(&user, id)
	err := query.Error
	if err != nil {
		log.Panic(err)
		return nil, err
	}
	return &user, nil
}

func UpdateUser(id uint, user model.User) error {
	user.ID = id
	query := db.Model(&user).Updates(user)
	err := query.Error
	if err != nil {
		log.Panic(err)
		return err
	}
	return nil
}