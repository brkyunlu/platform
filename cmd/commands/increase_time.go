package commands

import (
	"fmt"
	"platform/internal/models"
	"time"
)

type IncreaseTimeManager struct {
	TimeManager *TimeManager
}

func (m *IncreaseTimeManager) IncreaseTime(hours int) time.Time {
	fmt.Println("Zaman ilerliyor...")

	currentTime := m.TimeManager.Now()
	newTime := currentTime.Add(time.Duration(hours) * time.Hour)

	// Kampanyaların durumunu kontrol etmek için bir channel oluştur
	campaignStatusChan := make(chan models.Campaign)

	// Tüm kampanyaları al
	allCampaigns, err := models.GetAllCampaigns()
	if err != nil {
		fmt.Println("Kampanya verileri alınamadı:", err)
		return newTime
	}

	// Kampanyaları channel'e gönder
	for _, campaign := range allCampaigns {
		campaignStatusChan <- campaign
	}
	close(campaignStatusChan)

	// Kampanya durumlarını güncelleme işlemini paralel olarak yap
	// ve sonucu başka bir channel'e aktar
	updatedCampaignsChan := make(chan models.Campaign)
	for campaign := range campaignStatusChan {
		go m.updateCampaignStatus(currentTime, campaign, updatedCampaignsChan)
	}

	// Paralel işlemlerin tamamlanmasını bekle
	for range allCampaigns {
		<-updatedCampaignsChan
	}
	close(updatedCampaignsChan)

	return newTime
}

func (m *IncreaseTimeManager) updateCampaignStatus(currentTime time.Time, c models.Campaign, updatedCampaignsChan chan models.Campaign) {
	if c.Status == true && currentTime.After(c.Expiry) {
		// Kampanya başladıktan sonra belirli saatlerde indirim oranı güncelleme
		if c.CurrentPriceManipulation > 0 {
			discountHourly := c.CurrentPriceManipulation / float64(c.Duration) // Saat başına düşen indirim
			elapsedHours := int(currentTime.Sub(c.Expiry).Hours())
			newLimit := c.CurrentPriceManipulation - (float64(elapsedHours) * discountHourly)
			if newLimit < 0 {
				newLimit = 0
			}
			// Kampanya indirim limitini güncelle
			err := c.Update("current_price_manipulation", newLimit)
			if err != nil {
				fmt.Printf("Kampanya \"%s\" indirim limiti güncellenirken bir hata oluştu: %v\n", c.Name, err)
			}
		}
		// Kampanyanın indirim değerlerini sıfırla ve durumunu false olarak işaretle
		err := c.Updates(models.Campaign{CurrentPriceManipulation: 0, Status: false})
		if err != nil {
			fmt.Printf("Kampanya \"%s\" durumu güncellenirken bir hata oluştu: %v\n", c.Name, err)
		} else {
			updatedCampaignsChan <- c
		}
	}
}
