package and

import "io"

type ReadCloser struct {
	io.Reader
	CloseFunc func() error
}

type WriteCloser struct {
	io.Writer
	CloseFunc func() error
}

func (rac *ReadCloser) Close() error {
	panic("stub")
}

func (wac *WriteCloser) Close() error {
	panic("stub")
}

type Embedme interface{}
