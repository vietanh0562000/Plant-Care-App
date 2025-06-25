package services

import (
	"github.com/robfig/cron"
)

type Scheduler struct {
	c *cron.Cron
}

func (sch *Scheduler) CreateSchedulerAt(hour int, actionFunc func()) {
	sch.c = cron.New()
	sch.c.AddFunc("@every 20s", actionFunc)
	sch.c.Start() // Start the scheduler
}
