package service

import (
	"goBackendExample/internal/model"
	"goBackendExample/internal/repository"
)

func CreateUser(user model.User) error {
	err := repository.CreateUser(user)
	if err != nil {
		return err
	}
	return nil
}

func GetAllUser() ([]model.User, error) {
	users, err := repository.GetAllUser()
	if err != nil {
		return nil, err
	}

	return users, nil
}

func GetUserById(id int) (*model.User, error) {
	user, err := repository.GetUserById(uint(id))
	if err != nil {
		return nil, err
	}

	return user, nil
}

func UpdateUser(id int, user model.User) error {
	err := repository.UpdateUser(uint(id), user)
	if err != nil {
		return err
	}
	return nil
}
