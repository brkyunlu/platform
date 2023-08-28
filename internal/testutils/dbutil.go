// testutils/dbutil.go

package testutils

import (
	"errors"
	"fmt"
	"github.com/joho/godotenv"
	"platform/internal/client"
	"platform/internal/logger"
	"platform/internal/models"
)

var EnvFileError = errors.New("error loading .env file. Please make sure you are working in the correct directory")

func SetupTestDB() {
	// Load environment file
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Print(EnvFileError)
	}
	logger.GetLogger().Info("App Started")
	client.Connections()
	models.Migrate()
}
