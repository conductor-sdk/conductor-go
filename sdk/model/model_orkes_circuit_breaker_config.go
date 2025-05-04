package model

type OrkesCircuitBreakerConfig struct {
	AutomaticTransitionFromOpenToHalfOpenEnabled bool    `json:"automaticTransitionFromOpenToHalfOpenEnabled,omitempty"`
	FailureRateThreshold                         float32 `json:"failureRateThreshold,omitempty"`
	MaxWaitDurationInHalfOpenState               int64   `json:"maxWaitDurationInHalfOpenState,omitempty"`
	MinimumNumberOfCalls                         int32   `json:"minimumNumberOfCalls,omitempty"`
	PermittedNumberOfCallsInHalfOpenState        int32   `json:"permittedNumberOfCallsInHalfOpenState,omitempty"`
	SlidingWindowSize                            int32   `json:"slidingWindowSize,omitempty"`
	SlowCallDurationThreshold                    int64   `json:"slowCallDurationThreshold,omitempty"`
	SlowCallRateThreshold                        float32 `json:"slowCallRateThreshold,omitempty"`
	WaitDurationInOpenState                      int64   `json:"waitDurationInOpenState,omitempty"`
}
