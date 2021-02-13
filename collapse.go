package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)


func listen() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT)

	sig := <- sigs
	fmt.Println("exit", sig)
	os.Exit(-1)
}
