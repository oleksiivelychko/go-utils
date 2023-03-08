package logger

import (
	"github.com/hashicorp/go-hclog"
)

func NewLogger(name string) hclog.Logger {
	return hclog.New(&hclog.LoggerOptions{
		Name:       name,
		Level:      hclog.LevelFromString("DEBUG"),
		Color:      1,
		TimeFormat: "02/01/2006 15:04:05",
	})
}
