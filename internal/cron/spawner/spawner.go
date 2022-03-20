package cron

import (
    "github.com/go-co-op/gocron"
)

// Wrapper. Needed for user to be able to specify multiple jobs. Each with it's own function and interval
type Job struct {
    JobFunc interface{}
    Interval *gocron.Scheduler
}

// Inits and starts the jobs scheduler
func Spawn(scheduler *gocron.Scheduler, jobs ...Job) {
    for _, job := range jobs {
        job.Interval.Do(job.JobFunc)
    }
    scheduler.StartAsync()
}
