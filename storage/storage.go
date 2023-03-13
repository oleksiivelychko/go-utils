package storage

import "io"

type LocalStorage interface {
	Save(path string, file io.Reader) (int64, error)
}
