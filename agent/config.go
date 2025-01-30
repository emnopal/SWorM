package agent

import (
	"errors"
	"fmt"

	"github.com/SWorM/v2/tool"
)

func (agent Agent) GetOpenapiConfig() (*tool.ExtendedOperation, error) {
	opid := agent.Action.OperationID
	if opid != "" {
		fmt.Printf("Operation Id '%s' found!\n", opid)
		return agent.GetOpenapiConfigByID(opid), nil
	}

	path := agent.Action.Path
	method := agent.Action.Method
	if path != "" && method != "" {
		fmt.Printf("Operation Path '%s', Method '%s' found!\n", path, method)
		return agent.GetOpenapiConfigByPath(path, method), nil
	}

	// if path and method is empty then straight up return nil and print error
	err := errors.New("ERROR: (operationID) or (path & method) value in workflow.json required! ")
	return nil, err
}

func (agent Agent) GetOpenapiConfigByID(operationId string) *tool.ExtendedOperation {
	operation := agent.DocExt.OperationID.Value(operationId)
	return operation
}

func (agent Agent) GetOpenapiConfigByPath(path string, method string) *tool.ExtendedOperation {
	pathItem := agent.DocExt.Paths.Value(path)
	op := pathItem.GetOperation(method)

	localOperation := tool.Operation{Operation: op}
	extendedOperation := localOperation.Extend(path, method)
	return extendedOperation
}
