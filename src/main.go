package main

import (
	"fmt"

	"github.com/SWorM/v2/src/actions"
	"github.com/SWorM/v2/src/loader"
)

func main() {
	loader := loader.FolderLoader{DirPath: "action"} // Create a new instance of FileLoader
	action_list_string, err := loader.Load()         // Load actions from the directory
	if err != nil {
		panic(err)
	}

	action_handler := actions.NewActionsDict() // Create a new instance of ActionsDict
	err = action_handler.LoadActions(action_list_string)
	if err != nil {
		panic(err)
	}

	for _, action := range action_handler.GetActionKeys() {
		fmt.Println(action)
		action_item, err := action_handler.GetAction(action)
		if err != nil {
			panic(err)
		}
		action_item_dump, err := (*action_item).Dump()
		if err != nil {
			panic(err)
		}
		fmt.Println(action_item_dump)
	}

	// action_dict, err := actions.LoadActions("action")
	// if err != nil {
	// 	panic(err)
	// }

	// action_dump, err := action_dict.Dump()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(action_dump)
}
