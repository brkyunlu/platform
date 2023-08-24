package commands

import (
	"fmt"
	"platform/internal/models"
	"platform/internal/validator"
)

type ProductManager struct {
	Products models.Product
}

func (manager *ProductManager) GetProductInfo(code string) string {
	product, err := models.Product{}.Find("code", code)
	if err != nil {
		return "Ürün bulunamadı."
	}

	return fmt.Sprintf("Ürün %s bilgisi; fiyat %.2f, stok %d", product.Code, product.Price, product.Stock)

}

func (manager *ProductManager) CreateProduct(code string, price float64, stock int) string {
	manager.Products = models.Product{Code: code, Price: price, Stock: stock}

	v := validator.New()

	if models.ValidateProduct(v, &manager.Products); !v.Valid() {
		return fmt.Sprint(v.Errors)
	}
	// Veritabanına kaydet
	product, createErr := manager.Products.Create()
	if createErr != nil {
		return "Ürün oluşturulurken bir hata oluştu."
	}

	return fmt.Sprintf("Ürün oluşturuldu; kod %s, fiyat %.2f, stok %d", product.Code, product.Price, product.Stock)
}
