# go-testgolden

Oftentimes, tests are about comparing an instance of datastructure to an expected one. However,
it might quickly become a lot of work to update those expectations when modifying the code. Like
the butterfly effect, one tiny change to one line of code might have big consequences on the
outcome of the algorithm and can therefore impact a lot of tests.

Golden files are files containing a serialized version of the expected instance of the datastructure
that is used by the test as the expectation for the outcome of the algorithm. This repository offer
primitives for creating and maintaining golden files, also called goldens.


# Getting Started

```bash
go get github.com/clems4ever/go-testgolden
```

```go
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
```
