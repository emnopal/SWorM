package agent

import (
	"encoding/json"
	"fmt"

	"github.com/SWorM/v2/debug"
)

func DumpRequest(request *Request) {
	if !debug.CheckVerbosity(debug.REQUEST_CODE) {
		return
	}

	requestJSON, err := json.MarshalIndent(request, "", "  ")
	if err != nil {
		fmt.Println("Error marshalling request to JSON:", err)
		return
	}

	fmt.Println(string(requestJSON))
}
