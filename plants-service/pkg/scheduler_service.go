package services

import (
	"log"
	"plant-care-app/plants-service/internal/database"
	"plant-care-app/plants-service/internal/models"

	"github.com/robfig/cron"
)

type Scheduler struct {
	c *cron.Cron
}

func (sch *Scheduler) CreateSchedulerAt(hour int) {
	sch.c = cron.New()
	sch.c.AddFunc("@every 20s", func() {
		var plants []models.Plant
		result := database.DB.Find(&plants).Where("last_time_watering + (watering_interval || ' days) <= NOW()")
		if result.Error != nil {
			log.Printf("Database error: %v\n", result.Error)
			return
		}

	})
	sch.c.Start() // Start the scheduler
}
