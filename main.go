package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/SWorM/v2/tool"

	"github.com/getkin/kin-openapi/openapi3"
)

const (
	// ### VERBOSITY level (in binary/decimal)
	// # - 0001/1: print out workflow
	// # - 0010/2: print out action
	// # etc ...
	// # to print multiple things just add the decimal

	VERBOSITY = 3
	DEBUG     = true
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func RunAction(action tool.Action, doc *openapi3.T) {
	docExt := (&tool.T{T: doc}).Extend()

	openapi_refrence := docExt.OperationID.Value(action.Name)
	if openapi_refrence != nil {
		// WIP
	}
	// WIP

	fmt.Printf("running: '%s'  \n", action.Name)
}

func main() {
	openapi_raw, err := os.ReadFile("./openapi.yaml")
	check(err)

	workflow_raw, err := os.ReadFile("./workflow.json")
	check(err)

	var workflow tool.Workflow
	err = json.Unmarshal(workflow_raw, &workflow)
	check(err)

	loader := openapi3.NewLoader()
	doc, err := loader.LoadFromData(openapi_raw)
	check(err)

	// for debugging workflow
	if DEBUG {
		tool.DumpWorkflow(workflow, VERBOSITY)
	}

	action_list := workflow.Actions
	for _, action := range action_list {
		action.Endpoint = workflow.Baseurl + action.Path

		// for debugging action
		if DEBUG {
			tool.DumpAction(action, VERBOSITY)
		}

		fmt.Printf("starting: '%s'  \n", action.Name)
		RunAction(action, doc)
		fmt.Printf("returning: '%s'  \n", action.Name)
	}
}
