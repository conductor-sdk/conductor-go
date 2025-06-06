package model

type CircuitBreakerTransitionResponse struct {
	CurrentState        string `json:"currentState,omitempty"`
	Message             string `json:"message,omitempty"`
	PreviousState       string `json:"previousState,omitempty"`
	Service             string `json:"service,omitempty"`
	TransitionTimestamp int64  `json:"transitionTimestamp,omitempty"`
}
