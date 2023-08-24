package commands

import (
	"fmt"
	"platform/internal/models"
)

type CommandProcessor struct {
	ProductManager      ProductManager
	CampaignManager     CampaignManager
	OrderManager        OrderManager
	IncreaseTimeManager IncreaseTimeManager
	TimeManager         *TimeManager
}

func NewCommandProcessor(timeManager *TimeManager) *CommandProcessor {
	return &CommandProcessor{
		ProductManager:      ProductManager{Products: models.Product{}},
		CampaignManager:     CampaignManager{Campaigns: models.Campaign{}},
		OrderManager:        OrderManager{Order: models.Order{}},
		IncreaseTimeManager: IncreaseTimeManager{TimeManager: timeManager},
		TimeManager:         timeManager,
	}
}
func (p *CommandProcessor) Run() {

	for {
		command := getUserInput("Komutları girin (Çıkış için: exit)\n>>> ")

		switch command {
		case "exit":
			fmt.Println("Çıkış yapılıyor...")
			return
		case "get_product_info":
			productCode := getUserInput("-Ürün Kodu Girin: ")
			result := p.ProductManager.GetProductInfo(productCode)
			fmt.Println(result)
		case "create_product":
			productCode := getUserInput("-Ürün kodu girin: ")
			productPrice := getUserFloatInput("-Fiyat girin: ")
			productStock := getUserIntInput("-Stok girin: ")
			result := p.ProductManager.CreateProduct(productCode, productPrice, productStock)
			fmt.Println(result)
		case "get_campaign_info":
			campaignName := getUserInput("-Kampanya Adı Girin: ")
			result := p.CampaignManager.GetCampaignInfo(campaignName)
			fmt.Println(result)
		case "create_campaign":
			campaignName := getUserInput("-Kampanya adı girin: ")
			productID := getUserIntInput("-Ürün kodu girin: ")
			duration := getUserIntInput("-Süre girin: ")
			priceLimit := getUserFloatInput("-Fiyat manipülasyon limiti girin: ")
			targetSales := getUserIntInput("-Hedef satış girin: ")
			result := p.CampaignManager.CreateCampaign(campaignName, int64(productID), duration, priceLimit, targetSales)
			// Kampanya oluşturulduktan sonra süre dolduysa kampanya durumunu güncelle
			if p.CampaignManager.Campaigns.Duration <= 0 {
				p.CampaignManager.Campaigns.Duration = 0
				p.CampaignManager.Campaigns.Name = "Sona Erdi"
			}
			fmt.Println(result)
		case "increase_time":
			hours := getUserIntInput("-Kaç saat ilerletilsin: ")
			newTime := p.IncreaseTimeManager.IncreaseTime(hours)
			fmt.Println("Yeni zaman:", newTime.Format("15:04"))
		default:
			fmt.Println("Geçersiz komut. Lütfen geçerli bir komut girin.")
		}
	}
}
