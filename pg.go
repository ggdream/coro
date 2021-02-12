package main

import (
	"fmt"
	"io")



type reader struct {
	io.Reader
	Now				uint64
	End				uint64
}

func (r *reader) Read(buf []byte) (int, error) {
	n, err := r.Reader.Read(buf)

	r.Now += uint64(n)
	fmt.Printf("\r✐  %.0f%%", float64(r.Now * 10000 / r.End) / 100)

	return n, err
}
