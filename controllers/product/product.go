package product

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"platform/internal/error_handler"
	"platform/internal/logger"
	"platform/internal/models"
	time_simulation "platform/internal/time"
	"platform/internal/validator"
)

type ProductManager struct {
	Products    models.Product
	TimeManager time_simulation.DefaultTimeSimulator
}

type IProductManager interface {
	GetProductInfo(code string) (*models.Product, error)
	CreateProduct(code string, price float64, stock int) (*models.Product, error)
}

func (m *ProductManager) GetProductInfo(code string) (*models.Product, error) {
	product, err := m.Products.Find("code", code)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, error_handler.GetErrorMessage("ErrProductNotFound", logger.Error)
		}
		return nil, error_handler.GetErrorMessage("ErrProductQuery", logger.Error)
	}

	campaigns, err := m.Products.GetActiveCampaignsForProduct(product.ID, m.TimeManager.GetSimulatedTime())
	if err != nil {
		return nil, err
	}
	if len(campaigns) > 0 {
		currentPriceManipulation := campaigns[0].CurrentPriceManipulation
		product.Price *= (1 - currentPriceManipulation/100)
	}
	return &product, nil
}

func (m *ProductManager) CreateProduct(code string, price float64, stock int) (*models.Product, error) {
	m.Products = models.Product{Code: code, Price: price, Stock: stock}

	v := validator.New()

	models.ValidateProduct(v, &m.Products)
	if !v.Valid() {
		return nil, errors.New(fmt.Sprint(v.Errors))
	}
	createdProduct, err := m.Products.Create()
	if err != nil {
		return nil, error_handler.GetErrorMessage("ErrProductCreate", logger.Error)
	}
	return &createdProduct, nil
}
