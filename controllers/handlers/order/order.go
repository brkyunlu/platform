package order

import (
	"fmt"
	"github.com/spf13/cobra"
	"platform/controllers/order"
	"platform/internal/error_handler"
	"platform/internal/logger"
	"strconv"
)

func OrderHandler(cmd *cobra.Command, args []string) error {
	if len(args) != 2 {
		return error_handler.GetErrorMessage("OrderRequirement", logger.Error)
	}

	productCode := args[0]
	quantity, _ := strconv.Atoi(args[1])
	orderManager := &order.OrderManager{}
	createOrder, err := orderManager.CreateOrder(productCode, quantity)
	if err != nil {
		return err
	}

	fmt.Printf(`Order created; product %s, quantity %d, total price %.2f`,
		createOrder.Product.Code, createOrder.Quantity, createOrder.TotalPrice)
	return nil
}
