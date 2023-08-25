package commands

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"math/rand"
	"platform/internal/models"
	"platform/internal/validator"
	"time"
)

type CampaignManager struct {
	Campaigns   models.Campaign
	TimeManager *TimeManager
}

type CampaignManagerInterface interface {
	GetCampaignInfo(name string) (*models.Campaign, error)
	CreateCampaign(name string, productCode string, duration int, limit float64, targetSales int) (*models.Campaign, error)
}

func (m *CampaignManager) GetCampaignInfo(name string) (*models.Campaign, error) {
	campaign, err := m.Campaigns.Find("name", name)
	if err != nil {
		return nil, GetErrorMessage("ErrCampaignNotFound")
	}

	// Kampanya süresi dolmuşsa ve durum aktifse durumu güncelle
	currentTime := m.TimeManager.Now()
	if campaign.Duration <= 0 && campaign.Status && currentTime.After(campaign.Expiry) {
		campaign.Status = false
		if err := campaign.Update("status", false); err != nil {
			return nil, GetErrorMessage("ErrCampaignUpdate")
		}
	}

	return &campaign, nil
}

func (m *CampaignManager) CreateCampaign(name string, productCode string, duration int, limit float64, targetSales int) (*models.Campaign, error) {
	product, err := m.Campaigns.Product.Find("code = ?", productCode)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, GetErrorMessage("ErrProductNotFound")
		}
		return nil, GetErrorMessage("ErrProductQuery")
	}

	// Ürüne ait aktif kampanyayı kontrol et
	existingCampaign, _ := m.Campaigns.Find("product_id = ? AND status = ?", product.ID, true)

	if existingCampaign.Expiry.After(m.TimeManager.Now()) {
		return nil, GetErrorMessage("ErrActiveCampaignExists")
	}

	randomLimit := rand.Float64() * limit
	currentTime := m.TimeManager.Now()
	expiryTime := currentTime.Add(time.Hour * time.Duration(duration))

	campaign := &models.Campaign{
		Name:                     name,
		ProductID:                product.ID,
		Duration:                 duration,
		PriceManipulationLimit:   limit,
		TargetSales:              targetSales,
		CurrentPriceManipulation: randomLimit,
		Status:                   true,
		Expiry:                   expiryTime,
	}

	v := validator.New()
	if models.ValidateCampaign(v, campaign); !v.Valid() {
		return nil, errors.New(fmt.Sprint(v.Errors))
	}

	if _, createErr := m.Campaigns.Create(); createErr != nil {
		return nil, GetErrorMessage("ErrCampaignCreate")
	}

	return campaign, nil
}
func (m *CampaignManager) UpdateCampaignStatus(campaign *models.Campaign) error {
	currentTime := m.TimeManager.Now()

	if campaign.Status && currentTime.After(campaign.Expiry) {
		if campaign.CurrentPriceManipulation > 0 {
			discountHourly := campaign.CurrentPriceManipulation / float64(campaign.Duration)
			elapsedHours := int(currentTime.Sub(campaign.Expiry).Hours())
			newLimit := campaign.CurrentPriceManipulation - (float64(elapsedHours) * discountHourly)
			if newLimit < 0 {
				newLimit = 0
			}
			err := campaign.Update("current_price_manipulation", newLimit)
			if err != nil {
				return fmt.Errorf("Kampanya \"%s\" indirim limiti güncellenirken bir hata oluştu: %v", campaign.Name, err)
			}
		}
		err := campaign.Updates(models.Campaign{CurrentPriceManipulation: 0, Status: false})
		if err != nil {
			return fmt.Errorf("Kampanya \"%s\" durumu güncellenirken bir hata oluştu: %v", campaign.Name, err)
		}
	}
	return nil
}
