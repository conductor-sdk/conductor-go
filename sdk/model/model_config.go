package model

type Config struct {
	CircuitBreakerConfig *OrkesCircuitBreakerConfig `json:"circuitBreakerConfig,omitempty"`
}
