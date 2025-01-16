package tool

type Action struct {
	Name        string                 `json:"name"`
	Description string                 `json:"description"`
	Path        string                 `json:"path"`
	Method      string                 `json:"method"`
	Headers     map[string]interface{} `json:"headers"`
	Parameters  map[string]interface{} `json:"parameters"`
	Payload     string                 `json:"payload"`
	Endpoint    string
}

type Workflow struct {
	Name        string                 `json:"name"`
	Baseurl     string                 `json:"Baseurl"`
	Description string                 `json:"description"`
	Envs        map[string]interface{} `json:"envs"`
	Actions     []Action               `json:"actions"`
}