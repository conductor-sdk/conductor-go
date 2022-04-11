package http_model

type PollData struct {
	QueueName    string `json:"queueName,omitempty"`
	Domain       string `json:"domain,omitempty"`
	WorkerId     string `json:"workerId,omitempty"`
	LastPollTime int64  `json:"lastPollTime,omitempty"`
}
