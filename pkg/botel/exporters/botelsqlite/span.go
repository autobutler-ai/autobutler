package botelsqlite

import (
	"fmt"

	"go.opentelemetry.io/otel/codes"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

// SpanToJSON converts a ReadOnlySpan to the JSON structure shown in types.go
func SpanToJSON(span sdktrace.ReadOnlySpan) (map[string]interface{}, error) {
	spanCtx := span.SpanContext()
	parentCtx := span.Parent()

	result := map[string]interface{}{
		"Name": span.Name(),
		"SpanContext": map[string]interface{}{
			"TraceID":    spanCtx.TraceID().String(),
			"SpanID":     spanCtx.SpanID().String(),
			"TraceFlags": fmt.Sprintf("%02x", spanCtx.TraceFlags()),
			"TraceState": spanCtx.TraceState().String(),
			"Remote":     spanCtx.IsRemote(),
		},
		"Parent": map[string]interface{}{
			"TraceID":    parentCtx.TraceID().String(),
			"SpanID":     parentCtx.SpanID().String(),
			"TraceFlags": fmt.Sprintf("%02x", parentCtx.TraceFlags()),
			"TraceState": parentCtx.TraceState().String(),
			"Remote":     parentCtx.IsRemote(),
		},
		"SpanKind":          span.SpanKind(),
		"StartTime":         span.StartTime(),
		"EndTime":           span.EndTime(),
		"DroppedAttributes": span.DroppedAttributes(),
		"DroppedEvents":     span.DroppedEvents(),
		"DroppedLinks":      span.DroppedLinks(),
		"ChildSpanCount":    span.ChildSpanCount(),
	}

	// Convert attributes
	attrs := []map[string]interface{}{}
	for _, attr := range span.Attributes() {
		attrs = append(attrs, map[string]interface{}{
			"Key": string(attr.Key),
			"Value": map[string]interface{}{
				"Type":  attr.Value.Type().String(),
				"Value": attr.Value.AsInterface(),
			},
		})
	}
	result["Attributes"] = attrs

	// Convert events
	events := []map[string]interface{}{}
	for _, event := range span.Events() {
		eventAttrs := []map[string]interface{}{}
		for _, attr := range event.Attributes {
			eventAttrs = append(eventAttrs, map[string]interface{}{
				"Key": string(attr.Key),
				"Value": map[string]interface{}{
					"Type":  attr.Value.Type().String(),
					"Value": attr.Value.AsInterface(),
				},
			})
		}
		events = append(events, map[string]interface{}{
			"Name":       event.Name,
			"Time":       event.Time,
			"Attributes": eventAttrs,
		})
	}
	if len(events) == 0 {
		result["Events"] = nil
	} else {
		result["Events"] = events
	}

	// Convert links
	links := []map[string]interface{}{}
	for _, link := range span.Links() {
		linkAttrs := []map[string]interface{}{}
		for _, attr := range link.Attributes {
			linkAttrs = append(linkAttrs, map[string]interface{}{
				"Key": string(attr.Key),
				"Value": map[string]interface{}{
					"Type":  attr.Value.Type().String(),
					"Value": attr.Value.AsInterface(),
				},
			})
		}
		links = append(links, map[string]interface{}{
			"SpanContext": map[string]interface{}{
				"TraceID":    link.SpanContext.TraceID().String(),
				"SpanID":     link.SpanContext.SpanID().String(),
				"TraceFlags": fmt.Sprintf("%02x", link.SpanContext.TraceFlags()),
				"TraceState": link.SpanContext.TraceState().String(),
				"Remote":     link.SpanContext.IsRemote(),
			},
			"Attributes": linkAttrs,
		})
	}
	if len(links) == 0 {
		result["Links"] = nil
	} else {
		result["Links"] = links
	}

	// Status
	status := span.Status()
	result["Status"] = map[string]interface{}{
		"Code":        status.Code.String(),
		"Description": status.Description,
	}

	// Check status code explicitly
	if status.Code == codes.Unset {
		result["Status"].(map[string]interface{})["Code"] = "Unset"
	} else if status.Code == codes.Ok {
		result["Status"].(map[string]interface{})["Code"] = "Ok"
	} else if status.Code == codes.Error {
		result["Status"].(map[string]interface{})["Code"] = "Error"
	}

	// Resource attributes
	resourceAttrs := []map[string]interface{}{}
	for _, attr := range span.Resource().Attributes() {
		resourceAttrs = append(resourceAttrs, map[string]interface{}{
			"Key": string(attr.Key),
			"Value": map[string]interface{}{
				"Type":  attr.Value.Type().String(),
				"Value": attr.Value.AsInterface(),
			},
		})
	}
	result["Resource"] = resourceAttrs

	// Instrumentation scope
	scope := span.InstrumentationScope()
	scopeAttrs := []map[string]interface{}{}
	iter := scope.Attributes.Iter()
	for iter.Next() {
		attr := iter.Attribute()
		scopeAttrs = append(scopeAttrs, map[string]interface{}{
			"Key": string(attr.Key),
			"Value": map[string]interface{}{
				"Type":  attr.Value.Type().String(),
				"Value": attr.Value.AsInterface(),
			},
		})
	}
	if len(scopeAttrs) == 0 {
		scopeAttrs = nil
	}

	result["InstrumentationScope"] = map[string]interface{}{
		"Name":       scope.Name,
		"Version":    scope.Version,
		"SchemaURL":  scope.SchemaURL,
		"Attributes": scopeAttrs,
	}

	// For backward compatibility
	result["InstrumentationLibrary"] = result["InstrumentationScope"]

	return result, nil
}
