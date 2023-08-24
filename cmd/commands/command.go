package commands

import "time"

type TimeManager struct {
	CurrentTime time.Time
}

func (t *TimeManager) Now() time.Time {
	return t.CurrentTime
}

type Command interface {
	Execute() string
}

type CommandManager interface {
	GetProductInfo(code string) string
	CreateProduct(code string, price float64, stock int) string
	GetCampaignInfo(name string) string
	CreateCampaign(name string, productID int64, duration int, limit float64, targetSales int) string
	// Diğer komut işlemleri burada tanımlanabilir
}
