package workflow

type httpTask struct {
	task
}

type HttpMethod string

const (
	GET     HttpMethod = "GET"
	PUT     HttpMethod = "PUT"
	POST    HttpMethod = "POST"
	DELETE  HttpMethod = "DELETE"
	HEAD    HttpMethod = "HEAD"
	OPTIONS HttpMethod = "OPTIONS"
)

func Http(taskRefName string, input *HttpInput) *httpTask {
	httpTask := &httpTask{task{
		name:              taskRefName,
		taskReferenceName: taskRefName,
		description:       "",
		taskType:          HTTP,
		optional:          false,
		inputParameters:   map[string]interface{}{},
	}}
	if input.Method == "" {
		input.Method = GET
	}
	httpTask.inputParameters["http_request"] = input
	return httpTask
}

type HttpInput struct {
	Method            HttpMethod          `json:"method"`
	Uri               string              `json:"uri"`
	Headers           map[string][]string `json:"headers,omitempty"`
	Accept            string              `json:"accept,omitempty"`
	ContentType       string              `json:"contentType,omitempty"`
	ConnectionTimeOut int16               `json:"ConnectionTimeOut,omitempty"`
	ReadTimeout       int16               `json:"readTimeout,omitempty"`
	Body              interface{}         `json:"body,omitempty"`
}

// Input to the task
func (task *httpTask) Input(key string, value *interface{}) *httpTask {
	task.inputParameters[key] = value
	return task
}
func (task *httpTask) Description(description string) *httpTask {
	task.task.Description(description)
	return task
}
func (task *httpTask) Optional(optional bool) *httpTask {
	task.task.Optional(optional)
	return task
}
