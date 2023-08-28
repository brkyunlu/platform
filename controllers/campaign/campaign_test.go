package campaign_test

import (
	"platform/controllers/campaign"
	time_simulated "platform/internal/time"

	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"platform/internal/models"
	"platform/internal/testutils"
)

func TestCampaignManager_GetCampaignInfo(t *testing.T) {
	testutils.SetupTestDB()
	mockCampaignManager := campaign.CampaignManager{
		Campaigns: models.Campaign{},
	}
	product := models.Product{
		Code:  "PC01",
		Price: float64(499),
		Stock: 100,
	}
	newProduct, err := product.Create()
	assert.NoError(t, err)

	newCampaign := models.Campaign{
		Name:      "SummerSale",
		ProductID: newProduct.ID,
		Duration:  7,
		Status:    true,
		Expiry:    time.Now().AddDate(0, 0, 7), // Expires in 7 days
	}

	_, err = newCampaign.Create()
	assert.NoError(t, err)
	campaignInfo, err := mockCampaignManager.GetCampaignInfo(newCampaign.Name)
	assert.NoError(t, err)
	assert.NotNil(t, campaignInfo)

	campaignInfo, err = mockCampaignManager.GetCampaignInfo("non_existing_campaign")
	assert.Error(t, err)
	assert.Nil(t, campaignInfo)
}

func TestCampaignManager_CreateCampaign(t *testing.T) {
	testutils.SetupTestDB()

	mockCampaignManager := campaign.CampaignManager{
		Campaigns:   models.Campaign{},
		TimeManager: time_simulated.DefaultTimeSimulator{},
	}

	product := models.Product{
		Code:  "PC02",
		Price: 29.99,
		Stock: 100,
	}

	_, err := product.Create()
	assert.NoError(t, err)

	createdCampaign, err := mockCampaignManager.CreateCampaign("NewCampaign", "PC02", 7, 0.1, 50)
	assert.NoError(t, err)
	assert.NotNil(t, createdCampaign)

	invalidCampaignManager := campaign.CampaignManager{
		Campaigns:   models.Campaign{},
		TimeManager: time_simulated.DefaultTimeSimulator{},
	}

	invalidCreatedCampaign, err := invalidCampaignManager.CreateCampaign("InvalidCampaign", "non_existing_product", 7, 0.1, 50)
	assert.Error(t, err)
	assert.Nil(t, invalidCreatedCampaign)
}

func TestCampaignManager_UpdateCampaignStatus(t *testing.T) {
	testutils.SetupTestDB()
	product := models.Product{
		Code:  "PC03",
		Price: 29.99,
		Stock: 100,
	}

	newProduct, err := product.Create()
	assert.NoError(t, err)

	mockCampaignManager := campaign.CampaignManager{
		Campaigns: models.Campaign{
			Name:                     "C01",
			ProductID:                newProduct.ID,
			Duration:                 7,
			TargetSales:              100,
			Status:                   true,
			PriceManipulationLimit:   0.5,
			CurrentPriceManipulation: 0.4,
			Expiry:                   time.Now().AddDate(0, 0, -3), // Expired 3 days ago
		},
	}
	newCampaign, err := mockCampaignManager.Campaigns.Create()
	assert.NoError(t, err)

	err = mockCampaignManager.UpdateCampaignStatus(&newCampaign)
	assert.NoError(t, err)
}
