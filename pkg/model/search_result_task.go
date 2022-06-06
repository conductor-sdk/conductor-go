package model

type SearchResultTask struct {
	TotalHits int64  `json:"totalHits,omitempty"`
	Results   []Task `json:"results,omitempty"`
}
