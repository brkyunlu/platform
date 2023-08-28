package cmd

import (
	"github.com/spf13/cobra"
	"platform/controllers/handlers/order"
)

func init() {
	Command.AddCommand(&cobra.Command{
		Use:   "create_order [product_code] [quantity]",
		Short: "Create a new order",
		Args:  cobra.ExactArgs(2),
		RunE:  order.OrderHandler,
	})
}
