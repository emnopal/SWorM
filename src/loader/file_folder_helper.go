package loader

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
)

func getFileList(dirPath string) ([]fs.DirEntry, error) {
	entry, err := os.Stat(dirPath)
	action_list := []fs.DirEntry{}

	if !entry.IsDir() {
		return action_list, errors.New("path is not a directory")
	}

	if err != nil {
		os.Mkdir(dirPath, 0755)
	}

	files, _ := os.ReadDir(dirPath)
	if len(files) == 0 {
		return action_list, errors.New("action folder is empty")
	}

	action_list = files

	return action_list, nil
}

func loadFile(path string) ([]byte, error) {
	bytes, err := os.ReadFile(path)
	if err != nil {
		return bytes, fmt.Errorf("error reading file: %w", err)
	}
	return bytes, nil
}
