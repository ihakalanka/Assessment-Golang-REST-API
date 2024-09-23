package controllers

import (
	"net/http"
	"strconv"
	"github.com/ihakalanka/Assessment-Golang-REST-API/internal/dtos"
	"github.com/ihakalanka/Assessment-Golang-REST-API/internal/models"
	"github.com/ihakalanka/Assessment-Golang-REST-API/internal/services"
	"github.com/ihakalanka/Assessment-Golang-REST-API/pkg"
	"github.com/labstack/echo/v4"
)

func GetProducts(c echo.Context) error {
	products, err := services.GetProducts()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, pkg.JSONResponse(http.StatusInternalServerError, err.Error()))
	}

	return c.JSON(http.StatusOK, pkg.JSONResponse(http.StatusOK, products))
}

func GetProduct(c echo.Context) error {
	productID := c.Param("id")
	product, err := services.GetProduct(productID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, pkg.JSONResponse(http.StatusInternalServerError, err.Error()))
	}
	return c.JSON(http.StatusOK, pkg.JSONResponse(http.StatusOK, product))
}

func CreateProduct(c echo.Context) error {
	var input dto.ProductInput
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, pkg.JSONResponse(http.StatusBadRequest, err.Error()))
	}

	product := models.Product{
		Name:        input.Name,
		Price:       input.Price,
		Description: input.Description,
	}

	newProduct, err := services.CreateProduct(product)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, pkg.JSONResponse(http.StatusInternalServerError, err.Error()))
	}
	return c.JSON(http.StatusCreated, pkg.JSONResponse(http.StatusCreated, newProduct))
}

func UpdateProduct(c echo.Context) error {
	var input dto.ProductInput
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, pkg.JSONResponse(http.StatusBadRequest, err.Error()))
	}

	productID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, pkg.JSONResponse(http.StatusBadRequest, err.Error()))
	}

	product := models.Product{
		ID:          uint(productID),
		Name:        input.Name,
		Price:       input.Price,
		Description: input.Description,
	}

	updatedProduct, err := services.UpdateProduct(product)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, pkg.JSONResponse(http.StatusInternalServerError, err.Error()))
	}
	return c.JSON(http.StatusOK, pkg.JSONResponse(http.StatusOK, updatedProduct))
}

func DeleteProduct(c echo.Context) error {
	productID := c.Param("id")
	err := services.DeleteProduct(productID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, pkg.JSONResponse(http.StatusInternalServerError, err.Error()))
	}
	return c.JSON(http.StatusOK, pkg.JSONResponse(http.StatusOK, "Product deleted successfully"))
}