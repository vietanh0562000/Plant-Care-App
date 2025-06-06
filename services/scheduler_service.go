package services

import (
	"log"
	"plant-care-app/database"
	"plant-care-app/models"

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

		for i := 0; i < len(plants); i++ {
			mailService := MailService{}
			mailService.SendNewMail("aloalo", "vitanh562000@gmail.com")
		}
	})
	sch.c.Start() // Start the scheduler
}
