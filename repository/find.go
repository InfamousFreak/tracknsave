package repository

import (
	"errors"
	"tracknsave/models"
)

func Find(email, password string) (*models.User, error){
	//query the database for the registered user
	if email== "test@gmail.com" && password == "test12345" {
		return &models.User{
			ID: 		1,
			Email:		"test@gmail.com",
			Password:	"test12345",
		}, nil
	}
	return nil, errors.New("user not found")
}