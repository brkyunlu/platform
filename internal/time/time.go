package time_simulated

import (
	"encoding/json"
	"io/ioutil"
	"time"
)

// TimeSimulator interface defines functions to simulate time.
type TimeSimulator interface {
	GetSimulatedTime() time.Time
	UpdateSimulatedTime(newTime time.Time) error
}

// SimulatedTimeConfig represents the simulated time stored in the config.json file.
type SimulatedTimeConfig struct {
	SimulatedTime string `json:"simulated_time"`
}

// DefaultTimeSimulator provides the default time simulation.
type DefaultTimeSimulator struct{}

// readSimulatedTimeFromConfigFile reads the simulated time from the config.json file.
func readSimulatedTimeFromConfigFile() (time.Time, error) {
	// Read the config file
	configData, err := ioutil.ReadFile("config.json")
	if err != nil {
		return time.Time{}, err
	}

	var config SimulatedTimeConfig

	// Parse the JSON data
	err = json.Unmarshal(configData, &config)
	if err != nil {
		return time.Time{}, err
	}
	if config.SimulatedTime == "" {
		// If simulated_time is empty in the config file, return today's date at 00:00
		return time.Now().Truncate(24 * time.Hour), nil
	}
	// Parse the simulated time
	simulatedTime, err := time.Parse(time.RFC3339, config.SimulatedTime)
	if err != nil {
		return time.Time{}, err
	}

	return simulatedTime, nil
}

// GetSimulatedTime returns the default simulated time.
func (d DefaultTimeSimulator) GetSimulatedTime() time.Time {
	// If the config.json file exists and contains a simulated time, return that time.
	// Otherwise, return today's date at 00:00.
	simulatedTime, err := readSimulatedTimeFromConfigFile()
	if err != nil {
		return time.Now().Truncate(24 * time.Hour)
	}
	return simulatedTime
}

// UpdateSimulatedTime updates the simulated time in the config.json file.
func (d DefaultTimeSimulator) UpdateSimulatedTime(newTime time.Time) error {
	configPath := "config.json"

	// Read the config file
	configData, err := ioutil.ReadFile(configPath)
	if err != nil {
		return err
	}

	var config SimulatedTimeConfig

	// Parse the JSON data
	err = json.Unmarshal(configData, &config)
	if err != nil {
		return err
	}

	// Update the simulated time with the new time
	config.SimulatedTime = newTime.Format(time.RFC3339)

	// Encode the updated JSON data
	updatedConfigData, err := json.MarshalIndent(config, "", "    ")
	if err != nil {
		return err
	}

	// Update the config file
	err = ioutil.WriteFile(configPath, updatedConfigData, 0644)
	if err != nil {
		return err
	}

	//fmt.Printf("Simulated time updated: %v\n", newTime.Format(time.RFC3339))
	return nil
}
