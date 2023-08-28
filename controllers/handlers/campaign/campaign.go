package campaign

import (
	"fmt"
	"github.com/spf13/cobra"
	"platform/controllers/campaign"
	"platform/internal/error_handler"
	"platform/internal/logger"
	"strconv"
)

func GetCampaignInfoHandler(cmd *cobra.Command, args []string) error {
	if len(args) != 1 {
		return error_handler.GetErrorMessage("GetCampaignRequirement", logger.Error)
	}

	campaignName := args[0]

	campaignManager := &campaign.CampaignManager{}
	campaignInfo, err := campaignManager.GetCampaignInfo(campaignName)
	if err != nil {
		return err
	}
	fmt.Printf(`Campaign "%s" information; Status %t, Target Sales %d, Total Sales %d`,
		campaignInfo.Name, campaignInfo.Status, campaignInfo.TargetSales, campaignInfo.TotalSales)

	return nil
}

func CreateCampaignHandler(cmd *cobra.Command, args []string) error {
	if len(args) != 5 {
		return error_handler.GetErrorMessage("CreateCampaignRequirement", logger.Error)
	}

	campaignName := args[0]
	productCode := args[1]
	duration, _ := strconv.Atoi(args[2])
	priceLimit, _ := strconv.ParseFloat(args[3], 64)
	targetSales, _ := strconv.Atoi(args[4])
	campaignManager := &campaign.CampaignManager{}
	createdCampaign, err := campaignManager.CreateCampaign(campaignName, productCode, duration, priceLimit, targetSales)
	if err != nil {
		return err
	}

	fmt.Printf(`Campaign created; name %s, product %s, duration %d hours, limit %.2f, target sales %d`,
		createdCampaign.Name, createdCampaign.Product.Code, createdCampaign.Duration, createdCampaign.PriceManipulationLimit, createdCampaign.TargetSales)

	return nil
}
