package actions

import (
	"errors"
	"fmt"
)

type ActionFactory interface {
	Load(string, string) (*ActionItem, error)
}

func Load(actionType string, data string) (*ActionItem, error) {
	switch actionType {
	case "request":
		action := &ActionRequest{}
		loadedAction, err := action.Load(data)
		if err != nil {
			return nil, fmt.Errorf("error loading action: %w", err)
		}
		return loadedAction, nil
	default:
		return nil, errors.New("invalid action type")
	}
}
