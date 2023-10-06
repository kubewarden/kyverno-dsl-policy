package gzip

import "io"

func ReadCloser(r io.ReadCloser) io.ReadCloser {
	panic("stub")
}

func ReadCloserLevel(r io.ReadCloser, level int) io.ReadCloser {
	panic("stub")
}

func UnzipReadCloser(r io.ReadCloser) (io.ReadCloser, error) {
	panic("stub")
}

func Is(r io.Reader) (bool, error) {
	panic("stub")
}

type Embedme interface{}
