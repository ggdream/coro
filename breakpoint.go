package main

import "os"


func judge(name string) int64 {
	info, err := os.Stat(name)
	if err != nil {
		return 0
	}

	return info.Size()
}
