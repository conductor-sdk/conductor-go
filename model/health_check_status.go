package model

type HealthCheckStatus struct {
	HealthResults           []Health `json:"healthResults,omitempty"`
	SuppressedHealthResults []Health `json:"suppressedHealthResults,omitempty"`
	Healthy                 bool     `json:"healthy,omitempty"`
}
