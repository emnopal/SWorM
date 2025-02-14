package loader

import (
	"fmt"
)

// FolderLoader has interface Loader
type FolderLoader struct {
	DirPath string
}

func (fl *FolderLoader) Load() ([]string, error) {
	action_list := []string{}
	ActionEntry, err := getFileList(fl.DirPath)
	if err != nil {
		return action_list, fmt.Errorf("error getting action list: %w", err)
	}

	if len(ActionEntry) == 0 {
		return action_list, fmt.Errorf("no actions found")
	}

	for _, action := range ActionEntry {
		action_path := fl.DirPath + "/" + action.Name()
		action_bytes, err := loadFile(action_path)
		if err != nil {
			return action_list, fmt.Errorf("error loading an action: %w", err)
		}
		action_list = append(action_list, string(action_bytes))
	}

	return action_list, nil
}
