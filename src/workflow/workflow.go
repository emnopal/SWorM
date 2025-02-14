package workflow

import (
	"encoding/json"
	"fmt"
)

type Workflows struct {
	Workflow map[string]*Workflow
}

type Workflow struct {
	Name        string            `json:"name"`
	Description string            `json:"description"`
	Envs        map[string]string `json:"envs"`
	Actions     []WorkflowAction  `json:"actions"`
}

type WorkflowAction struct {
	ActionName string            `json:"action_name"`
	LocalEnv   map[string]string `json:"local_env"`
}

func NewWorkflows() *Workflows {
	return &Workflows{Workflow: make(map[string]*Workflow)}
}

func NewWorkflow() *Workflow {
	return &Workflow{}
}

func (w *Workflow) Dump() (string, error) {
	workflow_json, err := json.MarshalIndent(w, "", "  ")
	if err != nil {
		return "", fmt.Errorf("error dumping workflow: %w", err)
	}
	return string(workflow_json), nil
}

func (w *Workflow) LoadWorkflow(workflow_string string) error {
	err := json.Unmarshal([]byte(workflow_string), w)
	if err != nil {
		return fmt.Errorf("error reading workflow file: %w", err)
	}
	return nil
}

func (ws *Workflows) Dump() (string, error) {
	workflow_json, err := json.MarshalIndent(ws, "", "  ")
	if err != nil {
		return "", fmt.Errorf("error dumping workflows: %w", err)
	}
	return string(workflow_json), nil
}

func (ws *Workflows) LoadWorkflows(workflow_list []string) error {
	for _, workflow_string := range workflow_list {
		workflow := NewWorkflow()
		err := workflow.LoadWorkflow(workflow_string)
		if err != nil {
			return fmt.Errorf("error loading workflow: %w", err)
		}
		ws.Workflow[workflow.Name] = workflow
	}
	return nil
}
