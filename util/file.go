package util

import (
	"bytes"
	"fmt"
	"github.com/sanjid133/vault-kube/errors"
	"io/ioutil"
	"os"
)

func EnsurePath(path string) error {
	if err := os.MkdirAll(path, 0755); err != nil {
		return errors.Wrap(fmt.Sprintf("Failed to mkdir %v: %v", path, err))
	}
	return nil
}

func WriteData(filePath string, value interface{}) error {
	return ioutil.WriteFile(filePath, []byte(fmt.Sprintf("%v", value)), 0644)
}

func ReadData(filePath string) (string, error) {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return string(bytes.TrimSpace(content)), nil
}
