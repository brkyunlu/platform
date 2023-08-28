package increase_time

import (
	"fmt"
	"github.com/spf13/cobra"
	"platform/controllers/increase_time"
	"platform/internal/error_handler"
	"platform/internal/logger"
	"strconv"
)

func IncreaseTimeHandler(cmd *cobra.Command, args []string) error {
	if len(args) != 1 {
		return error_handler.GetErrorMessage("IncreaseTimeRequirement", logger.Error)
	}
	hours, _ := strconv.Atoi(args[0])
	increaseTimeManager := &increase_time.IncreaseTimeManager{}
	increased_time, err := increaseTimeManager.IncreaseTime(hours)
	if err != nil {
		return err
	}
	fmt.Printf(`Time is: %s`, increased_time)
	return nil
}
