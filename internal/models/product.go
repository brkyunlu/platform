package models

import (
	"platform/internal/client"
	"platform/internal/validator"
	"regexp"
)

type Product struct {
	ID    int64   `json:"id" gorm:"primarykey"` // Unique integer ID for the product
	Code  string  `json:"code" gorm:"unique"`
	Price float64 `json:"price" gorm:"type:integer;not null"`
	Stock int     `json:"stock" gorm:"type:integer;not null"`
}

func (product Product) Create() (Product, error) {
	result := client.PostgreSqlClient.Create(&product)
	return product, result.Error
}
func (product Product) Update(column string, value interface{}) error {
	result := client.PostgreSqlClient.Model(&product).Update(column, value)
	return result.Error
}

func (product Product) Updates(data Product) error {
	result := client.PostgreSqlClient.Model(&product).Updates(data)
	return result.Error
}
func (product Product) Find(query ...interface{}) (Product, error) {
	result := client.PostgreSqlClient.First(&product, query...)
	return product, result.Error
}

func (product Product) Count(column string, value interface{}) int64 {
	var counter int64
	postClient := client.PostgreSqlClient.Model(&product)
	if column != "" && value != "" {
		postClient.Where(column, value)
	}
	postClient.Count(&counter)
	return counter
}
func ValidateProduct(v *validator.Validator, product *Product) {
	v.Check(product.Code != "", "code", "boş bırakılamaz")

	v.Check(product.Price >= 0, "price", "pozitif bir sayı olmalıdır")

	v.Check(product.Stock >= 0, "stock", "pozitif bir sayı olmalıdır")
	product.ValidateUniqueCode(v) //check unique code
	product.ValidateCodeFormat(v) //check code format (just use 0-1 && A-Z)
}
func (product Product) ValidateUniqueCode(v *validator.Validator) {
	var existingProduct Product
	result := client.PostgreSqlClient.Where("code = ?", product.Code).First(&existingProduct)
	if result.Error == nil && existingProduct.ID != product.ID {
		v.AddError("code", "bu kod zaten kullanılıyor")
	}
}
func (product Product) ValidateCodeFormat(v *validator.Validator) {
	// Regular expression to match only letters and numbers
	validCodePattern := regexp.MustCompile("^[a-zA-Z0-9]+$")
	if !validCodePattern.MatchString(product.Code) {
		v.AddError("code", "sadece harf ve sayı içermelidir")
	}
}
