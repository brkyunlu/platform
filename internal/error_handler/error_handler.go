package error_handler

import (
	"errors"
	"platform/internal/logger"
)

var errorMessages = map[string]string{
	"ErrProductNotFound":        "Product not found.",
	"ErrProductCreate":          "An error occurred while creating the product.",
	"ErrProductQuery":           "An error occurred while querying the product.",
	"ErrProductUpdate":          "An error occurred while updating the product status.",
	"ErrActiveCampaignExists":   "An active campaign already exists for this product.",
	"ErrCampaignNotFound":       "Campaign not found.",
	"ErrCampaignUpdate":         "An error occurred while updating the campaign status.",
	"ErrCampaignCreate":         "An error occurred while creating the campaign.",
	"ErrCampaignValidation":     "Campaign validation error.",
	"ErrSimulatedTimeUpdate":    "Time update error.",
	"ErrGetAllCampaign":         "Unable to fetch campaigns.",
	"ErrNoStock":                "Insufficient stock.",
	"ErrStockUpdate":            "An error occurred while updating stock status.",
	"ErrOrderCreate":            "An error occurred while creating the order.",
	"GetCampaignRequirement":    "Campaign name is required.",
	"CreateCampaignRequirement": "Campaign name, product code, duration, manipulation limit, and target sales are required.",
	"IncreaseTimeRequirement":   "Hour is required (int).",
	"OrderRequirement":          "Product code and quantity information are required.",
	"GetProductRequirement":     "Product code is required.",
	"CreateProductRequirement":  "Product code, price, and stock are required.",
}

func GetErrorMessage(err string, level logger.LogLevel) error {
	if errMsg, ok := errorMessages[err]; ok {
		log := logger.GetLogger()

		switch level {
		case logger.Info:
			log.Info(errMsg)
		case logger.Warn:
			log.Warn(errMsg)
		case logger.Error:
			log.Error(errMsg)
			// ... Other log levels can be added here
		}
		return errors.New(errMsg)
	}
	log := logger.GetLogger()
	log.Error("Unknown error code:", err)
	return errors.New("An unexpected error occurred.")
}
