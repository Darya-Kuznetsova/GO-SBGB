package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	grShutdown()
}

func grShutdown() {
	channel := make(chan os.Signal, 1)
	signal.Notify(channel, syscall.SIGINT, syscall.SIGTERM)
	timer := time.After(5 * time.Second)
	select {
	case <-timer:
		fmt.Println("Over")
	case newSignal := <-channel:
		fmt.Println("Graceful shutdown", newSignal)
	}
}
