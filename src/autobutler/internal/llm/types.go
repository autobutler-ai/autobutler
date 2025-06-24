package llm

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/openai/openai-go"
)

type ChatRole string

const (
	ChatRoleUser   ChatRole = "user"
	ChatRoleSystem ChatRole = "system"
	ChatRoleDummy  ChatRole = "dummy"
	ChatRoleError  ChatRole = "error"
)

type ChatMessage struct {
	Role      ChatRole `json:"role"`
	Content   string   `json:"content"`
	Timestamp string   `json:"timestamp"` // ISO 8601 format
}

func ErrorChatMessage(err error) ChatMessage {
	return ChatMessage{
		Role:      ChatRoleError,
		Content:   fmt.Sprintf("An error occurred while processing your request: %v", err),
		Timestamp: GetTimestamp(time.Now()),
	}
}

func GetTimestamp(timestamp time.Time) string {
	// Matches JS new Date().toLocaleTimeString()
	return timestamp.Format("3:04:05 PM")
}

func FromCompletionToChatMessage(completion openai.ChatCompletion) ChatMessage {
	if len(completion.Choices) == 0 {
		return ChatMessage{
			Role:      ChatRoleError,
			Content:   "",
			Timestamp: GetTimestamp(time.Now()),
		}
	}
	return ChatMessage{
		Role:    ChatRoleSystem,
		Content: completion.Choices[0].Message.Content,
		// Matches JS new Date().toLocaleTimeString()
		Timestamp: GetTimestamp(time.Now()),
	}
}

type McpRegistry struct {
	Functions map[string]openai.FunctionDefinitionParam
}

func (r McpRegistry) toCompletionToolParam() []openai.ChatCompletionToolParam{
	var tools []openai.ChatCompletionToolParam
	for _, fn := range r.Functions {
		tools = append(tools, openai.ChatCompletionToolParam{
			Type:        "function",
			Function:    fn,
		})
	}
	return tools
}

func (r McpRegistry) MakeToolCall(completion *openai.ChatCompletion) error {
	toolCalls := completion.Choices[0].Message.ToolCalls
	if len(toolCalls) > 0 {
		for _, toolCall := range toolCalls {
			var args map[string]float64
			if err := json.Unmarshal([]byte(toolCall.Function.Arguments), &args); err != nil {
				return fmt.Errorf("failed to unmarshal function arguments: %w", err)
			}
			if _, ok := r.Functions[toolCall.Function.Name]; !ok {
				return fmt.Errorf("function %s not found in registry", toolCall.Function.Name)
			}
			switch toolCall.Function.Name {
			case "add":
				param0, ok1 := args["param0"]
				param1, ok2 := args["param1"]
				if !ok1 || !ok2 {
					return fmt.Errorf("invalid arguments for add function: expected 'param1' and 'param2'")
				}
				result := add(param0, param1)
				completion.Choices[0].Message.Content = fmt.Sprintf("The result of adding %f and %f is %f", param0, param1, result)
			}
		}
	}
	return nil
}
