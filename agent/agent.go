package agent

import (
	"errors"
	"fmt"

	"github.com/SWorM/v2/debug"
	"github.com/SWorM/v2/tool"

	"github.com/go-resty/resty/v2"
)

type Agent struct {
	Workflow tool.Workflow
	Action   tool.Action
	DocExt   tool.ExtendedT

	PathParam  map[string]string
	QueryParam map[string]string
}

func New(workflow tool.Workflow, action tool.Action, doc *tool.T) *Agent {
	agent := &Agent{
		Workflow: workflow,
		Action:   action,
		DocExt:   *doc.Extend(),

		PathParam:  make(map[string]string),
		QueryParam: make(map[string]string),
	}
	return agent
}

func (agent Agent) RunAction() (bool, *Response, error) {
	config, err := agent.GetOpenapiConfig()
	error_string := ""
	should_continue := true

	client := resty.New()

	// return a signal to not continue, and send the error
	if err != nil {
		return false, nil, err
	}

	for _, openapi_parameter := range config.Parameters {
		parameter_name := openapi_parameter.Value.Name

		parameter := agent.Action.GetParameter(parameter_name)

		fmt.Printf("parameter_name: %s\n", parameter_name)
		fmt.Printf("parameter: %s\n", parameter)

		if parameter == "" {
			if openapi_parameter.Value.Required {
				error_string += fmt.Sprintf("ERROR! the parameter value for `%s` is required and missing!\n", parameter_name)
				should_continue = false
			} else {
				error_string += fmt.Sprintf("WARN! the parameter value for `%s` is missing!\n", parameter_name)
			}
		} else {
			if openapi_parameter.Value.In == "path" {
				agent.PathParam[parameter_name] = parameter
			} else if openapi_parameter.Value.In == "query" {
				agent.QueryParam[parameter_name] = parameter
			}
		}

	}

	operation := agent.DocExt.OperationID.Value(agent.Action.OperationID)
	if operation == nil {
		opPath := agent.DocExt.Paths.Value(agent.Action.Path).GetOperation(config.Method)
		operation = (&tool.Operation{Operation: opPath}).Extend(agent.Action.Path, config.Method)
	}

	// create request
	request := &Request{client.R()}

	//## setup request
	//# -   setup method
	request.Method = operation.Method
	//# -   setup method
	request.URL = agent.Workflow.Baseurl + operation.Path
	//# -   setup parameters
	request = request.SetupParam(agent.PathParam, agent.QueryParam, agent.Workflow.Envs)
	//# -   setup header
	request = request.SetupHeader(agent.Action.Headers, agent.Workflow.Envs)
	//# -   setup body
	request = request.SetupBody(agent.Action.Payload, agent.Workflow.Envs)

	println("url: ", request.URL)

	if debug.DEBUG {
		DumpRequest(request)
	}

	// send request
	response, err := request.Send()
	if err != nil {
		return false, nil, err
	}

	err = errors.New(error_string)
	return should_continue, &Response{response}, err
}
