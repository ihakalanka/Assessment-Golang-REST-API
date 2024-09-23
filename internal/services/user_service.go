package services

import (
	"errors"
	"fmt"
	"github.com/ihakalanka/Assessment-Golang-REST-API/internal/config"
	dto "github.com/ihakalanka/Assessment-Golang-REST-API/internal/dtos"
	"github.com/ihakalanka/Assessment-Golang-REST-API/internal/models"
	"github.com/ihakalanka/Assessment-Golang-REST-API/pkg"
)

func CreateUser(user models.User) (models.User, error) {
	var existingUser models.User
	if err := config.DB.Where("email = ?", user.Email).First(&existingUser).Error; err == nil {
		return models.User{}, errors.New("user already exists with this email")
	}

	hashedPassword, err := pkg.HashPassword(user.Password)
	if err != nil {
		return models.User{}, fmt.Errorf("error hashing password: %w", err)
	}
	user.Password = hashedPassword

	if user.Role == "" {
		user.Role = "user"
	}

	if err := config.DB.Create(&user).Error; err != nil {
		return models.User{}, fmt.Errorf("error creating user: %w", err)
	}
	return user, nil
}

func AllUsers() ([]dto.UserDTO, error) {
	var users []models.User
	if err := config.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	var userDTOs []dto.UserDTO
	for _, user := range users {
		userDTOs = append(userDTOs, dto.UserDTO{
			ID:       user.ID,
			Name:     user.Name,
			Email:    user.Email,
			Role:     user.Role,
		})
	}
	return userDTOs, nil
}

func Delete(userID string) error {
	var user models.User
	if err := config.DB.Where("id = ?", userID).First(&user).Error; err != nil {
		return err
	}
	if err := config.DB.Delete(&user).Error; err != nil {
		return err
	}
	return nil
}

func Update(user models.User) (models.User, error) {
    var existingUser models.User

    if err := config.DB.Where("id = ?", user.ID).First(&existingUser).Error; err != nil {
        return models.User{}, err
    }

    if user.Password != "" {
        hashedPassword, err := pkg.HashPassword(user.Password)
        if err != nil {
            return models.User{}, fmt.Errorf("error hashing password: %w", err)
        }
        user.Password = hashedPassword
    } else {
        user.Password = existingUser.Password
    }

    if err := config.DB.Model(&existingUser).Updates(user).Error; err != nil {
        return models.User{}, err
    }

    return existingUser, nil
}