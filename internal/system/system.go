package system

import (
	"fmt"
	"platform/cmd/commands"
	"strings"
)

func ProcessCommand(command string) {
	parts := strings.Fields(command)
	if len(parts) == 0 {
		return
	}

	switch parts[0] {
	case "create_product":
		commands.CreateProduct(parts)
	case "get_product_info":
		commands.GetProductInfo(parts)
	case "create_order":
		commands.CreateOrder(parts)
	case "create_campaign":
		commands.CreateCampaign(parts)
	case "get_campaign_info":
		commands.GetCampaignInfo(parts)
	case "increase_time":
		commands.IncreaseTime(parts)
	default:
		fmt.Println("Geçersiz komut:", parts[0])
	}
}
