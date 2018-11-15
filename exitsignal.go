package main

import (
	"os"
	"os/signal"
	"syscall"
)

func ossignal() chan int {
	ossignal := make(chan os.Signal, 1)
	pulse := make(chan int)
	signal.Notify(ossignal,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)
	go func() {
		<-ossignal
		pulse <- 1
	}()
	return pulse
}
