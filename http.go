package main

import (
	"io"
	"os"
	"path"
	"strings"

	"github.com/ggdream/requests"
)



func downer(url string) error {
	res := requests.Get(url, nil)
	defer res.Close()

	ceil := getSuffix(url)
	
	return save(res.Raw().ContentLength, res.Body(), joinPath(".", ceil))
}

func save(size int64, r io.ReadCloser, name string) error {
	file, err := os.Create(name)
	if err != nil {
		return err
	}
	defer file.Close()

	read := &reader{
		Reader: r,
		End: uint64(size),
	}

	_, err = io.Copy(file, read)
	return err
}

func getSuffix(url string) string {
	ceils := strings.Split(url, "/")
	return ceils[len(ceils)-1]
}

func joinPath(_path, ceil string) string {
	return path.Join(_path, ceil)
}
