package json_file_reader

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func ReadJsonFile(jsonFilePath string) ([]byte, error) {
	if _, err := os.Stat(jsonFilePath); errors.Is(err, os.ErrNotExist) {
		return nil, errors.New(fmt.Sprintf("file %s does not exist", jsonFilePath))
	}

	wd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	jsonFile, err := os.Open(filepath.Join(wd, jsonFilePath))
	if err != nil {
		return nil, err
	}

	defer func(jsonFile *os.File) {
		_ = jsonFile.Close()
	}(jsonFile)

	return io.ReadAll(jsonFile)
}
