package model

type ProtoRegistryEntry struct {
	Data        string `json:"data,omitempty"`
	Filename    string `json:"filename,omitempty"`
	ServiceName string `json:"serviceName,omitempty"`
}
