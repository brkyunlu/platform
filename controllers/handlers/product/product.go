package product

import (
	"fmt"
	"github.com/spf13/cobra"
	"platform/controllers/product"
	"platform/internal/error_handler"
	"platform/internal/logger"
	"strconv"
)

func GetProductInfoHandler(cmd *cobra.Command, args []string) error {
	if len(args) != 1 {
		return error_handler.GetErrorMessage("GetProductRequirement", logger.Error)
	}

	productCode := args[0]

	productManager := &product.ProductManager{}
	product, err := productManager.GetProductInfo(productCode)
	if err != nil {
		return err
	}

	fmt.Printf(`Product %s information; price %.2f, stock %d`,
		product.Code, product.Price, product.Stock)

	return nil
}

func CreateProductHandler(cmd *cobra.Command, args []string) error {
	if len(args) != 3 {
		return error_handler.GetErrorMessage("CreateProductRequirement", logger.Error)
	}

	productCode := args[0]
	productPrice, _ := strconv.ParseFloat(args[1], 64)
	productStock, _ := strconv.Atoi(args[2])

	productManager := &product.ProductManager{}
	createdProduct, err := productManager.CreateProduct(productCode, productPrice, productStock)
	if err != nil {
		return err
	}

	fmt.Printf(`Product created; code %s, price %.2f, stock %d`,
		createdProduct.Code, createdProduct.Price, createdProduct.Stock)

	return nil
}
