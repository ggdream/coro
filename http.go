package main

import (
	"fmt"
	"io"
	"os"
	"path"
	"strings"

	"github.com/ggdream/requests"
)



func downer(url string) error {
	ceil := getSuffix(url)
	name := joinPath(".", ceil)

	headers := map[string]string{
		"User-Agent": userAgent,
		"Range": fmt.Sprintf("bytes=%d-", judge(name)),
	}

	res := requests.Get(url, &requests.Options{
		Headers: headers,
	})
	defer res.Close()


	return save(res.Raw().ContentLength, res.Body(), name)
}

func save(Totalsize int64, r io.ReadCloser, name string) error {
	file, err := os.OpenFile(name, os.O_APPEND | os.O_CREATE | os.O_RDWR, os.ModePerm)
	if err != nil {
		return err
	}
	defer file.Close()

	read := &reader{
		Reader: r,
		End: uint64(Totalsize),
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
