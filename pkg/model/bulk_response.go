package model

type BulkResponse struct {
	BulkErrorResults      map[string]string `json:"bulkErrorResults,omitempty"`
	BulkSuccessfulResults []string          `json:"bulkSuccessfulResults,omitempty"`
}
