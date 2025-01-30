package tool

type Workflow struct {
	Name        string            `json:"name"`
	Baseurl     string            `json:"Baseurl"`
	Description string            `json:"description"`
	Envs        map[string]string `json:"envs"`
	Actions     []Action          `json:"actions"`
}
