package model

type SearchResultWorkflowScheduleExecutionModel struct {
	Results   []WorkflowScheduleExecutionModel `json:"results,omitempty"`
	TotalHits int64                            `json:"totalHits,omitempty"`
}
