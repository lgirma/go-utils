package utils

import (
	"fmt"
	"sync"

	"github.com/google/uuid"
)

type JobCompleteCallback = func(jobId string, result any, err error)
type JobAction = func() (result any, err error)

type Job struct {
	Id       string
	Action   JobAction
	Callback JobCompleteCallback
}

const (
	JOB_STATUS_WORKING = iota
	JOB_STATUS_COMPLETE
)

type JobService interface {
	Start(job *Job) error
	Status(id string) int
	Stop(id string)
	Wait(id string)
}

func GetJobService() JobService {
	return &DefaultJobService{
		_jobs:       make(map[string]*Job),
		_waitGroups: make(map[string]*sync.WaitGroup),
	}
}

func NewJob(action JobAction, callback JobCompleteCallback) *Job {
	return &Job{
		Id:       uuid.NewString(), /*TODO: generate new uuid*/
		Action:   action,
		Callback: callback,
	}
}

type DefaultJobService struct {
	_jobs       map[string]*Job
	_waitGroups map[string]*sync.WaitGroup
}

func (service *DefaultJobService) Wait(id string) {
	service._waitGroups[id].Wait()
}

func (service *DefaultJobService) Start(job *Job) error {
	if job == nil {
		return fmt.Errorf("job is nil")
	} else if job.Action == nil {
		return fmt.Errorf("job action is nil")
	}
	service._jobs[job.Id] = job
	var wg sync.WaitGroup
	wg.Add(1)
	service._waitGroups[job.Id] =  &wg
	
	go func() {
		defer wg.Done()
		result, err := job.Action()
		if job.Callback != nil {
			job.Callback(job.Id, result, err)
		}
		service.Stop(job.Id)
	}()
	return nil
}

func (service *DefaultJobService) Status(id string) int {
	_, ok := service._jobs[id]
	if !ok {
		return JOB_STATUS_COMPLETE
	}
	return JOB_STATUS_WORKING
}

func (service *DefaultJobService) Stop(id string) {
	delete(service._jobs, id)
	delete(service._waitGroups, id)
}
