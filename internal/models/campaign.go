package models

import (
	"errors"
	"platform/internal/client"
	"platform/internal/validator"
	"time"
)

var ErrProductNotFound = errors.New("the product you wanted to buy was not found")

type Campaign struct {
	ID                       int64     `json:"id" gorm:"primary_key"`
	Name                     string    `json:"name" gorm:"not null"`
	ProductID                int64     `json:"-"`
	Product                  Product   `json:"product" gorm:"references:ID"`
	Duration                 int       `json:"duration" gorm:"not null;check:duration > 0"`
	Expiry                   time.Time `json:"expiry"`
	Status                   bool      `json:"status" gorm:"not null"`
	PriceManipulationLimit   float64   `json:"price_manipulation_limit" gorm:"not null;check:price_manipulation_limit >= 0 AND price_manipulation_limit <= 100"`
	CurrentPriceManipulation float64   `json:"current_price_manipulation" gorm:"not null;check:current_price_manipulation >= 0 AND current_price_manipulation <= price_manipulation_limit"`
	TargetSales              int       `json:"target_sales" gorm:"not null;check:target_sales >= 0"`
	TotalSales               int       `json:"total_sales"`
	Revenue                  float64   `json:"revenue"`
	AveragePrice             float64   `json:"average_price"`
}

func (campaign Campaign) Create() (Campaign, error) {
	result := client.PostgreSqlClient.Create(&campaign)
	return campaign, result.Error
}
func (campaign Campaign) Update(column string, value interface{}) error {
	result := client.PostgreSqlClient.Model(&campaign).Update(column, value)
	return result.Error
}

func (campaign Campaign) Updates(data Campaign) error {
	result := client.PostgreSqlClient.Model(&campaign).Updates(data)
	return result.Error
}
func (campaign Campaign) Find(query ...interface{}) (Campaign, error) {
	result := client.PostgreSqlClient.First(&campaign, query...)
	return campaign, result.Error
}
func (campaign Campaign) Count(column string, value interface{}) int64 {
	var counter int64
	postClient := client.PostgreSqlClient.Model(&campaign)
	if column != "" && value != "" {
		postClient.Where(column, value)
	}
	postClient.Count(&counter)
	return counter
}
func GetAllCampaigns() ([]Campaign, error) {
	var campaigns []Campaign
	result := client.PostgreSqlClient.Find(&campaigns)
	return campaigns, result.Error
}
func ValidateCampaign(v *validator.Validator, campaign *Campaign) {
	validateCampaignProduct(v, campaign.ProductID)
	v.Check(campaign.Name != "", "name", "boş bırakılamaz")
	v.Check(campaign.Duration > 0, "duration", "geçerli bir süre olmalıdır")
	v.Check(campaign.PriceManipulationLimit >= 0 && campaign.PriceManipulationLimit <= 100, "price_manipulation_limit", "0 ile 100 arasında olmalıdır")
	v.Check(campaign.TargetSales >= 0, "target_sales", "pozitif bir sayı olmalıdır")
}
func validateCampaignProduct(v *validator.Validator, productID int64) {
	product, _ := Product{}.Find("id", productID)
	v.Check(product.ID != 0, "product_id", ErrProductNotFound.Error())
}
