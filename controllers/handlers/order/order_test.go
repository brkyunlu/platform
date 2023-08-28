package order_test

import (
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"platform/controllers/handlers/order"
	"platform/internal/models"
	"platform/internal/testutils"
	"testing"
	"time"
)

func TestGetOrderHandler(t *testing.T) {
	testutils.SetupTestDB()

	newProduct := models.Product{
		Code:  "PH05",
		Price: 29.99,
		Stock: 100,
	}

	createdProduct, err := newProduct.Create()
	assert.NoError(t, err)
	newCampaign := models.Campaign{
		Name:                     "CH03",
		ProductID:                createdProduct.ID,
		Duration:                 7,
		TargetSales:              100,
		Status:                   true,
		PriceManipulationLimit:   0.5,
		CurrentPriceManipulation: 0.4,
		Expiry:                   time.Now().AddDate(0, 0, +3), // Expired 3 days ago
	}
	_, err = newCampaign.Create()
	assert.NoError(t, err)

	cmd := &cobra.Command{}
	args := []string{newProduct.Code, "1"}

	err = order.OrderHandler(cmd, args)
	assert.NoError(t, err)

	err = order.OrderHandler(nil, []string{})
	assert.Error(t, err)
}
