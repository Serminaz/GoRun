/*Поменять местами два числа без создания временной переменной.*/
package main

import (
	"fmt"
)

func main() {
 a := 50
 b := 3
 fmt.Println( "до обмена", a, "and ", b)

 a = a+b
 b = a - b
 a = a - b
  fmt.Println( "1 вариант", a, "and ", b)
  b, a = a, b
  fmt.Println( "2 вариант", a, "and ", b)
}  