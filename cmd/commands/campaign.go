package commands

import (
	"fmt"
	"math/rand"
	"platform/internal/models"
	"platform/internal/validator"
)

type CampaignManager struct {
	Campaigns   models.Campaign
	TimeManager *TimeManager
}

func (m *CampaignManager) GetCampaignInfo(name string) string {
	var err error
	m.Campaigns, err = models.Campaign{}.Find("name", name)
	if err != nil {
		return "Kampanya bulunamadı."
	}

	// Kampanya süresi dolmuşsa ve durum aktifse durumu güncelle
	currentTime := m.TimeManager.Now() // Şu anki zamanı al
	if m.Campaigns.Duration <= 0 && m.Campaigns.Status && currentTime.After(m.Campaigns.Expiry) {
		m.Campaigns.Status = false
		err = m.Campaigns.Update("status", false)
		if err != nil {
			return "Kampanya durumu güncellenirken bir hata oluştu."
		}
	}
	averagePrice := float64(m.Campaigns.TotalSales) / float64(m.Campaigns.TargetSales)
	return fmt.Sprintf(`Kampanya "%s" bilgisi; Durum Aktif, Hedef Satış %d, Toplam Satış %d, Ciro %.2f, Ortalama Ürün Fiyatı %.2f`,
		m.Campaigns.Name, m.Campaigns.TargetSales, m.Campaigns.TotalSales, float64(m.Campaigns.TotalSales)*averagePrice, averagePrice)
}

func (m *CampaignManager) CreateCampaign(name string, productID int64, duration int, limit float64, targetSales int) string {
	randomLimit := rand.Float64() * limit
	m.Campaigns = models.Campaign{
		Name:                     name,
		ProductID:                productID,
		Duration:                 duration,
		PriceManipulationLimit:   limit,
		TargetSales:              targetSales,
		CurrentPriceManipulation: randomLimit, // Yeni alanı ayarla
	}
	v := validator.New()

	if models.ValidateCampaign(v, &m.Campaigns); !v.Valid() {
		return fmt.Sprint(v.Errors)
	}
	// Veritabanına kaydet
	campaign, createErr := m.Campaigns.Create()
	if createErr != nil {
		return "kampanya oluşturulurken bir hata oluştu."
	}
	return fmt.Sprintf("Kampanya oluşturuldu; adı %s, ürün %s, süre %d saat, limit %.2f, hedef satış %d",
		campaign.Name, campaign.ProductID, campaign.Duration, campaign.PriceManipulationLimit, campaign.TargetSales)
}
