package services

import (
	"fmt"

	"github.com/robfig/cron"
)

type Scheduler struct {
	c *cron.Cron
}

func (sch *Scheduler) CreateSchedulerAt(hour int, actionFunc func()) {
	sch.c = cron.New()
	runRule := fmt.Sprintf("0 0 */%d * * *", hour)
	sch.c.AddFunc(runRule, actionFunc)
	sch.c.Start() // Start the scheduler
}
