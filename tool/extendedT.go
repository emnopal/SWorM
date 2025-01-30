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

func (T *T) Extend() *ExtendedT {
	paths := T.Paths.Map()
	op := make(map[string]*ExtendedOperation)

	for path, pathitem := range paths {

		operations := pathitem.Operations()

		for method, operation := range operations {

			operation := &Operation{operation}
			NewOperation := operation.Extend(path, method)
			op[operation.OperationID] = NewOperation

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
