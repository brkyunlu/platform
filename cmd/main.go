package main

import (
	"errors"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
	"platform/cmd/commands"
	"platform/internal/client"
	"platform/internal/database/seeders"
	"platform/internal/models"
	"time"
)

var EnvFileError = errors.New("error loading .env file. Please make sure you are working in the correct directory")

func main() {

	// Load environment file
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Print(EnvFileError)
	}
	client.Connections()
	models.Migrate()
	seeders.Seed()
	// Başlangıç saati
	startTime := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.Now().Location())

	timeManager := &commands.TimeManager{CurrentTime: startTime}
	RunCommandProcessor(timeManager)
}
func RunCommandProcessor(timeManager *commands.TimeManager) {
	commandProcessor := commands.NewCommandProcessor(timeManager)
	commandProcessor.Run()
}
