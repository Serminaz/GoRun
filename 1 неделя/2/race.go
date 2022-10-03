package main

import (
	"fmt"
	"time"
)

//глобальный канал для полной синхронизации данных
var ch = make(chan int) // Какие данные не знают о типе,
                        // определенном в канале.

// Определите принтер
func printer(s string) {
	for _, value := range s {
		fmt.Printf("%c", value)
		time.Sleep(time.Millisecond * 300)
	}
}


func person1() {
	printer("hello")
	ch <- 777        // Отправить значение в канал 
}

func person2() {
	<-ch            //получить и отбросить данные
	printer("world") 
}

func main() {
	go person1()
	go person2()
	time.Sleep(time.Second * 3) 
}
