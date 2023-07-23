package file

import (
	"errors"
	"io/ioutil"
	"os"
	"strings"
)

func ReadFile(filename string) (string, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}

	return strings.Trim(string(data), "\n\t "), nil
}

func FileExists(filename string) bool {
	_, err := os.Stat(filename)
	if errors.Is(err, os.ErrNotExist) {
		return false
	}
	return true
}
