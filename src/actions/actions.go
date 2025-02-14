package actions

import (
	"errors"
)

type ActionsDict struct {
	Actions map[string]*ActionItem
}

func (a *ActionsDict) GetActions() map[string]*ActionItem {
	return a.Actions
}

func (a *ActionsDict) GetActionKeys() []string {
	keys := make([]string, 0, len(a.Actions))
	for k := range a.Actions {
		keys = append(keys, k)
	}
	return keys
}

func (a *ActionsDict) GetAction(actionName string) (*ActionItem, error) {
	if _, ok := a.Actions[actionName]; !ok {
		return nil, errors.New("action not found")
	}
	return a.Actions[actionName], nil
}

func (a *ActionsDict) LoadActions(actionString []string) error {
	if len(actionString) == 0 {
		return errors.New("no actions found")
	}

	for _, action := range actionString {
		actionItem, err := Load("request", action)
		if err != nil {
			return err
		}
		a.Actions[(*actionItem).GetName()] = actionItem
	}
	return nil
}

func NewActionsDict() *ActionsDict {
	return &ActionsDict{Actions: make(map[string]*ActionItem)}
}
