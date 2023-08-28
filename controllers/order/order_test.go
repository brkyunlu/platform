package order_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"platform/controllers/order"
	"platform/internal/models"
	"platform/internal/testutils"
)

func TestOrderManager_CreateOrder(t *testing.T) {
	testutils.SetupTestDB()

	mockOrderManager := order.OrderManager{
		Order: models.Order{},
	}

	product := models.Product{
		Code:  "PC01",
		Price: float64(499),
		Stock: 100,
	}
	newProduct, err := product.Create()
	assert.NoError(t, err)
	campaign := models.Campaign{
		Name:                     "CO01",
		ProductID:                newProduct.ID,
		Duration:                 7,
		TargetSales:              100,
		Status:                   true,
		PriceManipulationLimit:   0.5,
		CurrentPriceManipulation: 0.4,
		Expiry:                   time.Now().AddDate(0, 0, -3), // Expired 3 days ago
	}
	_, err = campaign.Create()
	assert.NoError(t, err)

	createdOrder, err := mockOrderManager.CreateOrder("P01", 2)
	assert.NoError(t, err)
	assert.NotNil(t, createdOrder)

	insufficientStockOrder, err := mockOrderManager.CreateOrder("P01", 200)
	assert.Error(t, err)
	assert.Nil(t, insufficientStockOrder)
}
