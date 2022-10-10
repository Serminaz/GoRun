// К каким негативным последствиям может привести данный фрагмент кода,
// и как это исправить? Приведите корректный пример реализации.

//var justString string  - Глобальные переменные в большинстве случаев нарушают инкапсуляцию
package main

import (
	"fmt"
)

func someFunc() {
  v := createHugeString(1 << 10)
  justString = v[:100]
  fmt.Println(justString)
  return v[:100]
}

func main() {
  someFunc()
  
}