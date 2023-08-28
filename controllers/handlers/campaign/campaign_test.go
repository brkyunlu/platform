package campaign_test

import (
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"platform/controllers/handlers/campaign"
	"platform/internal/models"
	"platform/internal/testutils"
	"testing"
	"time"
)

func TestGetCampaignInfoHandler(t *testing.T) {
	testutils.SetupTestDB()

	newProduct := models.Product{
		Code:  "PH06",
		Price: 29.99,
		Stock: 100,
	}

	createdProduct, err := newProduct.Create()
	assert.NoError(t, err)
	newCampaign := models.Campaign{
		Name:                     "CH05",
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

	// Burada test için gerekli olan veritabanı veya sahte veri oluşturulabilir
	cmd := &cobra.Command{}
	args := []string{newCampaign.Name}

	// Test case 1: Ürün bilgisini alırken başarılı bir durum
	err = campaign.GetCampaignInfoHandler(cmd, args)
	assert.NoError(t, err)

	// Test case 2: Ürün kodu eksik olduğunda hata dönmeli
	err = campaign.GetCampaignInfoHandler(nil, []string{})
	assert.Error(t, err)
}
