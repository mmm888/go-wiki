package cron

import "github.com/robfig/cron"

type Schedule int

const (
	Minutely Schedule = iota + 1
	Hourly
	Daily
)

func NewCron() *Cron {
	return &Cron{
		c: cron.New(),
	}
}

type Cron struct {
	c *cron.Cron
}

func (c *Cron) Start() {
	c.c.Start()
}

func (c *Cron) Add(schedule Schedule, f func()) {
	switch schedule {
	case Minutely:
		c.c.AddFunc("0 * * * * *", f)

	case Hourly:
		c.c.AddFunc("0 0 * * * *", f)

	case Daily:
		c.c.AddFunc("0 0 0 * * *", f)
	}
}
