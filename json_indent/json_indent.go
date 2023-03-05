package json_indent

import (
	"bytes"
	"encoding/json"
)

func JsonIndent(jsonString string) []byte {
	var out bytes.Buffer
	_ = json.Indent(&out, []byte(jsonString), "", "	")

	return out.Bytes()
}
