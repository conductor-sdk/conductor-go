package worker

import (
	"fmt"
	"sync"
	"time"

	"github.com/conductor-sdk/conductor-go/pkg/concurrency"
	log "github.com/sirupsen/logrus"
)

type AvailableWorkerChannel chan int

const (
	availableWorkerChannelTimeout = 30 * time.Second
)

type AvailableWorkerMonitor struct {
	taskType                   string
	availableWorkerAmount      int
	availableWorkerAmountMutex sync.RWMutex
	availableWorkerChannel     AvailableWorkerChannel
	runningGoRoutines          sync.WaitGroup
}

func NewAvailableWorkerMonitor(taskType string) *AvailableWorkerMonitor {
	return &AvailableWorkerMonitor{
		taskType:               taskType,
		availableWorkerAmount:  0,
		availableWorkerChannel: make(AvailableWorkerChannel),
	}
}

func (a *AvailableWorkerMonitor) Start() error {
	a.runningGoRoutines.Add(1)
	go a.mergeNextAvailableWorkerAmountDaemon()
	log.Debug("Started available worker monitor of taskType: ", a.taskType)
	return nil
}

func (a *AvailableWorkerMonitor) Wait() error {
	log.Debug("Waiting for available worker monitor of taskType: ", a.taskType)
	a.runningGoRoutines.Wait()
	log.Debug("Done waiting for available worker monitor of taskType: ", a.taskType)
	return nil
}

func (a *AvailableWorkerMonitor) IncreaseAvailableWorkerAmount(amount int) error {
	a.availableWorkerChannel <- amount
	return nil
}

func (a *AvailableWorkerMonitor) DecreaseAvailableWorkerAmount() error {
	a.availableWorkerChannel <- -1
	return nil
}

func (a *AvailableWorkerMonitor) GetAvailableWorkerAmount() (int, error) {
	a.availableWorkerAmountMutex.RLock()
	defer a.availableWorkerAmountMutex.RUnlock()
	return a.availableWorkerAmount, nil
}

func (a *AvailableWorkerMonitor) mergeNextAvailableWorkerAmountDaemon() {
	defer a.runningGoRoutines.Done()
	defer concurrency.HandlePanicError("merge_next_available_worker")
	for {
		err := a.mergeNextAvailableWorkerAmount()
		if err != nil {
			log.Warning(
				"Failed to merge next available worker amount. Reason: ", err.Error(),
				", taskType: ", a.taskType,
			)
			break
		}
	}
}

func (a *AvailableWorkerMonitor) mergeNextAvailableWorkerAmount() error {
	availableWorkerAmount, err := a.getNextValueFromChannel()
	if err != nil {
		return err
	}
	log.Debug(
		"Received available worker amount from channel: ", availableWorkerAmount,
		", taskType: ", a.taskType,
	)
	return a.updateAvailableWorkerAmount(availableWorkerAmount)
}

func (a *AvailableWorkerMonitor) getNextValueFromChannel() (int, error) {
	select {
	case availableWorkers, ok := <-a.availableWorkerChannel:
		if !ok {
			return 0, fmt.Errorf("channel closed")
		}
		return availableWorkers, nil
	case <-time.After(availableWorkerChannelTimeout):
		return 0, fmt.Errorf("timeout")
	}
}

func (a *AvailableWorkerMonitor) updateAvailableWorkerAmount(amount int) error {
	a.availableWorkerAmountMutex.Lock()
	defer a.availableWorkerAmountMutex.Unlock()
	a.availableWorkerAmount += amount
	log.Debug(
		"Updated available worker amount by: ", amount,
		", taskType: ", a.taskType,
	)
	return nil
}
