//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
//  the License. You may obtain a copy of the License at
//
//  http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
//  an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
//  specific language governing permissions and limitations under the License.

package worker

import (
	"fmt"
	"sync"
	"time"

	"github.com/conductor-sdk/conductor-go/sdk/model"
	log "github.com/sirupsen/logrus"
)

type WorkerProperties struct {
	TaskName        string
	ExecuteFunction model.ExecuteTaskFunction

	batchSize            int
	paused               bool
	pollInterval         time.Duration
	runningWorkerCounter int
	taskDomain           string

	batchSizeMutex            sync.RWMutex
	pausedMutex               sync.RWMutex
	pollIntervalMutex         sync.RWMutex
	runningWorkerCounterMutex sync.RWMutex
	taskDomainMutex           sync.RWMutex
}

func NewWorkerProperties(taskName string, executeFunction model.ExecuteTaskFunction) *WorkerProperties {
	return NewWorkerPropertiesCustom(taskName, executeFunction, 1, 100*time.Millisecond)
}

func NewWorkerPropertiesCustom(taskName string, executeFunction model.ExecuteTaskFunction, batchSize int, pollInterval time.Duration) *WorkerProperties {
	return &WorkerProperties{
		TaskName:             taskName,
		ExecuteFunction:      executeFunction,
		batchSize:            batchSize,
		runningWorkerCounter: 0,
		pollInterval:         pollInterval,
		paused:               false,
	}
}

func (w *WorkerProperties) GetBatchSize() int {
	w.batchSizeMutex.RLock()
	defer w.batchSizeMutex.RUnlock()
	return w.batchSize
}

func (w *WorkerProperties) SetBatchSize(batchSize int) error {
	if batchSize < 0 {
		return fmt.Errorf("batchSize can not be negative")
	}
	if batchSize == w.GetBatchSize() {
		return nil
	}
	w.batchSizeMutex.Lock()
	defer w.batchSizeMutex.Unlock()
	previousValue := w.batchSize
	w.batchSize = batchSize
	if batchSize == 0 {
		log.Info("Stopped worker for task: ", w.TaskName)
	} else if previousValue == 0 {
		log.Info("Started worker for task: ", w.TaskName)
	}
	return nil
}

func (w *WorkerProperties) Resume() {
	w.pausedMutex.Lock()
	defer w.pausedMutex.Unlock()
	w.paused = false
	log.Info("Resumed worker for task: ", w.TaskName)
}

func (w *WorkerProperties) Pause() {
	w.pausedMutex.Lock()
	defer w.pausedMutex.Unlock()
	w.paused = true
	log.Info("Paused worker for task: ", w.TaskName)
}

func (w *WorkerProperties) IsPaused() bool {
	w.pausedMutex.RLock()
	defer w.pausedMutex.RUnlock()
	return w.paused
}

func (w *WorkerProperties) GetPollInterval() time.Duration {
	w.pollIntervalMutex.RLock()
	defer w.pollIntervalMutex.RUnlock()
	return w.pollInterval
}

func (w *WorkerProperties) SetPollInterval(pollInterval time.Duration) {
	w.pollIntervalMutex.Lock()
	defer w.pollIntervalMutex.Unlock()
	previousValue := w.pollInterval
	w.pollInterval = pollInterval
	log.Trace(
		"Updated pollInterval for task: ", w.TaskName,
		", from: ", previousValue.Milliseconds(), "ms",
		", to: ", w.pollInterval.Milliseconds(), "ms",
	)
}

func (w *WorkerProperties) GetTaskDomain() string {
	w.taskDomainMutex.RLock()
	defer w.taskDomainMutex.RUnlock()
	return w.taskDomain
}

func (w *WorkerProperties) SetTaskDomain(taskDomain string) {
	w.taskDomainMutex.Lock()
	defer w.taskDomainMutex.Unlock()
	previousValue := w.taskDomain
	w.taskDomain = taskDomain
	log.Trace(
		"Updated taskDomain for task: ", w.TaskName,
		", from: ", previousValue,
		", to: ", w.taskDomain,
	)
}

func (w *WorkerProperties) GetRunningWorkers() int {
	w.runningWorkerCounterMutex.RLock()
	defer w.runningWorkerCounterMutex.RUnlock()
	return w.runningWorkerCounter
}

func (w *WorkerProperties) IncrementRunningWorker(amount int) {
	w.runningWorkerCounterMutex.Lock()
	defer w.runningWorkerCounterMutex.Unlock()
	previousValue := w.runningWorkerCounter
	w.runningWorkerCounter += amount
	log.Trace(
		"Increased running workers for task: ", w.TaskName,
		", from: ", previousValue,
		", to: ", w.runningWorkerCounter,
	)
}

func (w *WorkerProperties) DecreaseRunningWorker() {
	w.runningWorkerCounterMutex.Lock()
	defer w.runningWorkerCounterMutex.Unlock()
	previousValue := w.runningWorkerCounter
	w.runningWorkerCounter -= 1
	log.Trace(
		"Decreased runningWorkerCounter for task: ", w.TaskName,
		", from: ", previousValue,
		", to: ", w.runningWorkerCounter,
	)
}
