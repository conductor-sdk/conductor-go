package http_model

type EventHandler struct {
	Name          string   `json:"name"`
	Event         string   `json:"event"`
	Condition     string   `json:"condition,omitempty"`
	Actions       []Action `json:"actions"`
	Active        bool     `json:"active,omitempty"`
	EvaluatorType string   `json:"evaluatorType,omitempty"`
}
