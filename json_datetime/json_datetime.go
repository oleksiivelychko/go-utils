package json_datetime

import (
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
