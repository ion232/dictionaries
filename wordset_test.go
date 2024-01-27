package dictionaries_test

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/santhosh-tekuri/jsonschema"
)

func TestWordsetV1(t *testing.T) {
	t.Parallel()

	sch, err := jsonschema.Compile(filepath.Join("wordset", "base", "schema.json"))
	if err != nil {
		t.Error(err)
	}

	data, err := os.ReadFile(filepath.Join("wordset", "base", "wordset.json"))
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
