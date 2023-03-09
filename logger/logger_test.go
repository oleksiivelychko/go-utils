package logger

import (
	"bytes"
	"github.com/hashicorp/go-hclog"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestLoggerLevel(t *testing.T) {
	var buf bytes.Buffer

	logger := hclog.New(&hclog.LoggerOptions{
		Name:   "test",
		Level:  hclog.LevelFromString("DEBUG"),
		Output: &buf,
	})

	logger.Debug("this is test", "who", "programmer", "why", "testing")

	str := buf.String()
	dataIdx := strings.IndexByte(str, ' ')
	rest := str[dataIdx+1:]

	assert.Equal(t, "[DEBUG] test: this is test: who=programmer why=testing\n", rest)
}
