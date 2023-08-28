package product_test

import (
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"platform/controllers/handlers/product"
	"platform/internal/models"
	"platform/internal/testutils"
	"testing"
)

func TestGetProductInfoHandler(t *testing.T) {
	testutils.SetupTestDB()

	newProduct := models.Product{
		Code:  "PH01",
		Price: 29.99,
		Stock: 100,
	}

	createdProduct, err := newProduct.Create()
	assert.NoError(t, err)

	// Burada test için gerekli olan veritabanı veya sahte veri oluşturulabilir
	cmd := &cobra.Command{}
	args := []string{createdProduct.Code}

	// Test case 1: Ürün bilgisini alırken başarılı bir durum
	err = product.GetProductInfoHandler(cmd, args)
	assert.NoError(t, err)

	// Test case 2: Ürün kodu eksik olduğunda hata dönmeli
	err = product.GetProductInfoHandler(nil, []string{})
	assert.Error(t, err)
}
