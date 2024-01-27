package dictionaries_test

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/santhosh-tekuri/jsonschema"
)

func ValidateFile(t *testing.T, sch *jsonschema.Schema, filename string) {
	path := filepath.Join("wordset", "base", "sections", filename)

	data, err := os.ReadFile(path)
	if err != nil {
		t.Errorf("%#v", err)
	}

	var v interface{}
	if err := json.Unmarshal(data, &v); err != nil {
		t.Errorf("%#v", err)
	}

	reader := bytes.NewReader(data)
	if err = sch.Validate(reader); err != nil {
		t.Error(err)
	}
}

func TestWordsetBase(t *testing.T) {
	t.Parallel()

	sch, err := jsonschema.Compile(filepath.Join("wordset", "base", "schema.json"))
	if err != nil {
		t.Error(err)
	}

	for _, r := range "abcdefghijklmnopqrstuvwxyz" {
		filename := string(r) + ".json"
		ValidateFile(t, sch, filename)
	}
}
