package json_indent

import (
	"bytes"
	"encoding/json"
	"fmt"
	"testing"
)

func TestJsonIndent(t *testing.T) {
	prettyJson := JsonIndent(`{"hello":"world"}`)

	fmt.Printf("%s\n", prettyJson)

	buffer := new(bytes.Buffer)
	if err := json.Compact(buffer, prettyJson); err != nil {
		t.Error(err)
	}
}
