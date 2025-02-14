package loader

import (
	"errors"
)

type (
	DbLoader struct {
		DbPath string
		DbMode string
	}

	Loader interface {
		Load(path string) ([]string, error)
	}
)

func (dl *DbLoader) Load() ([]string, error) {
	return nil, errors.New("not implemented")
}
