package model

// WorkflowConsistency defines the consistency level for workflow execution
type WorkflowConsistency string

const (
	// SynchronousConsistency - Request is kept in memory until the evaluation completes and flushed to the persistence afterwards
	// If the node dies before the writes are flushed, the workflow request is gone
	SynchronousConsistency WorkflowConsistency = "SYNCHRONOUS"

	// DurableConsistency - Default
	// Request is stored in a persistence store before the evaluation - Guarantees the execution
	// Implements durable workflows -- Suitable for most use cases
	DurableConsistency WorkflowConsistency = "DURABLE"

	// RegionDurableConsistency - In the multi-region setup, guarantees that the start request is replicated across the region
	// Safest
	// Slowest
	RegionDurableConsistency WorkflowConsistency = "REGION_DURABLE"
)

// String returns the string representation of WorkflowConsistency
func (wc WorkflowConsistency) String() string {
	return string(wc)
}

// IsValid checks if the WorkflowConsistency value is valid
func (wc WorkflowConsistency) IsValid() bool {
	switch wc {
	case SynchronousConsistency, DurableConsistency, RegionDurableConsistency:
		return true
	default:
		return false
	}
}

// GetDefault returns the default WorkflowConsistency
func GetDefaultWorkflowConsistency() WorkflowConsistency {
	return DurableConsistency
}
