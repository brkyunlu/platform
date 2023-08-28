package campaign

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"platform/internal/error_handler"
	"platform/internal/logger"
	"platform/internal/models"
	time_simulated "platform/internal/time"
	"platform/internal/validator"
	"time"
)

type CampaignManager struct {
	Campaigns   models.Campaign
	TimeManager time_simulated.DefaultTimeSimulator
}

type ICampaignManager interface {
	GetCampaignInfo(name string) (*models.Campaign, error)
	CreateCampaign(name string, productCode string, duration int, limit float64, targetSales int) (*models.Campaign, error)
	UpdateCampaignStatus(campaign *models.Campaign) error
}

func (m *CampaignManager) GetCampaignInfo(name string) (*models.Campaign, error) {
	campaign, err := m.Campaigns.Find("name", name)
	if err != nil {
		return nil, error_handler.GetErrorMessage("ErrCampaignNotFound", logger.Error)
	}
	currentTime := m.TimeManager.GetSimulatedTime()
	if campaign.Duration <= 0 && campaign.Status && currentTime.After(campaign.Expiry) {
		campaign.Status = false
		if err := campaign.Update("status", false); err != nil {
			return nil, error_handler.GetErrorMessage("ErrCampaignUpdate", logger.Error)
		}
	}

	return &campaign, nil
}

func (m *CampaignManager) CreateCampaign(name string, productCode string, duration int, limit float64, targetSales int) (*models.Campaign, error) {
	product, err := m.Campaigns.Product.Find("code = ?", productCode)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, error_handler.GetErrorMessage("ErrProductNotFound", logger.Error)
		}
		return nil, error_handler.GetErrorMessage("ErrProductQuery", logger.Error)
	}

	// Check the active campaign of the product
	existingCampaign, _ := m.Campaigns.Find("product_id = ? AND status = ?", product.ID, true)

	if existingCampaign.Expiry.After(m.TimeManager.GetSimulatedTime()) {
		return nil, error_handler.GetErrorMessage("ErrActiveCampaignExists", logger.Error)
	}

	currentTime := m.TimeManager.GetSimulatedTime()
	expiryTime := currentTime.Add(time.Hour * time.Duration(duration))

	campaign := models.Campaign{
		Name:                     name,
		ProductID:                product.ID,
		Duration:                 duration,
		PriceManipulationLimit:   limit,
		TargetSales:              targetSales,
		CurrentPriceManipulation: 0,
		Status:                   true,
		Expiry:                   expiryTime,
	}

	v := validator.New()
	if models.ValidateCampaign(v, &campaign); !v.Valid() {
		return nil, errors.New(fmt.Sprint(v.Errors))
	}

	if _, createErr := campaign.Create(); createErr != nil {
		return nil, error_handler.GetErrorMessage("ErrCampaignCreate", logger.Error)
	}

	return &campaign, nil
}

func (m *CampaignManager) UpdateCampaignStatus(campaign *models.Campaign) error {
	currentTime := m.TimeManager.GetSimulatedTime()

	if !campaign.Status || currentTime.After(campaign.Expiry) {
		// If the campaign has expired or is already inactive, update its status and reset the discount.
		err := campaign.Updates(models.Campaign{CurrentPriceManipulation: 0, Status: false})
		if err != nil {
			return fmt.Errorf("An error occurred while updating the status of campaign \"%s\": %v", campaign.Name, err)
		}
	} else {
		// If the campaign is still active and has not expired:
		// Increase the discount rate, but ensure it does not exceed a certain price manipulation limit.
		discountIncrement := campaign.PriceManipulationLimit / float64(campaign.Duration)
		timeRemaining := campaign.Expiry.Sub(currentTime)
		remainingHours := int(timeRemaining.Hours())

		// Increase the discount rate and ensure it does not exceed the specified limit.
		newDiscount := campaign.CurrentPriceManipulation + (float64(remainingHours) * discountIncrement)
		if newDiscount > campaign.PriceManipulationLimit {
			newDiscount = campaign.PriceManipulationLimit
		}

		err := campaign.Update("current_price_manipulation", newDiscount)
		if err != nil {
			return fmt.Errorf("An error occurred while updating the discount limit of campaign \"%s\": %v", campaign.Name, err)
		}
	}

	return nil
}
