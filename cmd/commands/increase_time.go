package commands

import (
	"fmt"
	"platform/internal/models"
	"time"
)

type IncreaseTimeManager struct {
	TimeManager     *TimeManager
	CampaignManager CampaignManager
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

	// Kampanya durumlarını güncelleme işlemini paralel olarak yap ve sonucu başka bir channel'e aktar

	updatedCampaignsChan := make(chan models.Campaign)
	for campaign := range campaignStatusChan {
		go m.CampaignManager.UpdateCampaignStatus(&campaign) // UpdateCampaignStatus fonksiyonunu CampaignManager üzerinden çağırdık

	}

	// Paralel işlemlerin tamamlanmasını bekle
	for range allCampaigns {
		<-updatedCampaignsChan
	}
	close(updatedCampaignsChan)

	return newTime
}
