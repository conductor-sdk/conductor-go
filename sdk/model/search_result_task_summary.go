package model

type SearchResultTaskSummary struct {
	TotalHits int64         `json:"totalHits,omitempty"`
	Results   []TaskSummary `json:"results,omitempty"`
}
