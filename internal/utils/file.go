package utils

import (
	"os"
	"path/filepath"
	"strings"
)

func FileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func Extension(path string) string {
	return strings.TrimPrefix(filepath.Ext(path), ".")
}

func ReadFile(path string) ([]byte, error) {
	return os.ReadFile(path)
}
