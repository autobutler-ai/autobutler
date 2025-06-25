package llm

import (
	"autobutler/pkg/util"
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/openai/openai-go"
)

var (
	mcpRegistry = &McpRegistry{
		Functions: make(map[string]McpFunction),
	}
)

func (r McpRegistry) Add(param0 float64, param1 float64) float64 {
	return param0 + param1
}

func init() {
	addFn, err := NewMcpFunction(mcpRegistry.Add, "Adds two numbers together and returns the result.")
	if err != nil {
		panic(fmt.Sprintf("failed to generate JSON schema for add function: %v", err))
	}
	mcpRegistry.Functions[addFn.Name()] = *addFn
}

type McpRegistry struct {
	Functions map[string]McpFunction
}

type McpFunction struct {
	Definition openai.FunctionDefinitionParam
	fn 	   interface{}
	OutputHandler func(result any) (string, error)
}

func (f McpFunction) Name() string {
	return f.Definition.Name
}

func (f McpFunction) Parameters() openai.FunctionParameters {
	return f.Definition.Parameters
}

func (f McpFunction) Description() string {
	return f.Definition.Description.String()
}

func NewMcpFunction(fn interface{}, description string) (*McpFunction, error) {
	t := reflect.TypeOf(fn)
	if t.Kind() != reflect.Func {
		return nil, fmt.Errorf("expected a function, got %s", t.Kind())
	}

	params := map[string]any{}
	required := []string{}
	for i := 0; i < t.NumIn(); i++ {
		paramType := t.In(i)
		paramName := fmt.Sprintf("param%d", i)
		params[paramName] = util.TypeToJsonschema(paramType)
		required = append(required, paramName)
	}
	schema := map[string]any{
		"type":       "object",
		"properties": params,
		"required":   required,
	}

	return &McpFunction{
		Definition: openai.FunctionDefinitionParam{
			Name:        util.GetFunctionName(fn),
			Strict:      openai.Bool(false),
			Description: openai.String(description),
			Parameters:  schema,
		},
		fn: fn,
	}, nil
}

func (r McpRegistry) MakeToolCall(completion *openai.ChatCompletion) ([]any, error) {
	toolCalls := completion.Choices[0].Message.ToolCalls
	results := make([]any, 0, len(toolCalls))
	if len(toolCalls) > 0 {
		for _, toolCall := range toolCalls {
			var args map[string]float64
			if err := json.Unmarshal([]byte(toolCall.Function.Arguments), &args); err != nil {
				return nil, fmt.Errorf("failed to unmarshal function arguments: %w", err)
			}
			if _, ok := r.Functions[toolCall.Function.Name]; !ok {
				return nil, fmt.Errorf("function %s not found in registry", toolCall.Function.Name)
			}
			fnDef, ok := r.Functions[toolCall.Function.Name]
			if !ok {
				return nil, fmt.Errorf("function %s not found in registry", toolCall.Function.Name)
			}

			// Prepare argument list in order as defined in fnDef.Parameters
			var paramNames []string
			if fnDef.Parameters() != nil {
				if props, ok := fnDef.Parameters()["properties"].(map[string]interface{}); ok {
					for name := range props {
						paramNames = append(paramNames, name)
					}
				}
			}

			var argValues []interface{}
			for _, name := range paramNames {
				val, exists := args[name]
				if !exists {
					return nil, fmt.Errorf("missing argument '%s' for function %s", name, toolCall.Function.Name)
				}
				argValues = append(argValues, val)
			}

			var returnValue any
			var err error
			if returnValue, err = r.callByName(toolCall.Function.Name, argValues...); err != nil {
				return nil, fmt.Errorf("failed to call function %s: %w", toolCall.Function.Name, err)
			}
			results = append(results, returnValue)
		}
	}
	return results, nil
}

func (r McpRegistry) callByName(fnName string, args ...any) (any, error) {
	fn := reflect.ValueOf(&r).MethodByName(fnName)
	if fn.Kind() != reflect.Func {
		return nil, fmt.Errorf("function %s not found", fnName)
	}

	if fn.Type().NumIn() != len(args) {
		return nil, fmt.Errorf("function %s expects %d arguments, got %d", fnName, fn.Type().NumIn(), len(args))
	}

	in := make([]reflect.Value, len(args))
	for i, arg := range args {
		in[i] = reflect.ValueOf(arg)
	}

	out := fn.Call(in)
	if len(out) == 0 {
		return nil, nil // No return value
	}
	return out[0].Interface(), nil
}

func (r McpRegistry) toCompletionToolParam() []openai.ChatCompletionToolParam{
	var tools []openai.ChatCompletionToolParam
	for _, fn := range r.Functions {
		tools = append(tools, openai.ChatCompletionToolParam{
			Type:        "function",
			Function:    fn.Definition,
		})
	}
	return tools
}
