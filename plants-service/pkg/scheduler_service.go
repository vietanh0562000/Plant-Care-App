package services

import (
	"github.com/robfig/cron"
)

type Scheduler struct {
	c *cron.Cron
}

func (sch *Scheduler) CreateSchedulerAt(rule string, actionFunc func()) {
	sch.c = cron.New()
	sch.c.AddFunc(rule, actionFunc)
	sch.c.Start() // Start the scheduler
}
