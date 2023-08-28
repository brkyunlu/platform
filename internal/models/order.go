package models

import "platform/internal/client"

type Order struct {
	ID         int64     `gorm:"primary_key"`
	ProductID  int64     `json:"-"`
	Product    Product   `json:"product" gorm:"references:ID"`
	CampaignID *int64    `json:"-"`
	Campaign   *Campaign `json:"campaign" gorm:"references:ID"`
	Quantity   int
	TotalPrice float64
}

func (order Order) Create() (Order, error) {
	result := client.PostgreSqlClient.Create(&order)
	return order, result.Error
}
func (order Order) Find(query ...interface{}) (Order, error) {
	result := client.PostgreSqlClient.Joins("Product", "products on orders.product_id = product.id").
		Joins("Campaign", "campaign on orders.campaign_id = campaign.id").
		First(&order, query...)
	return order, result.Error
}
func (order Order) Select(query ...interface{}) (Order, error) {
	result := client.PostgreSqlClient.Select(query)
	return order, result.Error
}
func (order Order) Count(column string, value interface{}) int64 {
	var counter int64
	postClient := client.PostgreSqlClient.Model(&order)
	if column != "" && value != "" {
		postClient.Where(column, value)
	}
	postClient.Count(&counter)
	return counter
}
