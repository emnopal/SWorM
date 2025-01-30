package debug

const (
	WORKFLOW_CODE int = 1
	ACTION_CODE   int = 2
	REQUEST_CODE  int = 4
)

func LoadDebugConfig() {
	ReadOrCreateDebugConfig()
}

func CheckVerbosity(code int) bool {
	return (VERBOSITY & code) == code
}

func Check(e error) {
	if e != nil {
		panic(e)
	}
}
