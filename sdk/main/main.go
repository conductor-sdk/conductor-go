package main

import (
	"fmt"
	"github.com/conductor-sdk/conductor-go/sdk/client"
	"github.com/conductor-sdk/conductor-go/sdk/model"
	"github.com/conductor-sdk/conductor-go/sdk/settings"
	"github.com/conductor-sdk/conductor-go/sdk/worker"
)

type Restaurant struct {
	Restaurant RestaurantData `json:"restaurant"`
}

type RestaurantData struct {
	Name  string `json:"name"`
	Owner Owner  `json:"owner"`
}

type Owner struct {
	Name string `json:"name"`
}

func authSettings() *settings.AuthenticationSettings {
	key := "api_key_user_03"
	secret := "api_key_user_03"
	if key != "" && secret != "" {
		return settings.NewAuthenticationSettings(
			key,
			secret,
		)
	}

	return nil
}

func httpSettings() *settings.HttpSettings {
	url := "http://localhost:8080/api"
	return settings.NewHttpSettings(url)
}

type Data struct {
	InputKey InputKey `json:"inputKey"`
}

type InputKey struct {
	Headers      Headers `json:"headers"`
	ReasonPhrase string  `json:"reasonPhrase"`
	Body         Body    `json:"body"`
	StatusCode   int     `json:"statusCode"`
}

type Headers struct {
	Date                    []string `json:"Date"`
	ContentType             []string `json:"Content-Type"`
	ContentLength           []string `json:"Content-Length"`
	Connection              []string `json:"Connection"`
	StrictTransportSecurity []string `json:"Strict-Transport-Security"`
}

type Body struct {
	RandomString   string                 `json:"randomString"`
	RandomInt      int                    `json:"randomInt"`
	HostName       string                 `json:"hostName"`
	APIRandomDelay string                 `json:"apiRandomDelay"`
	SleepFor       string                 `json:"sleepFor"`
	StatusCode     string                 `json:"statusCode"`
	QueryParams    map[string]interface{} `json:"queryParams"`
}

func ProcessBody(body *Data) (*Data, error) {
	if body == nil {
		fmt.Println("Body is nil")
		return nil, nil
	}
	body.InputKey.Body.HostName = "locally hosting"
	fmt.Println("returning", body)
	return body, nil
}

func ProcessBody2(task *model.Task) (interface{}, error) {
	fmt.Println("executing", task.TaskId)
	return Data{InputKey: InputKey{
		Headers:      Headers{},
		ReasonPhrase: "",
		Body:         Body{},
		StatusCode:   0,
	}}, nil
}

func ProcessBody3(ctx worker.TaskContext, data *Data) (Data, error) {

	fmt.Println("taskId", ctx.GetTaskId())

	return Data{InputKey: InputKey{
		Headers:      Headers{},
		ReasonPhrase: "reason reason reason",
		Body:         Body{},
		StatusCode:   data.InputKey.StatusCode + 2,
	}}, nil
}

func main() {

	apiClient := client.NewAPIClient(
		authSettings(),
		httpSettings(),
	)
	taskRunner := worker.NewTaskRunnerWithApiClient(apiClient)
	//taskRunner.StartWorker("hello", ProcessBody2, 1, 100)
	//taskRunner.StartWorkerWithDomain("hello", ProcessBody2, 1, 100, "d1")
	worker2 := worker.NewWorkerWithCtx("hello", ProcessBody3)
	worker2.Options = &worker.TaskWorkerOptions{
		BatchSize:    10,
		PollInterval: 100,
	}

	//worker.Options.Domain = "d1"
	err := worker2.Start(taskRunner)
	if err != nil {
		return
	}
	taskRunner.WaitWorkers()
	////worker := worker.NewWorker("hello", ProcessBody)
	//
	////taskClient := client.NewTaskClient(apiClient)
	////workerFn, _ := worker.GetWorker(ProcessBody)
	////_ = taskRunner.StartWorker2("hello", workerFn, 1, 1)
	//
	//opts := &client.TaskResourceApiBatchPollOpts{}
	//tasks, _, err := taskClient.BatchPollTask(context.Background(), "hello", opts)
	////task := model.PolledTask{
	////	TaskType:   "abcd",
	////	InputData:  "{}",
	////	OutputData: "{}",
	////}
	//if err != nil {
	//	fmt.Println("Got error ", err.Error())
	//}
	//fmt.Println("Got", len(tasks))

	//for i := range tasks {
	//	fmt.Println("Task.Id", tasks[i].TaskId)
	//	res, err := te.Execute(&tasks[i])
	//	if err != nil {
	//		fmt.Println("Error executing process body ", err.Error())
	//	} else {
	//		fmt.Println("Response of process body", res)
	//	}
	//}

	//data, _ := json.Marshal(task.InputData)
	//jsons := string(data)
	//fmt.Println(jsons)
}
