package services

import (
	"github.com/ihakalanka/Assessment-Golang-REST-API/internal/config"
	"github.com/ihakalanka/Assessment-Golang-REST-API/internal/models"
)

func CreateProduct(product models.Product) (models.Product, error) {
	if err := config.DB.Create(&product).Error; err != nil {
		return models.Product{}, err
	}
	return product, nil
}

func UpdateProduct(product models.Product) (models.Product, error) {
	if err := config.DB.Save(&product).Error; err != nil {
		return models.Product{}, err
	}
	return product, nil
}

func DeleteProduct(productID string) error {
	var product models.Product
	if err := config.DB.Where("id = ?", productID).First(&product).Error; err != nil {
		return err
	}
	if err := config.DB.Delete(&product).Error; err != nil {
		return err
	}
	return nil
}

func GetProduct(productID string) (models.Product, error) {
	var product models.Product
	if err := config.DB.Where("id = ?", productID).First(&product).Error; err != nil {
		return models.Product{}, err
	}
	return product, nil
}

func GetProducts() ([]models.Product, error) {
	var products []models.Product
	if err := config.DB.Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}