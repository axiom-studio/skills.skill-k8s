package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/axiom-studio/skills.sdk/executor"
	"github.com/axiom-studio/skills.sdk/grpc"
	"github.com/axiom-studio/skills.sdk/resolver"
)

func main() {
	// Get port from env or use default
	port := os.Getenv("SKILL_PORT")
	if port == "" {
		port = "50051"
	}

	// Create skill server
	server := grpc.NewSkillServer("skill-k8s", "1.0.0")

	// Register K8s executors with schemas
	server.RegisterExecutorWithSchema("k8s-get", &K8sGetExecutor{}, makeNodeSchema("k8s-get", "Get a Kubernetes resource", K8sGetSchema))
	server.RegisterExecutorWithSchema("k8s-list", &K8sListExecutor{}, makeNodeSchema("k8s-list", "List Kubernetes resources", K8sListSchema))
	server.RegisterExecutorWithSchema("k8s-logs", &K8sLogsExecutor{}, makeNodeSchema("k8s-logs", "Get logs from a pod", K8sLogsSchema))
	server.RegisterExecutorWithSchema("k8s-events", &K8sEventsExecutor{}, makeNodeSchema("k8s-events", "Get events for a resource", K8sEventsSchema))
	server.RegisterExecutorWithSchema("k8s-restart", &K8sRestartExecutor{}, makeNodeSchema("k8s-restart", "Restart a deployment", K8sRestartSchema))
	server.RegisterExecutorWithSchema("k8s-scale", &K8sScaleExecutor{}, makeNodeSchema("k8s-scale", "Scale a deployment", K8sScaleSchema))
	server.RegisterExecutorWithSchema("k8s-patch", &K8sPatchExecutor{}, makeNodeSchema("k8s-patch", "Patch a resource", K8sPatchSchema))
	server.RegisterExecutorWithSchema("k8s-delete", &K8sDeleteExecutor{}, makeNodeSchema("k8s-delete", "Delete a resource", K8sDeleteSchema))

	fmt.Printf("Starting skill-k8s gRPC server on port %s\n", port)
	if err := server.Serve(port); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to serve: %v\n", err)
		os.Exit(1)
	}
}

func makeNodeSchema(name, description string, inputSchema map[string]interface{}) *resolver.NodeSchema {
	return &resolver.NodeSchema{
		Name:        name,
		DisplayName: name,
		Category:    "action",
		Description: description,
		Sections: []*resolver.ConfigSection{
			{
				Title: "Configuration",
				Fields: []*resolver.FieldSchema{
					{Key: "cluster", Type: resolver.FieldTypeText, Label: "Cluster"},
					{Key: "namespace", Type: resolver.FieldTypeText, Label: "Namespace"},
					{Key: "kind", Type: resolver.FieldTypeText, Label: "Kind"},
					{Key: "name", Type: resolver.FieldTypeText, Label: "Name"},
				},
			},
		},
	}
}

func mustMarshal(v interface{}) []byte {
	data, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	return data
}

// K8sGetExecutor handles k8s-get node type
type K8sGetExecutor struct{}

func (e *K8sGetExecutor) Type() string { return "k8s-get" }

func (e *K8sGetExecutor) Execute(ctx context.Context, step *executor.StepDefinition, resolver executor.TemplateResolver) (*executor.StepResult, error) {
	// TODO: Implement actual K8s get logic
	// For now, return a placeholder
	return &executor.StepResult{
		Output: map[string]interface{}{
			"message": "k8s-get executed",
			"config":  step.Config,
		},
	}, nil
}

// K8sListExecutor handles k8s-list node type
type K8sListExecutor struct{}

func (e *K8sListExecutor) Type() string { return "k8s-list" }

func (e *K8sListExecutor) Execute(ctx context.Context, step *executor.StepDefinition, resolver executor.TemplateResolver) (*executor.StepResult, error) {
	return &executor.StepResult{
		Output: map[string]interface{}{
			"message": "k8s-list executed",
			"config":  step.Config,
		},
	}, nil
}

// K8sLogsExecutor handles k8s-logs node type
type K8sLogsExecutor struct{}

func (e *K8sLogsExecutor) Type() string { return "k8s-logs" }

func (e *K8sLogsExecutor) Execute(ctx context.Context, step *executor.StepDefinition, resolver executor.TemplateResolver) (*executor.StepResult, error) {
	return &executor.StepResult{
		Output: map[string]interface{}{
			"message": "k8s-logs executed",
			"config":  step.Config,
		},
	}, nil
}

// K8sEventsExecutor handles k8s-events node type
type K8sEventsExecutor struct{}

func (e *K8sEventsExecutor) Type() string { return "k8s-events" }

func (e *K8sEventsExecutor) Execute(ctx context.Context, step *executor.StepDefinition, resolver executor.TemplateResolver) (*executor.StepResult, error) {
	return &executor.StepResult{
		Output: map[string]interface{}{
			"message": "k8s-events executed",
			"config":  step.Config,
		},
	}, nil
}

// K8sRestartExecutor handles k8s-restart node type
type K8sRestartExecutor struct{}

func (e *K8sRestartExecutor) Type() string { return "k8s-restart" }

func (e *K8sRestartExecutor) Execute(ctx context.Context, step *executor.StepDefinition, resolver executor.TemplateResolver) (*executor.StepResult, error) {
	return &executor.StepResult{
		Output: map[string]interface{}{
			"message": "k8s-restart executed",
			"config":  step.Config,
		},
	}, nil
}

// K8sScaleExecutor handles k8s-scale node type
type K8sScaleExecutor struct{}

func (e *K8sScaleExecutor) Type() string { return "k8s-scale" }

func (e *K8sScaleExecutor) Execute(ctx context.Context, step *executor.StepDefinition, resolver executor.TemplateResolver) (*executor.StepResult, error) {
	return &executor.StepResult{
		Output: map[string]interface{}{
			"message": "k8s-scale executed",
			"config":  step.Config,
		},
	}, nil
}

// K8sPatchExecutor handles k8s-patch node type
type K8sPatchExecutor struct{}

func (e *K8sPatchExecutor) Type() string { return "k8s-patch" }

func (e *K8sPatchExecutor) Execute(ctx context.Context, step *executor.StepDefinition, resolver executor.TemplateResolver) (*executor.StepResult, error) {
	return &executor.StepResult{
		Output: map[string]interface{}{
			"message": "k8s-patch executed",
			"config":  step.Config,
		},
	}, nil
}

// K8sDeleteExecutor handles k8s-delete node type
type K8sDeleteExecutor struct{}

func (e *K8sDeleteExecutor) Type() string { return "k8s-delete" }

func (e *K8sDeleteExecutor) Execute(ctx context.Context, step *executor.StepDefinition, resolver executor.TemplateResolver) (*executor.StepResult, error) {
	return &executor.StepResult{
		Output: map[string]interface{}{
			"message": "k8s-delete executed",
			"config":  step.Config,
		},
	}, nil
}

// Node schemas (simplified for now)
var K8sGetSchema = map[string]interface{}{
	"type": "object",
	"properties": map[string]interface{}{
		"cluster": map[string]interface{}{
			"type":        "string",
			"description": "Target cluster name",
		},
		"namespace": map[string]interface{}{
			"type":        "string",
			"description": "Target namespace",
		},
		"kind": map[string]interface{}{
			"type":        "string",
			"description": "Resource kind (e.g., Pod, Deployment)",
		},
		"name": map[string]interface{}{
			"type":        "string",
			"description": "Resource name",
		},
	},
	"required": []string{"kind", "name"},
}

var K8sListSchema = map[string]interface{}{
	"type": "object",
	"properties": map[string]interface{}{
		"cluster":   map[string]interface{}{"type": "string"},
		"namespace": map[string]interface{}{"type": "string"},
		"kind":      map[string]interface{}{"type": "string"},
		"labelSelector": map[string]interface{}{
			"type":        "string",
			"description": "Label selector to filter resources",
		},
	},
	"required": []string{"kind"},
}

var K8sLogsSchema = map[string]interface{}{
	"type": "object",
	"properties": map[string]interface{}{
		"cluster":   map[string]interface{}{"type": "string"},
		"namespace": map[string]interface{}{"type": "string"},
		"pod":       map[string]interface{}{"type": "string"},
		"container": map[string]interface{}{"type": "string"},
		"tailLines": map[string]interface{}{"type": "integer"},
		"follow":    map[string]interface{}{"type": "boolean"},
	},
	"required": []string{"pod"},
}

var K8sEventsSchema = map[string]interface{}{
	"type": "object",
	"properties": map[string]interface{}{
		"cluster":   map[string]interface{}{"type": "string"},
		"namespace": map[string]interface{}{"type": "string"},
		"kind":      map[string]interface{}{"type": "string"},
		"name":      map[string]interface{}{"type": "string"},
	},
}

var K8sRestartSchema = map[string]interface{}{
	"type": "object",
	"properties": map[string]interface{}{
		"cluster":   map[string]interface{}{"type": "string"},
		"namespace": map[string]interface{}{"type": "string"},
		"kind":      map[string]interface{}{"type": "string"},
		"name":      map[string]interface{}{"type": "string"},
	},
	"required": []string{"kind", "name"},
}

var K8sScaleSchema = map[string]interface{}{
	"type": "object",
	"properties": map[string]interface{}{
		"cluster":   map[string]interface{}{"type": "string"},
		"namespace": map[string]interface{}{"type": "string"},
		"kind":      map[string]interface{}{"type": "string"},
		"name":      map[string]interface{}{"type": "string"},
		"replicas":  map[string]interface{}{"type": "integer"},
	},
	"required": []string{"kind", "name", "replicas"},
}

var K8sPatchSchema = map[string]interface{}{
	"type": "object",
	"properties": map[string]interface{}{
		"cluster":   map[string]interface{}{"type": "string"},
		"namespace": map[string]interface{}{"type": "string"},
		"kind":      map[string]interface{}{"type": "string"},
		"name":      map[string]interface{}{"type": "string"},
		"patch":     map[string]interface{}{"type": "object"},
	},
	"required": []string{"kind", "name", "patch"},
}

var K8sDeleteSchema = map[string]interface{}{
	"type": "object",
	"properties": map[string]interface{}{
		"cluster":   map[string]interface{}{"type": "string"},
		"namespace": map[string]interface{}{"type": "string"},
		"kind":      map[string]interface{}{"type": "string"},
		"name":      map[string]interface{}{"type": "string"},
	},
	"required": []string{"kind", "name"},
}