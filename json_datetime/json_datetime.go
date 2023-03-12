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

func (t *JsonDateTime) MarshalJSON() ([]byte, error) {
	stamp := time.Now().Format(time.RFC3339)
	return []byte("\"" + stamp + "\""), nil
}

func (t *JsonDateTime) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), "\"")
	date, err := time.Parse(time.RFC3339, s)
	if err != nil {
		return err
	}
	t.Time = date
	return
}

//goland:noinspection GoMixedReceiverTypes
func (t *JsonDateTime) Value() (driver.Value, error) {
	var zeroTimestamp time.Time
	if t.Time.UnixNano() == zeroTimestamp.UnixNano() {
		return nil, nil
	}
	return t.Time, nil
}

//goland:noinspection GoMixedReceiverTypes
func (t *JsonDateTime) Scan(v interface{}) error {
	if value, ok := v.(time.Time); ok {
		*t = JsonDateTime{value}
		return nil
	}
	return fmt.Errorf("unable to convert %v to timestamp", v)
}
