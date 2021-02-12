package main

import (
	"os"
	"github.com/ggdream/mini"
)


func cmder() string {
	rest := mini.New(os.Args[1:]).GetRest()
	return rest[len(rest)-1]
}
