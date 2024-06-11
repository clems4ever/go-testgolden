package testgolden

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"
)

func updateGoldenFile(t *testing.T, goldenFilePath string, reader io.Reader) {
	t.Helper()

	err := os.MkdirAll(filepath.Dir(goldenFilePath), 0700)
	require.NoError(t, err)

	f, err := os.Create(goldenFilePath)
	require.NoError(t, err)
	defer f.Close()

	_, err = io.Copy(f, reader)
	require.NoError(t, err)
}

func RequireDataEqualToJSON[T any](t *testing.T, goldenFilePath string, data T) {
	var buf bytes.Buffer

	enc := json.NewEncoder(&buf)
	enc.SetIndent("", "\t")

	err := enc.Encode(data)
	require.NoError(t, err)

	RequireBytesEqualToGolden(t, goldenFilePath, &buf)
}

func RequireDataEqualToYaml[T any](t *testing.T, goldenFilePath string, data T) {
	var buf bytes.Buffer

	enc := yaml.NewEncoder(&buf)

	err := enc.Encode(data)
	require.NoError(t, err)

	RequireBytesEqualToGolden(t, goldenFilePath, &buf)
}

func RequireDataEqualToString(t *testing.T, goldenFilePath string, data string) {
	buf := bytes.NewBufferString(data)
	RequireBytesEqualToGolden(t, goldenFilePath, buf)
}

func RequireBytesEqualToGolden(t *testing.T, goldenFilePath string, reader io.Reader) {
	t.Helper()

	var goldenExists = true
	_, err := os.Stat(goldenFilePath)
	if err != nil && errors.Is(err, os.ErrNotExist) {
		goldenExists = false
	}

	// if dev mode
	if os.Getenv("CI") == "" {
		if !goldenExists {
			updateGoldenFile(t, goldenFilePath, reader)
			return
		}
	} else {
		if !goldenExists {
			require.FailNow(t, "golden file does not exist")
		}
	}

	f, err := os.OpenFile(goldenFilePath, os.O_RDONLY, 0600)
	require.NoError(t, err)
	defer f.Close()

	raw, err := io.ReadAll(f)
	require.NoError(t, err)

	expected, err := io.ReadAll(reader)
	require.NoError(t, err)

	require.Equal(t, string(expected), string(raw))
}
