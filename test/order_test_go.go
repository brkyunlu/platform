package test

import (
	"platform/cmd/commands"
	"platform/internal/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrderManager_CreateOrder(t *testing.T) {
	orderManager := commands.OrderManager{}

	// Mock bir Product verisi oluştur
	mockProduct := models.Product{
		ID:    11,
		Code:  "P001",
		Price: 9.99,
		Stock: 100,
	}
	orderManager.Order = models.Order{}
	orderManager.Order.Product = mockProduct

	// Mock bir Campaign verisi oluştur (kampanya varsa)
	mockCampaign := models.Campaign{
		ID:                     1,
		ProductID:              1,
		PriceManipulationLimit: 10,
	}
	orderManager.Order.Campaign = mockCampaign

	// Test CreateOrder fonksiyonunu (kampanyasız)
	result := orderManager.CreateOrder("P001", 5)

	expected := "Sipariş oluşturuldu; ürün 1, miktar 5, toplam fiyat 49.95"
	assert.Equal(t, expected, result, "Beklenen değer dönmedi")

	// Test CreateOrder fonksiyonunu (stok yetersiz)
	result = orderManager.CreateOrder("P001", 150)

	expected = "Stok yetersiz"
	assert.Equal(t, expected, result, "Beklenen değer dönmedi")

	// Kampanya oluştur
	mockCampaign = models.Campaign{
		ID:                     1,
		ProductID:              1,
		PriceManipulationLimit: 10,
	}
	orderManager.Order.Campaign = mockCampaign

	// Test CreateOrder fonksiyonunu (kampanyalı)
	result = orderManager.CreateOrder("P001", 10)

	expected = "Sipariş oluşturuldu; ürün 1, miktar 10, toplam fiyat 89.91"
	assert.Equal(t, expected, result, "Beklenen değer dönmedi")
}
