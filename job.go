package main

import "sync"

type Priority uint8
type Status uint8

const (
	Low    Priority = 1
	Medium          = 2
	High            = 3
)

const (
	Created  Status = 1
	Pending         = 2
	Approved        = 3
	Rejected        = 4
)

type Job struct {
	Name     string
	Action   func() error
	Priority Priority
	Status   Status
}

func (job *Job) Execute(wg *sync.WaitGroup) {
	job.Status = Created
	wg.Add(1)
	job.Status = Pending
	defer wg.Done()
	result := job.Action()
	if result != nil {
		job.Status = Rejected
		return
	}
	job.Status = Approved
}
