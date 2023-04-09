package model

type SearchResultWorkflowSchedule struct {
	Results   []WorkflowScheduleExecution `json:"results,omitempty"`
	TotalHits int64                       `json:"totalHits,omitempty"`
}
