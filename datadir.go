package testgolden

import (
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

func GetProjectRootDir() string {
	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)
	rootDir := filepath.Join(dir, "..", "..")

	return rootDir
}

func GetTestsDir() string {
	rootDir := GetProjectRootDir()
	return filepath.Join(rootDir, "tests")
}

// GetCurrentTestDataDir calculates the path to the test data corresponding to this exact test.
// This function must be called directly withing the test function but not deeper in the stack.
func GetCurrentTestDataDir(t *testing.T) string {
	return getFuncFilePath(t, 2)
}

func getFuncFilePath(t *testing.T, skip int) string {
	_, filename, _, _ := runtime.Caller(skip)
	name := strings.ReplaceAll(filepath.Base(filename), ".go", "")

	return filepath.Join(name, t.Name())
}

func GetTestDataFilePath(t *testing.T, path string) string {
	dir := getFuncFilePath(t, 2)
	return filepath.Join(dir, path)
}
