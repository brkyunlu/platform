package cmd

import (
	"github.com/spf13/cobra"
	"platform/controllers/handlers/increase_time"
)

func init() {
	Command.AddCommand(&cobra.Command{
		Use:   "increase_time [time(int)]",
		Short: "Increase the system time",
		Args:  cobra.ExactArgs(1),
		RunE:  increase_time.IncreaseTimeHandler,
	})
}
