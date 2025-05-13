package model

import (
	"encoding/json"
)

// ChatCompletion extends LLMWorkerInput and represents a chat completion request
type ChatCompletion struct {
	LLMWorkerInput
	Messages     []ChatMessage `json:"messages,omitempty"`
	Instructions string        `json:"instructions,omitempty"`
	JSONOutput   bool          `json:"jsonOutput,omitempty"`
}

// GetMessages returns the chat messages
func (c *ChatCompletion) GetMessages() []ChatMessage {
	return c.Messages
}

// SetMessages sets the chat messages
func (c *ChatCompletion) SetMessages(messages []ChatMessage) {
	c.Messages = messages
}

// GetInstructions returns the instructions for the chat
func (c *ChatCompletion) GetInstructions() string {
	return c.Instructions
}

// SetInstructions sets the instructions for the chat
func (c *ChatCompletion) SetInstructions(instructions string) {
	c.Instructions = instructions
}

// GetJSONOutput returns whether the output should be in JSON format
func (c *ChatCompletion) GetJSONOutput() bool {
	return c.JSONOutput
}

// SetJSONOutput sets whether the output should be in JSON format
func (c *ChatCompletion) SetJSONOutput(jsonOutput bool) {
	c.JSONOutput = jsonOutput
}

// MarshalJSON implements the json.Marshaler interface
func (c ChatCompletion) MarshalJSON() ([]byte, error) {
	type Alias ChatCompletion
	return json.Marshal(&struct {
		Alias
	}{
		Alias: Alias(c),
	})
}

// NewChatCompletion creates a new ChatCompletion instance
func NewChatCompletion() *ChatCompletion {
	return &ChatCompletion{}
}