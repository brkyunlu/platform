package seeders

import (
	"fmt"
	"math/rand"
	"platform/internal/models"
)

func ProductSeed() {
	counter := models.Product{}.Count("", "")
	if counter < 1 {
		for i := 0; i < 10; i++ {
			product := &models.Product{
				Code:  fmt.Sprintf("P%d", i+1),
				Price: float64(rand.Intn(250)),
				Stock: rand.Intn(250),
			}
			_, err := product.Create()
			if err != nil {
				return
			}
		}
	}
}
