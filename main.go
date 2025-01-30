package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/SWorM/v2/agent"
	"github.com/SWorM/v2/debug"
	"github.com/SWorM/v2/tool"

	"github.com/getkin/kin-openapi/openapi3"
)

func executeAction(workflow tool.Workflow, action tool.Action, doc *openapi3.T) (bool, *agent.Response, error) {
	agent := agent.New(workflow, action, &tool.T{T: doc})
	should_continue, response, err := agent.RunAction()
	return should_continue, response, err
}

func main() {
	debug.LoadDebugConfig()

	openapi_raw, err := os.ReadFile("./openapi.yaml")
	debug.Check(err)

	workflow_raw, err := os.ReadFile("./workflow.json")
	debug.Check(err)

	var workflow tool.Workflow
	err = json.Unmarshal(workflow_raw, &workflow)
	debug.Check(err)

	loader := openapi3.NewLoader()
	doc, err := loader.LoadFromData(openapi_raw)
	debug.Check(err)

	if debug.DEBUG {
		tool.DumpWorkflow(workflow)
	}

	action_list := workflow.Actions
	for _, action := range action_list {
		action.Endpoint = workflow.Baseurl + action.Path

		if debug.DEBUG {
			tool.DumpAction(action)
		}

		println("endpoint", action.Endpoint)

		fmt.Printf("starting: '%s'  \n", action.Name)

		//## run action from the workflow
		should_continue, response, err := executeAction(workflow, action, doc)
		if !debug.SKIPCHECKS && !should_continue {
			debug.Check(err)
		} else if err != nil {
			fmt.Println(err.Error())
		}

		//## show result of action

		if response != nil {
			fmt.Printf("returning: '%s'  \n", response.Status())
			// fmt.Printf("returning: '%s'  \n", response.Body())
		} else {
			fmt.Printf("returning: '%v'  \n", nil)
		}
	}

	//## TODO: save history as {date}.log

	//## exit
}
