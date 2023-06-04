package service

import "github.com/jnhu76/dwz/models"

type Auth struct {
	ID       int
	Username string
	Password string
}

func (a *Auth) ExistByUsername() (bool, error) {
	return models.ExistAuthByUsername(a.Username)
}

func (a *Auth) Add() error {
	return models.AddUser(a.Username, a.Password)
}

func (a *Auth) ChangePassword() error {
	return models.EditUser(a.ID, a.Password)
}

func (a *Auth) Check() (bool, error) {
	return models.CheckAuth(a.Username, a.Password)
}

func (a *Auth) Get(id int) (*models.Auth, error) {
	user, err := models.GetUserByID(a.ID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (a *Auth) GetUserByUsername() (*models.Auth, error) {
	user, err := models.GetUserByUsername(a.Username)

	if err != nil {
		return nil, err
	}
	return user, nil
}
