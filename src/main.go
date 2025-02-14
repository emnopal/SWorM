package main

import (
	"fmt"

	"github.com/SWorM/v2/src/actions"
	"github.com/SWorM/v2/src/loader"
	"github.com/SWorM/v2/src/workflow"
)

func main() {
	actionLoader := loader.FolderLoader{DirPath: "action"} // Create a new instance of FileLoader
	action_list_string, err := actionLoader.Load()         // Load actions from the directory
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

	workflowLoader := loader.FileLoader{FilePath: "workflow.json"} // Create a new instance of FileLoader
	workflow_string, err := workflowLoader.Load()                  // Load workflow from the file
	if err != nil {
		panic(err)
	}

	workflow_handler := workflow.NewWorkflow() // Create a new instance of Workflow
	err = workflow_handler.LoadWorkflow(workflow_string)
	if err != nil {
		panic(err)
	}

	workflow_dump, err := workflow_handler.Dump()
	if err != nil {
		panic(err)
	}
	fmt.Println(workflow_dump)

}
