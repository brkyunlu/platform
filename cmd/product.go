package cmd

import (
	"github.com/spf13/cobra"
	"platform/controllers/handlers/product"
)

func init() {
	Command.AddCommand(&cobra.Command{
		Use:   "get_product_info [product_code]",
		Short: "Get information about a product",
		Args:  cobra.ExactArgs(1),
		RunE:  product.GetProductInfoHandler,
	})
	Command.AddCommand(&cobra.Command{
		Use:   "create_product [product_code] [product_price] [product_stock]",
		Short: "Create a new product",
		Args:  cobra.ExactArgs(3),
		RunE:  product.CreateProductHandler,
	})
}
