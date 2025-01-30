package debug

import (
	"encoding/json"
	"os"
)

var (
	// ### VERBOSITY level (in binary/decimal)
	// # - 0001/1: print out workflow
	// # - 0010/2: print out action
	// # etc ...
	// # to print multiple things just add the decimal

	VERBOSITY  = 1 + 2 + 4
	DEBUG      = false
	SKIPCHECKS = false
)

// ReadOrCreateDebugConfig reads the .debug file or creates it with default values if it doesn't exist
func ReadOrCreateDebugConfig() {
	file, err := os.Open(".debug")
	if os.IsNotExist(err) {
		file, err = os.Create(".debug")
		if err != nil {
			// Handle the error appropriately
			return
		}
		defer file.Close()

		encoder := json.NewEncoder(file)
		config := struct {
			Verbosity  int  `json:"verbosity"`
			Debug      bool `json:"debug"`
			SkipChecks bool `json:"skipChecks"`
		}{
			Verbosity:  VERBOSITY,
			Debug:      DEBUG,
			SkipChecks: SKIPCHECKS,
		}

		encoder.Encode(config)
		return
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	config := struct {
		Verbosity  int  `json:"verbosity"`
		Debug      bool `json:"debug"`
		SkipChecks bool `json:"skipChecks"`
	}{}

	if err := decoder.Decode(&config); err != nil {
		return
	}

	VERBOSITY = config.Verbosity
	DEBUG = config.Debug
	SKIPCHECKS = config.SkipChecks
}
