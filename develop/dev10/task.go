package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	timeout := flag.Duration("timeout", 10*time.Second, "connection timeout")
	flag.Parse()

	host := flag.Arg(0)
	port := flag.Arg(1)
	if host == "" || port == "" {
		fmt.Println("Usage: go-telnet --timeout=10s host port")
		return
	}

	conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%s", host, port), *timeout)
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}
	defer conn.Close()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigCh
		fmt.Println("\nCtrl+C pressed. Closing connection.")
		conn.Close()
		os.Exit(0)
	}()

	go func() {
		_, err := io.Copy(conn, os.Stdin)
		if err != nil {
			fmt.Println("Error writing to server:", err)
			return
		}
	}()

	_, err = io.Copy(os.Stdout, conn)
	if err != nil {
		fmt.Println("Error reading from server:", err)
		return
	}
}
