package main

import (
	"fmt"
	"time"
)

func printer(s string) {
	for _, value := range s {
		fmt.Printf("%c", value)
		time.Sleep(time.Millisecond * 300)
	}
}

func person1() {
	printer("hello")
}

func person2() {
	printer("world")
}

func main() {
	go person1()
	go person2()
	time.Sleep(time.Second * 5) 
}