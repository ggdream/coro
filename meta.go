package main

import (
	"fmt"
	"runtime"
	"strings"
)


const (
	name    			= "coro"
	version 			= "0.1.0"
)

var (
	platform  			= runtime.GOOS
	arch				= runtime.GOARCH
	userAgent 			= fmt.Sprintf("%s/%s (%s; %s)", strings.Title(name), version, platform, arch)
)
