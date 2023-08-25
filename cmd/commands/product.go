package commands

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"platform/internal/models"
	"platform/internal/validator"
)

type ProductManager struct {
	Products    models.Product
	TimeManager *TimeManager
}
type ProductManagerInterface interface {
	GetProductInfo(code string) (*models.Product, error)
	CreateProduct(code string, price float64, stock int) (*models.Product, error)
}

func (m *ProductManager) GetProductInfo(code string) (*models.Product, error) {
	product, err := m.Products.Find("code", code)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, GetErrorMessage("ErrProductNotFound")
		}
		return nil, GetErrorMessage("ErrProductQuery")
	}

	return &product, nil
}

func (m *ProductManager) CreateProduct(code string, price float64, stock int) (*models.Product, error) {
	product := &models.Product{Code: code, Price: price, Stock: stock}

	v := validator.New()

	models.ValidateProduct(v, product)
	if !v.Valid() {
		return nil, errors.New(fmt.Sprint(v.Errors))
	}
	createdProduct, err := m.Products.Create()
	if err != nil {
		return nil, GetErrorMessage("ErrProductCreate")
	}
	return &createdProduct, nil
}
