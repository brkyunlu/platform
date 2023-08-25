package test

import (
	"platform/cmd/commands"
	"platform/internal/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProductManager_GetProductInfo(t *testing.T) {
	productManager := commands.ProductManager{}

	// Mock bir Product verisi oluştur
	mockProduct := models.Product{
		Code:  "P001",
		Price: 9.99,
		Stock: 100,
	}
	productManager.Products = mockProduct

	// Test GetProductInfo fonksiyonunu
	result := productManager.GetProductInfo("P001")

	expected := "Ürün P001 bilgisi; fiyat 9.99, stok 100"
	assert.Equal(t, expected, result, "Beklenen değer dönmedi")
}

func TestProductManager_CreateProduct(t *testing.T) {
	productManager := commands.ProductManager{}

	// Test CreateProduct fonksiyonunu
	result := productManager.CreateProduct("P002", 19.99, 50)

	expected := "Ürün oluşturuldu; kod P002, fiyat 19.99, stok 50"
	assert.Equal(t, expected, result, "Beklenen değer dönmedi")
}
