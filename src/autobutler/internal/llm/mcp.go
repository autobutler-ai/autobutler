package llm

import (
	"autobutler/pkg/util"
	"fmt"
	"reflect"
	"strings"

	"github.com/openai/openai-go"
)

var (
	mcpRegistry = &McpRegistry{
		Functions: make(map[string]openai.FunctionDefinitionParam),
	}
)

func (r McpRegistry) Add(param0 float64, param1 float64) float64 {
	return param0 + param1
}

func init() {
	addFn, err := util.GenerateJSONSchema(mcpRegistry.Add, "Adds two numbers together and returns the result.")
	if err != nil {
		panic(fmt.Sprintf("failed to generate JSON schema for add function: %v", err))
	}
	mcpRegistry.Functions[addFn.Name] = *addFn
}

func (r McpRegistry) callByName(fnName string, args ...any) (any, error) {
	fn := reflect.ValueOf(&r).MethodByName(strings.TrimSuffix(fnName,"-fm"))
	kind := fn.Kind()
	fmt.Sprintf("Function %s kind: %s", fnName, kind)
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
