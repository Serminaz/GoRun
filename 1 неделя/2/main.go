/*  Разработать программу, которая будет последовательно 
отправлять значения в канал, а с другой стороны канала — читать.
По истечению N секунд программа должна завершаться.*/

package main

import (
  "fmt"
  "time"
)

func otpravitel(n int, c chan int) {
  for i := 0; i < n; i++ {
    c <- i // отправляет в канал 
    i = +i
  }
  time.Sleep( 2 * time.Second) 
  close(c) // закрыть канал может только отправитель
}

func main() {
    start := time.Now()
  c := make(chan int)  
  go otpravitel(15, c)
  for i := range c { // читаем из канала, пока он открыт
    fmt.Println(i)
  }
  fmt.Println( time.Since(start))
}