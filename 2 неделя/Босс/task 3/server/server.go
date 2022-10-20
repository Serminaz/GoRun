package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

var host string
var port string

func init() {
	flag.StringVar(&host, "h", "127.0.0.1", "hostname")
	flag.StringVar(&port, "p", "8081", "port")
	flag.Parse()
}

func main() {
	log.Println("server launching")
	defer log.Println("server stopped")

	ln, err := net.Listen("tcp", host+":"+port)
	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close()

	conn, err := ln.Accept()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	osSignals := make(chan os.Signal, 1)
	signal.Notify(osSignals, syscall.SIGINT)
	done := make(chan interface{})

	in := make(chan string)
	go func() {
		reader := bufio.NewReader(conn)
		for {
			select {
			case <-done:
				close(in)
				return
			default:
				msg, err := reader.ReadString('\n')
				if err == nil {
					in <- msg
				}
			}
		}
	}()

	go func() {
		for {
			select {
			case <-done:
				return
			case msg := <-in:
				fmt.Fprint(conn, strings.ToUpper(msg))
			}
		}
	}()

	<-osSignals
	close(done)
}
