package http_model

type SearchResultWorkflowSummary struct {
	TotalHits int64             `json:"totalHits,omitempty"`
	Results   []WorkflowSummary `json:"results,omitempty"`
}
