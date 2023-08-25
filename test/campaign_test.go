package test

import (
	"platform/cmd/commands"
	"platform/internal/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCampaignManager_GetCampaignInfo(t *testing.T) {
	campaignManager := commands.CampaignManager{}

	// Mock bir Campaign verisi oluştur
	mockCampaign := models.Campaign{
		Name:                   "Campaign A",
		Duration:               7,
		PriceManipulationLimit: 10,
		TargetSales:            100,
		TotalSales:             50,
	}
	campaignManager.Campaigns = mockCampaign

	// Test GetCampaignInfo fonksiyonunu
	result := campaignManager.GetCampaignInfo("Campaign A")

	expected := "Kampanya \"Campaign A\" bilgisi; Durum Aktif, Hedef Satış 100, Toplam Satış 50, Ciro 250.00, Ortalama Ürün Fiyatı 5.00"
	assert.Equal(t, expected, result, "Beklenen değer dönmedi")
}

func TestCampaignManager_CreateCampaign(t *testing.T) {
	campaignManager := commands.CampaignManager{}

	// Test CreateCampaign fonksiyonunu
	result := campaignManager.CreateCampaign("Campaign B", "P001", 30, 5, 200)

	expected := "Kampanya oluşturuldu; adı Campaign B, ürün 1, süre 30 saat, limit 5.00, hedef satış 200"
	assert.Equal(t, expected, result, "Beklenen değer dönmedi")
}
