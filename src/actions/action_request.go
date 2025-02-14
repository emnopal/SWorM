package actions

import (
	"encoding/json"
	"fmt"
)

type ActionRequest struct {
	Name        string            `json:"name"`
	Description string            `json:"description"`
	Site        string            `json:"site"`
	Path        string            `json:"path"`
	Method      string            `json:"method"`
	Header      map[string]string `json:"header"`
	Parameter   map[string]string `json:"parameter"`
	Payload     string            `json:"payload"`
}

func NewActionRequest() *ActionRequest {
	return &ActionRequest{}
}

func (a *ActionRequest) Dump() (string, error) {
	json_dump, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "", fmt.Errorf("error dumping action map: %w", err)
	}
	return string(json_dump), nil
}

func (a *ActionRequest) Load(jsonString string) (*ActionItem, error) {
	action_bytes := []byte(jsonString)
	err := json.Unmarshal(action_bytes, a)

	if a.Site[len(a.Site)-1:] == "/" {
		a.Site = a.Site[:len(a.Site)-1]
	}

	if err != nil {
		return nil, fmt.Errorf("error reading action file: %w", err)
	}

	var item ActionItem = a
	return &item, nil
}

func (a *ActionRequest) GetName() string {
	return a.Name
}

func (a *ActionRequest) GetType() string {
	return "request"
}
