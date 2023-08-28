package increase_time

import (
	"fmt"
	"platform/controllers/campaign"
	"platform/internal/error_handler"
	"platform/internal/logger"
	"platform/internal/models"
	time_simulation "platform/internal/time"
	"sync"
	"time"
)

type IncreaseTimeManager struct {
	CampaignManager campaign.CampaignManager
	TimeManager     time_simulation.DefaultTimeSimulator
}

func (m *IncreaseTimeManager) IncreaseTime(hours int) (time.Time, error) {
	currentTime := m.TimeManager.GetSimulatedTime()
	// Get the current simulated time
	newTime := currentTime.Add(time.Duration(hours) * time.Hour)
	// Calculate the new simulated time by adding the specified hours

	// Update the new simulated time in the config.json file
	err := m.TimeManager.UpdateSimulatedTime(newTime)
	if err != nil {
		return currentTime, error_handler.GetErrorMessage("Error updating simulated time in ErrSimulatedTimeUpdate", logger.Error)
	}
	// Get all campaigns
	allCampaigns, err := models.GetAllCampaigns()
	if err != nil {
		return currentTime, error_handler.GetErrorMessage("Error getting all campaigns in ErrGetAllCampaign", logger.Error)
	}

	// Update campaign statuses in parallel and wait for the result
	var waitGroup sync.WaitGroup
	waitGroup.Add(len(allCampaigns))
	for _, c := range allCampaigns {
		go func(campaign models.Campaign) {
			defer waitGroup.Done()
			err := m.CampaignManager.UpdateCampaignStatus(&campaign)
			if err != nil {
				fmt.Println(err)
			}
		}(c)
	}
	waitGroup.Wait() // Wait for all operations to complete

	return newTime, nil
}
