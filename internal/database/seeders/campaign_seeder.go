package seeders

import (
	"fmt"
	"math/rand"
	"platform/internal/models"
	"time"
)

func CampaignSeed() {
	counter := models.Campaign{}.Count("", "")
	if counter < 1 {

		for i := 0; i < 10; i++ {

			duration := time.Duration(rand.Intn(10)+1) * time.Hour
			priceManipulationLimit := float64(rand.Intn(10)) + 1.0
			targetSales := rand.Intn(50) + 50

			campaign := &models.Campaign{
				Name:                   fmt.Sprintf("Campaign %d", i+1),
				ProductID:              int64(i + 1),
				Duration:               int(duration),
				PriceManipulationLimit: priceManipulationLimit,
				TargetSales:            targetSales,
			}

			_, err := campaign.Create()
			if err != nil {
				return
			}
			fmt.Printf("Kampanya oluşturuldu; adı %s, ürün kodu %s\n", campaign.Name, fmt.Sprintf("P%d", i+1))
		}
	}
}
