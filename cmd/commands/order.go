package commands

import (
	"fmt"
	"platform/internal/models"
)

type OrderManager struct {
	Order models.Order
}

func (m *OrderManager) CreateOrder(productCode string, quantity int) string {
	product, err := models.Product{}.Find("code", productCode)

	if err != nil {
		return fmt.Sprintf("Ürün bilgileri alınırken bir hata oluştu: %v", err)
	}

	// Kampanyayı bul ve indirimli fiyatı hesapla
	campaign, campaignErr := models.Campaign{}.Find("product_id", product.ID)
	if campaignErr == nil {
		discountedPrice := product.Price * (1 - (campaign.PriceManipulationLimit / 100))
		product.Price = discountedPrice
	}

	if quantity > product.Stock {
		return "Stok yetersiz"
	}

	totalPrice := product.Price * float64(quantity)
	order := models.Order{
		ProductID:  product.ID,
		Quantity:   quantity,
		TotalPrice: totalPrice,
	}

	order, err = order.Create()
	if err != nil {
		return fmt.Sprintf("Sipariş oluşturulurken bir hata oluştu: %v", err)
	}

	// Stok güncelleme işlemi
	discountedStock := product.Stock - quantity
	err = product.Update("stock", discountedStock)
	if err != nil {
		return fmt.Sprintf("Stok güncellenirken bir hata oluştu: %v", err)
	}

	return fmt.Sprintf("Sipariş oluşturuldu; ürün %s, miktar %d, toplam fiyat %.2f",
		order.ProductID, order.Quantity, order.TotalPrice)
}
