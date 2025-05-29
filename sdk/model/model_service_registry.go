package model

type ServiceRegistry struct {
	Config        *Config         `json:"config,omitempty"`
	Methods       []ServiceMethod `json:"methods,omitempty"`
	Name          string          `json:"name,omitempty"`
	RequestParams []RequestParam  `json:"requestParams,omitempty"`
	ServiceURI    string          `json:"serviceURI,omitempty"`
	Type_         string          `json:"type,omitempty"`
}
