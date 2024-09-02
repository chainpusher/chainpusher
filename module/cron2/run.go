package cron2

import "github.com/robfig/cron"

func Imediate(cron *cron.Cron) {
	for _, entry := range cron.Entries() {
		entry.Job.Run()
	}
}

func ImediateRun(cron *cron.Cron) {
	Imediate(cron)
	cron.Run()
}

func ImediateStart(cron *cron.Cron) {
	Imediate(cron)
	cron.Start()
}
