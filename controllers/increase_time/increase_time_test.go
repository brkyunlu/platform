package increase_time_test

import (
	"github.com/stretchr/testify/assert"
	"platform/controllers/campaign"
	"platform/controllers/increase_time"
	"platform/internal/models"
	"platform/internal/testutils"
	time_simulated "platform/internal/time"
	"testing"
	"time"
)

func TestIncreaseTimeManager_IncreaseTime(t *testing.T) {
	testutils.SetupTestDB()
	increaseTimeManager := increase_time.IncreaseTimeManager{
		CampaignManager: campaign.CampaignManager{},
		TimeManager:     time_simulated.DefaultTimeSimulator{},
	}
	product := models.Product{
		Code:  "PI01",
		Price: float64(499),
		Stock: 100,
	}
	newProduct, err := product.Create()
	assert.NoError(t, err)

	input := models.Campaign{

		Expiry:                   time.Now().Add(48 * time.Hour), // 2 gün sonra sona eriyor
		Name:                     "CI01",
		ProductID:                newProduct.ID,
		Duration:                 7,
		TargetSales:              100,
		Status:                   true,
		PriceManipulationLimit:   0.5,
		CurrentPriceManipulation: 0.4,
	}
	newCampaign, err := input.Create()
	assert.NoError(t, err)

	elapsedTime, err := increaseTimeManager.IncreaseTime(24)
	campaignInfo, err := models.Campaign{}.Find("id", newCampaign.ID)
	assert.NoError(t, err)
	assert.True(t, campaignInfo.Expiry.After(elapsedTime))

}
