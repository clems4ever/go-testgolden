package testgolden_test

import (
	"testing"

	"github.com/clems4ever/go-testgolden"
)

func TestExampleString(t *testing.T) {
	// Generate a path of a file to be stored in the test directory dedicated to this specific test.
	goldenPath := testgolden.GetTestDataFilePath(t, "golden.json")

	// If the env variable CI=1 is not set, the golden is updated with the latest version of the struct,
	// but if the variable is set, the CI will check that the golden and the actual data is equal.
	testgolden.RequireDataEqualToString(t, goldenPath, "test\nwith a newline")
}
