package order

import (
	"gorm.io/gorm"
	"platform/internal/error_handler"
	"platform/internal/logger"
	"platform/internal/models"
)

type OrderManager struct {
	Order models.Order
}

func (m *OrderManager) CreateOrder(productCode string, quantity int) (*models.Order, error) {
	product, err := m.Order.Product.Find("code", productCode)
	if err != nil {
		return nil, error_handler.GetErrorMessage("ErrProductNotFound", logger.Error)
	}

	campaign, campaignErr := models.Campaign{}.Find("product_id", product.ID)
	if campaignErr == nil || campaignErr == gorm.ErrRecordNotFound {
		discountedPrice := product.Price * (1 - (campaign.CurrentPriceManipulation / 100))
		product.Price = discountedPrice
	} else {
		return nil, error_handler.GetErrorMessage("ErrCampaignNotFound", logger.Error)
	}

	if quantity > product.Stock {
		return nil, error_handler.GetErrorMessage("ErrNoStock", logger.Error)
	}

	totalPrice := product.Price * float64(quantity)
	order := models.Order{
		ProductID:  product.ID,
		CampaignID: &campaign.ID,
		Quantity:   quantity,
		TotalPrice: totalPrice,
	}

	_, err = order.Create()
	if err != nil {
		return nil, error_handler.GetErrorMessage("ErrOrderCreate", logger.Error)
	}

	updatedStock := product.Stock - quantity
	err = product.Update("stock", updatedStock)
	if err != nil {
		return nil, error_handler.GetErrorMessage("ErrStockUpdate", logger.Error)
	}

	return &order, nil
}
