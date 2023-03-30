package utils

import (
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewJob(t *testing.T) {
	job := NewJob(func() (result any, err error) {
		return nil, nil
	}, func(jobId string, result any, err error) {

	})

	assert.NotEmpty(t, job.Id)
}

func TestStartJob(t *testing.T) {
	var service = GetJobService()
	count := 0
	job := NewJob(func() (result any, err error) {
		count += 1
		return 1, nil
	}, func(jobId string, result any, err error) {
		count += 1
		fmt.Printf("completed with error: %v", err)
	})

	err := service.Start(job)
	service.Wait(job.Id)

	assert.Nil(t, err)
	assert.Equal(t, 2, count)
}

func TestStartJobWithResultAndError(t *testing.T) {
	var service = GetJobService()
	job := NewJob(func() (result any, err error) {
		return "RES", errors.New("ERR_JOB")
	}, func(jobId string, result any, err error) {
		assert.Equal(t, "RES", result)
		assert.Equal(t, "ERR_JOB", err.Error())
	})

	err := service.Start(job)
	service.Wait(job.Id)

	assert.Nil(t, err)
}

func TestJobStatus(t *testing.T) {
	var service = GetJobService()
	job := NewJob(func() (result any, err error) {
		time.Sleep(3 * time.Second)
		return nil, nil
	}, func(jobId string, result any, err error) {

	})

	err := service.Start(job)
	assert.Nil(t, err)

	assert.Equal(t, JOB_STATUS_WORKING, service.Status(job.Id))
	service.Wait(job.Id)
	assert.Equal(t, JOB_STATUS_COMPLETE, service.Status(job.Id))
}
