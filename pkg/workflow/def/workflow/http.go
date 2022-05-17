package workflow

type HttpTask struct {
	Task
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

func Http(taskRefName string, input *HttpInput) *HttpTask {
	if input.Method == "" {
		input.Method = GET
	}
	return &HttpTask{
		Task{
			name:              taskRefName,
			taskReferenceName: taskRefName,
			description:       "",
			taskType:          HTTP,
			optional:          false,
			inputParameters: map[string]interface{}{
				"http_request": input,
			},
		},
	}
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
