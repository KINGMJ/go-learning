package main

import (
	"fmt"
	"sync"
	"time"
)

type Job struct {
	Name     string
	Schedule string
	Action   func()
}

type Scheduler struct {
	jobs      []Job
	stopChan  chan struct{}
	waitGroup sync.WaitGroup
}

func NewScheduler() *Scheduler {
	return &Scheduler{
		stopChan: make(chan struct{}),
	}
}

func (s *Scheduler) AddJob(job Job) {
	s.jobs = append(s.jobs, job)
}

func (s *Scheduler) Run() {
	for _, job := range s.jobs {
		s.scheduleJob(job)
	}
	<-s.stopChan
}

func (s *Scheduler) Stop() {
	close(s.stopChan)
	s.waitGroup.Wait()
}

func (s *Scheduler) scheduleJob(job Job) {
	s.waitGroup.Add(1)
	go func() {
		defer s.waitGroup.Done()
		for {
			select {
			case <-time.After(parseDuration(job.Schedule)):
				fmt.Printf("Running job '%s' at %s\n", job.Name, time.Now().Format(time.RFC3339))
				job.Action()
			case <-s.stopChan:
				return
			}
		}
	}()
}

func parseDuration(schedule string) time.Duration {
	duration, err := time.ParseDuration(schedule)
	if err != nil {
		panic(err)
	}
	return duration
}

func main() {
	scheduler := NewScheduler()
	scheduler.AddJob(Job{
		Name:     "ExampleJob",
		Schedule: "5s",
		Action: func() {
			fmt.Println("Hello from ExampleJob!")
		},
	})

	go scheduler.Run()

	select {}
	// time.Sleep(20*time.S)
}

// 使用 channel、for、select 实现 crontab 功能
