package controllers

import (
	"github.com/ihakalanka/Assessment-Golang-REST-API/dtos"
	"github.com/ihakalanka/Assessment-Golang-REST-API/models"
	"github.com/ihakalanka/Assessment-Golang-REST-API/utils"
	"github.com/labstack/echo/v4"
	"net/http"
	"github.com/ihakalanka/Assessment-Golang-REST-API/services"
	"fmt"
	"strconv"
)

func RegisterUser(c echo.Context) error {
    var input dto.RegisterUserInput
    if err := c.Bind(&input); err != nil {
        return c.JSON(http.StatusBadRequest, utils.JSONResponse(http.StatusBadRequest, "Invalid input"))
    }

    validator := utils.NewValidator()
    if err := validator.Struct(input); err != nil {
        utils.Log.Warn("Invalid input")
        return c.JSON(http.StatusBadRequest, utils.JSONResponse(http.StatusBadRequest, fmt.Sprintf("Validation error: %s", err.Error())))
    }

    newUser, err := services.CreateUser(models.User{
        Name:     input.Name,
        Email:    input.Email,
        Password: input.Password,
        Role:     input.Role,
    })

    if err != nil {
        return c.JSON(http.StatusConflict, utils.JSONResponse(http.StatusConflict, err.Error()))
    }

    utils.Log.Infof("User %s registered successfully", newUser.Email)
    return c.JSON(http.StatusCreated, utils.JSONResponse(http.StatusCreated, newUser))
}

func GetAllUsers(c echo.Context) error {
	users, err := services.AllUsers()
	if err != nil {
		utils.Log.Errorf("Error retrieving users: %v", err)
		return c.JSON(http.StatusInternalServerError, utils.JSONResponse(http.StatusInternalServerError, "Failed to retrieve users"))
	}

	return c.JSON(http.StatusOK, utils.JSONResponse(http.StatusOK, users))
}

func DeleteUser(c echo.Context) error {
    userID := c.Param("id")

	if err := services.Delete(userID); err != nil {
		utils.Log.Warnf("Error deleting user with ID %s: %v", userID, err)
		return c.JSON(http.StatusInternalServerError, utils.JSONResponse(http.StatusInternalServerError, "Failed to delete user"))
	}

	utils.Log.Infof("User with ID %s deleted successfully", userID)
	return c.JSON(http.StatusOK, utils.JSONResponse(http.StatusOK, "User deleted successfully"))
}

func UpdateUser(c echo.Context) error {
    var input dto.UpdateUserInput
    if err := c.Bind(&input); err != nil {
        return c.JSON(http.StatusBadRequest, utils.JSONResponse(http.StatusBadRequest, "Invalid input"))
    }

    userID, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        return c.JSON(http.StatusBadRequest, utils.JSONResponse(http.StatusBadRequest, "Invalid user ID"))
    }

    updatedUser, err := services.Update(models.User{
        ID:       uint(userID),
        Name:     input.Name,
        Email:    input.Email,
        Password: input.Password,
        Role:     input.Role,
    })

    if err != nil {
        utils.Log.Errorf("Error updating user %d: %v", userID, err)
        return c.JSON(http.StatusInternalServerError, utils.JSONResponse(http.StatusInternalServerError, "Failed to update user"))
    }

    utils.Log.Infof("User %s updated successfully", updatedUser.Email)
    return c.JSON(http.StatusOK, utils.JSONResponse(http.StatusOK, updatedUser))
}