package main

import (
	"errors"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
	"net/http"
	"os"
	"platform/cmd"
	"platform/internal/client"
	"platform/internal/logger"
	"platform/internal/models"
	time_simulation "platform/internal/time"
	"strconv"
	"time"
)

var EnvFileError = errors.New("error loading .env file. Please make sure you are working in the correct directory")

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Print(EnvFileError)
	}
	appPort, _ := strconv.Atoi(os.Getenv("APP_PORT"))
	logger.InitLogger()
	client.Connections()
	models.Migrate()

	var timeSimulator time_simulation.TimeSimulator = time_simulation.DefaultTimeSimulator{}
	timeSimulator.GetSimulatedTime()

	cmd.Execute()

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", appPort), //burayı bi kaldırıp dene bakalım
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	err = srv.ListenAndServe()
	logger.GetLogger().Error(err)
}
