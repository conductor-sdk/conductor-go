package model

// ChatCompletion extends LLMWorkerInput and represents a chat completion request
type ChatCompletion struct {
	LLMWorkerInput
	Messages     []ChatMessage `json:"messages,omitempty"`
	Instructions string        `json:"instructions,omitempty"`
	JSONOutput   bool          `json:"jsonOutput,omitempty"`
}