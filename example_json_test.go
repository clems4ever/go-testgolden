package testgolden_test

import (
	"testing"

	"github.com/clems4ever/go-testgolden"
)

type MyStruct struct {
	ID      int
	Title   string
	Content string
}

func TestExample(t *testing.T) {
	s := MyStruct{
		ID:      1,
		Title:   "this is a title",
		Content: "Here is a little story",
	}

	// Generate a path of a file to be stored in the test directory dedicated to this specific test.
	goldenPath := testgolden.GetTestDataFilePath(t, "golden.json")

	// If the env variable CI=1 is not set, the golden is updated with the latest version of the struct,
	// but if the variable is set, the CI will check that the golden and the actual data is equal.
	testgolden.RequireDataEqualToJSON(t, goldenPath, s)
}
