package worker

import (
	"sync"
	"time"

	"github.com/antihax/optional"
	"github.com/conductor-sdk/conductor-go/sdk/model"
	log "github.com/sirupsen/logrus"
)

type WorkerOrkestrator struct {
	taskName string

	pollInterval      time.Duration
	pollIntervalMutex sync.RWMutex

	executeTaskFunction      model.ExecuteTaskFunction
	executeTaskFunctionMutex sync.RWMutex

	domain      optional.String
	domainMutex sync.RWMutex

	batchSizeLimit      int
	batchSizeLimitMutex sync.RWMutex

	runningWorkers      int
	runningWorkersMutex sync.RWMutex
}

func NewWorkerOrkestrator(
	taskName string,
	batchSizeLimit int,
	pollInterval time.Duration,
	executeTaskFunction model.ExecuteTaskFunction,
	domain optional.String,
) *WorkerOrkestrator {
	return &WorkerOrkestrator{
		taskName:            taskName,
		pollInterval:        pollInterval,
		executeTaskFunction: executeTaskFunction,
		domain:              domain,
		batchSizeLimit:      batchSizeLimit,
		runningWorkers:      0,
	}
}

func (wo *WorkerOrkestrator) GetPollInterval() (pollInterval time.Duration, err error) {
	wo.pollIntervalMutex.RLock()
	defer wo.pollIntervalMutex.RUnlock()
	return wo.pollInterval, nil
}

func (wo *WorkerOrkestrator) SetPollInterval(pollInterval time.Duration) (err error) {
	wo.pollIntervalMutex.Lock()
	defer wo.pollIntervalMutex.Unlock()
	previous := wo.pollInterval
	wo.pollInterval = pollInterval
	log.Debug(
		"Updated poll interval for task: ", wo.taskName,
		", from: ", previous,
		", to: ", wo.pollInterval,
	)
	return nil
}

func (wo *WorkerOrkestrator) GetExecuteTaskFunction() (executeTaskFunction model.ExecuteTaskFunction, err error) {
	wo.executeTaskFunctionMutex.RLock()
	defer wo.executeTaskFunctionMutex.RUnlock()
	return wo.executeTaskFunction, nil
}

func (wo *WorkerOrkestrator) SetExecuteTaskFunction(executeTaskFunction model.ExecuteTaskFunction) (err error) {
	wo.executeTaskFunctionMutex.Lock()
	defer wo.executeTaskFunctionMutex.Unlock()
	wo.executeTaskFunction = executeTaskFunction
	return nil
}

func (wo *WorkerOrkestrator) GetDomain() (domain optional.String, err error) {
	wo.domainMutex.RLock()
	defer wo.domainMutex.RUnlock()
	return wo.domain, nil
}

func (wo *WorkerOrkestrator) SetDomain(domain optional.String) (err error) {
	wo.domainMutex.Lock()
	defer wo.domainMutex.Unlock()
	previous := wo.domain
	wo.domain = domain
	log.Debug(
		"Updated domain for task: ", wo.taskName,
		", from: ", previous,
		", to: ", wo.domain,
	)
	return nil
}

func (wo *WorkerOrkestrator) GetAvailableWorkers() (availableWorkers int, err error) {
	batchSizeLimit, err := wo.GetBatchSizeLimit()
	if err != nil {
		return -1, err
	}
	runningWorkers, err := wo.getRunningWorkers()
	if err != nil {
		return -1, err
	}
	return batchSizeLimit - runningWorkers, nil
}

func (wo *WorkerOrkestrator) getRunningWorkers() (runningWorkers int, err error) {
	wo.runningWorkersMutex.RLock()
	defer wo.runningWorkersMutex.RUnlock()
	return wo.runningWorkers, nil
}

func (wo *WorkerOrkestrator) IncreaseRunningWorkers(quantity int) (err error) {
	wo.runningWorkersMutex.Lock()
	defer wo.runningWorkersMutex.Unlock()
	previous := wo.runningWorkers
	wo.runningWorkers += quantity
	log.Trace(
		"Increased running workers for task: ", wo.taskName,
		", from: ", previous,
		", to: ", wo.runningWorkers,
	)
	return nil
}

func (wo *WorkerOrkestrator) DecreaseRunningWorker() (err error) {
	wo.runningWorkersMutex.Lock()
	defer wo.runningWorkersMutex.Unlock()
	previous := wo.runningWorkers
	wo.runningWorkers -= 1
	log.Trace(
		"Decreased running workers for task: ", wo.taskName,
		", from: ", previous,
		", to: ", wo.runningWorkers,
	)
	return nil
}

func (wo *WorkerOrkestrator) IncreaseBatchSizeLimitForTask(batchSize int) (err error) {
	wo.batchSizeLimitMutex.Lock()
	defer wo.batchSizeLimitMutex.Unlock()
	previous := wo.batchSizeLimit
	wo.batchSizeLimit += batchSize
	log.Debug(
		"Increased batchSize for taskName: ", wo.taskName,
		", from: ", previous,
		", to: ", wo.batchSizeLimit,
	)
	return nil
}

func (wo *WorkerOrkestrator) DecreaseBatchSizeLimit(batchSize int) (err error) {
	wo.batchSizeLimitMutex.Lock()
	defer wo.batchSizeLimitMutex.Unlock()
	previous := wo.batchSizeLimit
	wo.batchSizeLimit -= batchSize
	log.Debug(
		"Decreased batchSize for taskName: ", wo.taskName,
		", from: ", previous,
		", to: ", wo.batchSizeLimit,
	)
	return nil
}

func (wo *WorkerOrkestrator) GetBatchSizeLimit() (batchSizeLimit int, err error) {
	wo.batchSizeLimitMutex.RLock()
	defer wo.batchSizeLimitMutex.RUnlock()
	return wo.batchSizeLimit, nil
}

func (wo *WorkerOrkestrator) SetBatchSizeLimit(batchSize int) (err error) {
	wo.batchSizeLimitMutex.Lock()
	defer wo.batchSizeLimitMutex.Unlock()
	previous := wo.batchSizeLimit
	wo.batchSizeLimit = batchSize
	log.Debug(
		"Updated batch size for taskName: ", wo.taskName,
		", from: ", previous,
		", to: ", batchSize,
	)
	return nil
}
