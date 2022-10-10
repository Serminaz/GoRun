/*Реализовать собственную функцию sleep.*/

package main

import (
//	"fmt"
	"time"
	"log"
)

func Sleep1(sec int) {
    <-time.After(time.Second* time.Duration(sec))
}
func main() {
	log.Println( " begin")
	//fmt.Println( "begin")
	Sleep1(3)
	log.Println( " end")
}
