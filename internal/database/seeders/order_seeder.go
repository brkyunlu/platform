package seeders

import (
	"fmt"
	"math/rand"
	"platform/internal/models"
)

func OrderSeed() {
	counter := models.Order{}.Count("", "")
	if counter < 1 {

		for i := 0; i < 10; i++ {

			product, _ := models.Product{}.Find("id", i+1)
			if product.ID == 0 {
				return
			}
			quantity := rand.Intn(10) + 1

			order := &models.Order{
				ProductID:  product.ID,
				Quantity:   quantity,
				TotalPrice: product.Price * float64(quantity),
			}

			_, err := order.Create()
			if err != nil {
				return
			}
			fmt.Printf("Sipariş oluşturuldu; ürün %s, miktar %d\n", product.Code, quantity)
		}
	}
}
