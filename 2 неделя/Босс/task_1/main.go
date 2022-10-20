package main

import "fmt"

func main() {
	s := "hello"
	s = s[:len(s)-1]
	fmt.Println(s)
}
