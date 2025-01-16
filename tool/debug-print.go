package tool

import (
	"fmt"
)

const (
	WORKFLOW_CODE int = 1
	ACTION_CODE   int = 2
)

func CheckVerbosity(verbosity int, code int) bool {
	return (verbosity&code == code)
}

func DumpWorkflow(workflow Workflow, verbosity int) {
	if !CheckVerbosity(verbosity, WORKFLOW_CODE) {
		return
	}

	fmt.Println("#=========================================##=========================================#")
	fmt.Printf("| wf-name:        %-66s |\n", workflow.Name)
	fmt.Printf("| wf-description: %-66s |\n", workflow.Description)
	fmt.Printf("| wf-baseurl:     %-66s |\n", workflow.Baseurl)
	fmt.Printf("| wf-envs:        %-66s |\n", "")
	if len(workflow.Envs) > 0 {
		for key, value := range workflow.Envs {

			fmt.Printf("| - %-80s |\n", fmt.Sprintf("%s: %s", key, value))
		}
	} else {
		fmt.Printf("| - %-80s |\n", "<EMPTY>")
	}
	fmt.Println("#=========================================##=========================================#")
}

func DumpAction(action Action, verbosity int) {
	if !CheckVerbosity(verbosity, ACTION_CODE) {
		return
	}

	fmt.Println(" -  name: " + action.Name)
	fmt.Println(" -  description: " + action.Description)
	fmt.Println(" -  path: " + action.Path)
	fmt.Println(" -  method: " + action.Method)
	fmt.Println(" -  endpoint: " + action.Endpoint)

	fmt.Println("[-] header: ")
	if len(action.Headers) > 0 {
		for key, value := range action.Headers {
			fmt.Printf("    -  %s: %s \n", key, value)
		}
	} else {
		fmt.Println("    -  <EMPTY>")
	}

	fmt.Println("[-] parameters: ")
	if len(action.Parameters) > 0 {
		for key, value := range action.Parameters {
			fmt.Printf("    -  %s: %s \n", key, value)
		}
	} else {
		fmt.Println("    -  <EMPTY>")
	}
}
