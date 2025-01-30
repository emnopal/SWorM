package tool

type Action struct {
	Name        string            `json:"name"`
	Description string            `json:"description"`
	Path        string            `json:"path"`
	Method      string            `json:"method"`
	OperationID string            `json:"operationID"`
	Headers     map[string]string `json:"headers"`
	Parameters  map[string]string `json:"parameters"`
	Payload     string            `json:"payload"`
	Endpoint    string
}

func (Action Action) GetParameter(key string) string {
	if len(Action.Parameters) == 0 {
		return ""
	}
	return Action.Parameters[key]
}

func (Action Action) GetHeader(key string) string {
	if len(Action.Parameters) == 0 {
		return ""
	}
	return Action.Parameters[key]
}
