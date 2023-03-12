package json_datetime

import (
	"database/sql/driver"
	"fmt"
	"strings"
	"time"
)

type JsonDateTime struct {
	time.Time
}

func (dt *JsonDateTime) MarshalJSON() ([]byte, error) {
	stamp := time.Now().Format(time.RFC3339)
	return []byte("\"" + stamp + "\""), nil
}

func (dt *JsonDateTime) UnmarshalJSON(b []byte) (err error) {
	dtString := strings.Trim(string(b), "\"")
	date, err := time.Parse(time.RFC3339, dtString)
	if err != nil {
		return err
	}
	dt.Time = date
	return
}

//goland:noinspection GoMixedReceiverTypes
func (dt *JsonDateTime) Value() (driver.Value, error) {
	var zeroTimestamp time.Time
	if dt.Time.UnixNano() == zeroTimestamp.UnixNano() {
		return nil, nil
	}
	return dt.Time, nil
}

//goland:noinspection GoMixedReceiverTypes
func (dt *JsonDateTime) Scan(v interface{}) error {
	var dtBytes []byte
	for _, b := range v.([]uint8) {
		dtBytes = append(dtBytes, b)
	}

	dtString, err := time.Parse(time.DateTime, string(dtBytes))
	if err != nil {
		return fmt.Errorf("unable to convert %v to timestamp", v)
	}

	*dt = JsonDateTime{dtString}
	return nil
}
