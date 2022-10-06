/*
Разработать конвейер чисел.Даны два канала:
в первый пишутся числа (x) из массива,
во второй — результат операции x*2,
после чего данные из второго канала должны выводиться в stdout.
почитать про то каким образом управлять рутиной через канал.
*/
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	gorutina()
}

func gorutina() {
	rand.Seed(time.Now().Unix())
	var size int = 5

	massiv := make([]int, size)
	for i := 0; i < size; i++ {
		massiv[i] = 1 + rand.Intn(16)
	}
   
	workWithchanel(massiv)
	time.Sleep(10 * time.Millisecond)
}

 func workWithchanel (massiv []int) {
	c1 := make(chan int)
	c2 := make(chan int)
		go func() {
		for _, value := range massiv {
			c1 <- value
		}
	}()

	go func() {
		for value := range c1 {
			c2 <- value * 2
		}
	}()
    
	go func() {
		for value := range c2 {
			fmt.Println(value)
		}
	}()
	fmt.Println(massiv)
 }

