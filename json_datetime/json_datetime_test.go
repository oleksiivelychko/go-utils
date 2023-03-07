package json_datetime

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

type TestStruct struct {
	DateTime JsonDateTime `json:"dateTime"`
}

func TestMarshalJsonDateTime(t *testing.T) {
	testStruct := &TestStruct{DateTime: JsonDateTime{}}
	dateTimeNowString, err := json.Marshal(testStruct)
	if err != nil {
		t.Errorf("unable to marshal json datetime: %s", err)
	}

	unmarshalStruct := &TestStruct{}
	_ = json.Unmarshal(dateTimeNowString, &unmarshalStruct)

	if time.Now().Format(time.DateOnly) != unmarshalStruct.DateTime.Format(time.DateOnly) {
		t.Error("date are not same")
	}
}

func TestUnmarshalJsonDateTime(t *testing.T) {
	dtSample := "2022-10-09T08:23:55.267Z"

	testStruct := &TestStruct{}
	jsonString := []byte(fmt.Sprintf(`{"dateTime":"%s"`, dtSample))

	err := json.Unmarshal(jsonString, &testStruct)
	if err != nil {
		t.Errorf("unable to unmarshal json datetime: %s", err)
	}

	timeRFC3339, err := time.Parse(time.RFC3339, dtSample)
	if err != nil {
		t.Errorf("unable to parse json datetime %s : %s", dtSample, err)
	}

	marshalStruct := &TestStruct{DateTime: JsonDateTime{timeRFC3339}}
	if marshalStruct.DateTime.Format(time.RFC3339Nano) != dtSample {
		t.Error("datetime are not same")
	}
}
