package commands

import (
	"fmt"
	"platform/internal/models"
	"time"
)

type CommandProcessor struct {
	ProductManager      ProductManagerInterface
	CampaignManager     CampaignManagerInterface
	OrderManager        OrderManagerInterface
	IncreaseTimeManager IncreaseTimeManager
	TimeManager         *TimeManager
}
type TimeManager struct {
	CurrentTime time.Time
}

func (t *TimeManager) Now() time.Time {
	return t.CurrentTime
}
func NewCommandProcessor(timeManager *TimeManager) *CommandProcessor {
	return &CommandProcessor{
		ProductManager:      &ProductManager{Products: models.Product{}},
		CampaignManager:     &CampaignManager{Campaigns: models.Campaign{}},
		OrderManager:        &OrderManager{Order: models.Order{}},
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
			product, err := p.ProductManager.GetProductInfo(productCode)
			if err != nil {
				fmt.Println("Hata:", err)
				continue
			}
			fmt.Printf(`Ürün %s bilgisi; fiyat %.2f, stok %d`,
				product.Code, product.Price, product.Stock)
		case "create_product":
			productCode := getUserInput("-Ürün kodu girin: ")
			productPrice := getUserFloatInput("-Fiyat girin: ")
			productStock := getUserIntInput("-Stok girin: ")
			product, err := p.ProductManager.CreateProduct(productCode, productPrice, productStock)
			if err != nil {
				fmt.Println("Hata:", err)
			}
			fmt.Printf(`Ürün oluşturuldu; kod %s, fiyat %.2f, stok %d`,
				product.Code, product.Price, product.Stock)
		case "get_campaign_info":
			campaignName := getUserInput("-Kampanya Adı Girin: ")
			campaign, err := p.CampaignManager.GetCampaignInfo(campaignName)
			if err != nil {
				fmt.Println("Hata:", err)
			} else {
				fmt.Printf(`Kampanya "%s" bilgisi; Durum %t, Hedef Satış %d, Toplam Satış %d`,
					campaign.Name, campaign.Status, campaign.TargetSales, campaign.TotalSales)
			}
		case "create_campaign":
			campaignName := getUserInput("-Kampanya adı girin: ")
			productCode := getUserInput("-Ürün kodu girin: ")
			duration := getUserIntInput("-Süre girin: ")
			priceLimit := getUserFloatInput("-Fiyat manipülasyon limiti girin: ")
			targetSales := getUserIntInput("-Hedef satış girin: ")
			campaign, err := p.CampaignManager.CreateCampaign(campaignName, productCode, duration, priceLimit, targetSales)
			if err != nil {
				fmt.Println("Hata:", err)
			} else {
				fmt.Println(`Kampanya oluşturuldu; adı %s, ürün %s, süre %d saat, limit %.2f, hedef satış %d`,
					campaign.Name, campaign.ProductID, campaign.PriceManipulationLimit, campaign.TargetSales)
			}
		case "increase_time":
			hours := getUserIntInput("-Kaç saat ilerletilsin: ")
			newTime := p.IncreaseTimeManager.IncreaseTime(hours)
			fmt.Println("Yeni zaman:", newTime.Format("15:04"))
		default:
			fmt.Println("Geçersiz komut. Lütfen geçerli bir komut girin.")
		}
	}
}
