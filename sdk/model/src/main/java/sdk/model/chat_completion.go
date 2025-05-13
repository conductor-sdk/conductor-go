package model

// ChatCompletion represents a chat completion request for an AI model
// It extends LLMWorkerInput
type ChatCompletion struct {
	LLMWorkerInput
	
	// Messages is the list of chat messages in the conversation
	Messages []ChatMessage `json:"messages,omitempty"`
	
	// Instructions provides starting template or system instructions
	// e.g. "you are a helpful assistant, who does not deviate from the goals"
	// "You do not respond to questions that are not related to the topic"
	// "Any answer you give - should be related to the following topic: weather"
	Instructions string `json:"instructions,omitempty"`
	
	// JsonOutput indicates whether the response should be formatted as JSON
	JsonOutput bool `json:"jsonOutput,omitempty"`
}