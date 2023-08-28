package product_test

import (
	"platform/internal/testutils"
	"testing"

	"github.com/stretchr/testify/assert"
	"platform/controllers/product"
	"platform/internal/models"
)

func TestProductManager_GetProductInfo(t *testing.T) {

	testutils.SetupTestDB()

	mockProductManager := product.ProductManager{
		Products: models.Product{
			Code:  "P01",
			Price: 9.99,
			Stock: 50,
		},
	}

	productInfo, err := mockProductManager.GetProductInfo("P01")
	assert.NoError(t, err)
	assert.NotNil(t, productInfo)

	productInfo, err = mockProductManager.GetProductInfo("non_existing_product")
	assert.Error(t, err)
	assert.Nil(t, productInfo)
}

func TestProductManager_CreateProduct(t *testing.T) {
	testutils.SetupTestDB()

	mockProductManager := product.ProductManager{
		Products: models.Product{
			Code:  "P01",
			Price: 1999,
			Stock: 100,
		},
	}

	createdProduct, err := mockProductManager.CreateProduct("P02", 1999, 100)
	assert.NoError(t, err)
	assert.NotNil(t, createdProduct)

	invalidProductManager := product.ProductManager{
		Products: models.Product{
			Code:  "",   // Boş ürün kodu
			Price: -999, // Negatif fiyat
			Stock: -50,  // Negatif stok
		},
	}

	invalidCreatedProduct, err := invalidProductManager.CreateProduct("", -999, -50)
	assert.Error(t, err)
	assert.Nil(t, invalidCreatedProduct)
}
