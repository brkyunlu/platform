package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var Command = &cobra.Command{
	Use:   "Platform",
	Short: "Platform",
	Long:  "Platform Application",
}

func Execute() {
	if err := Command.Execute(); err != nil {
		fmt.Println(err)

		os.Exit(1)
	}
}

func init() {
}
