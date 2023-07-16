package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)

var (
	LogPath         = "/var/log/nginx/access.log"
	TopRequestCount = 10
	UpdateInterval  = time.Millisecond * 500
	ResetInterval   = time.Minute
	startedAt       = time.Now()
)

func main() {
	fmt.Print("\u001b[?47h")

	go ResetGoroutine()
	go ReaderGoroutine()
	go PrinterGoroutine()

	s := make(chan os.Signal, 1)
	signal.Notify(s, os.Interrupt)
	<-s

	fmt.Print("\u001b[?47l\u001b[?25h")
}
