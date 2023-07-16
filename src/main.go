package main

import (
	"os"
	"os/signal"
	"time"
)

var (
	UpdateInterval = time.Millisecond * 250
	ResetInterval  = time.Minute
)

func main() {
	go ResetGoroutine()
	go ReaderGoroutine()
	go PrinterGoroutine()

	s := make(chan os.Signal, 1)
	signal.Notify(s, os.Interrupt)
	<-s
}
