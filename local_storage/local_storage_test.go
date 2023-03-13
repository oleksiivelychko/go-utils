package local_storage

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"io"
	"os"
	"path/filepath"
	"testing"
	"unsafe"
)

const savePath = "/1/test.png"
const fileContent = "Hello World"

func setup(t *testing.T) (*LocalStorage, func()) {
	tmpDir, err := os.MkdirTemp(os.TempDir(), "files")
	if err != nil {
		t.Fatal(err)
	}

	// 1mb = 1000000 bytes
	localStorage, err := NewLocalStorage(tmpDir, 1000000)
	if err != nil {
		t.Fatal(err)
	}

	return localStorage, func() {
		_ = os.RemoveAll(tmpDir)
	}
}

func TestLocalSave(t *testing.T) {
	localStorage, cleanup := setup(t)
	defer cleanup()

	content := bytes.NewBuffer([]byte(fileContent))
	writtenBytes, err := localStorage.Save(savePath, content)
	assert.NoError(t, err)
	assert.GreaterOrEqual(t, writtenBytes, int64(unsafe.Sizeof(content)))

	file, err := os.Open(filepath.Join(localStorage.basePath, savePath))
	assert.NoError(t, err)

	data, err := io.ReadAll(file)
	assert.NoError(t, err)
	assert.Equal(t, fileContent, string(data))
}

func TestLocalGet(t *testing.T) {
	localStorage, cleanup := setup(t)
	defer cleanup()

	content := bytes.NewBuffer([]byte(fileContent))
	_, err := localStorage.Save(savePath, content)
	assert.NoError(t, err)

	file, err := localStorage.Get(savePath)
	assert.NoError(t, err)
	defer file.Close()

	data, err := io.ReadAll(file)
	assert.Equal(t, fileContent, string(data))
}
