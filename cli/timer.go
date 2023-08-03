package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var quit = make(chan os.Signal, 1)

func startTimer() {
	inc := 1
	t := time.NewTicker(time.Second * time.Duration(inc))
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	count := 0
	for {
		select {
		case <-t.C:
			count += inc
			fmt.Printf("\r%d", count)
		case <-quit:
			t.Stop()
			fmt.Println("timer stop")
			return
		}
	}
}
