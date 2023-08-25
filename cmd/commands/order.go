// order.go
package commands

import (
	"fmt"
	"platform/internal/models"
)

type OrderManager struct {
	Order models.Order
}

type OrderManagerInterface interface {
	CreateOrder(productCode string, quantity int) string
}

func (m *OrderManager) CreateOrder(productCode string, quantity int) string {
	product, err := m.Order.Product.Find("code", productCode)
	if err != nil {
		return fmt.Sprintf("Ürün bilgileri alınırken bir hata oluştu: %v", err)
	}

	campaign, campaignErr := m.Order.Campaign.Find("product_id", product.ID)
	if campaignErr == nil {
		discountedPrice := product.Price * (1 - (campaign.CurrentPriceManipulation / 100))
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

	_, err = m.Order.Create()
	if err != nil {
		return fmt.Sprintf("Sipariş oluşturulurken bir hata oluştu: %v", err)
	}

	updatedStock := product.Stock - quantity
	err = m.Order.Product.Update("stock", updatedStock)
	if err != nil {
		return fmt.Sprintf("Stok güncellenirken bir hata oluştu: %v", err)
	}

	return fmt.Sprintf("Sipariş oluşturuldu; ürün %s, miktar %d, toplam fiyat %.2f",
		order.ProductID, order.Quantity, order.TotalPrice)
}
