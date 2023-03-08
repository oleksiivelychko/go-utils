package utils

import (
	"bytes"
	"testing"
)

type JsonStruct struct {
	Attr string
}

func TestToJSON(t *testing.T) {
	buffer := new(bytes.Buffer)
	err := ToJSON(&JsonStruct{Attr: "testAttr"}, buffer)

	if err != nil {
		t.Errorf("unable to serialize into json: %s", err)
	}

	if buffer.String() != "{\"Attr\":\"testAttr\"}\n" {
		t.Errorf("unable to compare json string: %s", buffer.String())
	}
}

func TestFromJSON(t *testing.T) {
	buffer := new(bytes.Buffer)
	buffer.Write([]byte("{\"Attr\":\"testAttr\"}\n"))

	jsonStruct := &JsonStruct{}
	err := FromJSON(jsonStruct, buffer)

	if err != nil {
		t.Errorf("unable to deserialize from json: %s", err)
	}

	if jsonStruct.Attr != "testAttr" {
		t.Errorf("unable to compare json string: %s", jsonStruct.Attr)
	}
}
