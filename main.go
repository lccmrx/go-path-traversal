package main

import (
	"fmt"
	"os"
	"os/signal"
)

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	go defaultServer()
	go echoServer()

	fmt.Println("Waiting for shutdown signal...")
	<-c
	os.Exit(0)
}
