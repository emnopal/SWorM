package actions

type ActionItem interface {
	Load(string) (*ActionItem, error)
	Dump() (string, error)
	GetName() string
	GetType() string
}
