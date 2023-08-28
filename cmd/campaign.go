package cmd

import (
	"github.com/spf13/cobra"
	"platform/controllers/handlers/campaign"
)

func init() {
	Command.AddCommand(&cobra.Command{
		Use:   "get_campaign_info [campaign_name]",
		Short: "Get information about a campaign",
		Args:  cobra.ExactArgs(1),
		RunE:  campaign.GetCampaignInfoHandler,
	})
	Command.AddCommand(&cobra.Command{
		Use:   "create_campaign [campaign_name] [product_code] [campaign_duration] [campaign_price_limit] [campaign_target_sales]",
		Short: "Create a new campaign",
		Args:  cobra.ExactArgs(5),
		RunE:  campaign.CreateCampaignHandler,
	})
}
