package http_model

type SearchResultWorkflow struct {
	TotalHits int64      `json:"totalHits,omitempty"`
	Results   []Workflow `json:"results,omitempty"`
}
