package tool

import (
	"github.com/getkin/kin-openapi/openapi3"
)

type ExtendedT struct {
	*T
	OperationID *OperationID
}

type T struct {
	*openapi3.T
}

type OperationID struct {
	m map[string]*ExtendedOperation
}

type ExtendedOperation struct {
	Path         string
	Method       string
	Extensions   map[string]any
	Tags         []string
	Summary      string
	Description  string
	Parameters   openapi3.Parameters
	RequestBody  *openapi3.RequestBodyRef
	Responses    *openapi3.Responses
	Callbacks    openapi3.Callbacks
	Deprecated   bool
	Security     *openapi3.SecurityRequirements
	Servers      *openapi3.Servers
	ExternalDocs *openapi3.ExternalDocs
}

type Operation struct {
	*openapi3.Operation
}

func (Operation *Operation) Extend(Path string, Method string) ExtendedOperation {
	return ExtendedOperation{
		Path:         Path,
		Method:       Method,
		Extensions:   Operation.Extensions,
		Tags:         Operation.Tags,
		Summary:      Operation.Summary,
		Description:  Operation.Description,
		Parameters:   Operation.Parameters,
		RequestBody:  Operation.RequestBody,
		Responses:    Operation.Responses,
		Callbacks:    Operation.Callbacks,
		Deprecated:   Operation.Deprecated,
		Security:     Operation.Security,
		Servers:      Operation.Servers,
		ExternalDocs: Operation.ExternalDocs,
	}
}

func (T *T) Extend() *ExtendedT {
	paths := T.Paths.Map()
	op := make(map[string]*ExtendedOperation)

	for path, pathitem := range paths {

		operations := pathitem.Operations()

		for method, operation := range operations {

			operation := &Operation{operation}
			NewOperation := operation.Extend(path, method)
			op[operation.OperationID] = &NewOperation

		}
	}
	OperationID := &OperationID{
		op,
	}

	extendedT := &ExtendedT{
		T:           T,
		OperationID: OperationID,
	}

	return extendedT
}

func (operationID *OperationID) Value(key string) *ExtendedOperation {
	if len(operationID.m) == 0 {
		return nil
	}
	return operationID.m[key]
}
