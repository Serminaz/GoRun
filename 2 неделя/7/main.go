/*Дана переменная int64. Разработать программу,
 которая устанавливает i-й бит в 1 или 0.*/
package main

import (
	"fmt"
	"math/rand"
	"time"
	"strconv"
	//"unsafe"
)

func On(n int64, i int) int64 {
	return n | 1<<(i-1)
}

func Off(n int64, i int) int64 {
	return n & ^(1 << (i - 1))
}

func main() {
    var  ff int16
	var p int64 
	var mesto int 
	for  ff != 2 { 
	fmt.Println("Введи число")
	fmt.Scanln(&p)
	numToBits := strconv.FormatInt(int64(p), 2)
	fmt.Println(numToBits)
	fmt.Println("какой бит?")
	fmt.Scanln( &mesto)
	rand.Seed(time.Now().Unix())

	on := On(p, mesto)
	off := Off(p, mesto)
	fmt.Printf(" введеное значение %d, место бита %d, измененное число %d\n",p, mesto,on )
	fmt.Printf(" введеное значение %d, место бита %d, измененное число %d\n",p, mesto,off )
	fmt.Println("желаете продолжить ? (1-y/2-n)")
    fmt.Scanln(&ff) 

	}

}

































// /*
// Дана переменная int64. Разработать программу,
// которая устанавливает i-й бит в 1 или 0.
// */
// package main

// import (
// 	"fmt"
// 	// "math/rand"
// 	//"time"
// 	//"unsafe"
// 	"strconv"
// )

// func SetOneBit(NumberInt64, nthBitCheck int64) int64 {
// 	return NumberInt64 | (1 << (nthBitCheck - 1))
// }

// // смена с 0 на 1
// func SetZeroBit(NumberInt64, nthBitCheck int64) int64 {
// 	return NumberInt64 &^ (1 << (nthBitCheck - 1))
// }
// func main() {
// 	//	 rand.Seed(time.Now().Unix())
// 	var NumberInt64 int64 = 10
// 	var nthBitCheck int64 = 1
// 	//var newOneOrZerobit int64 = 0
// 	//numToBits := strconv.FormatInt(int64(NumberInt64), 2)
// 	NewBitDigitOne := strconv.FormatInt(int64(SetOneBit(NumberInt64, nthBitCheck)), 2)
// 	NewBitDigitZero := strconv.FormatInt(int64(SetZeroBit(NumberInt64, nthBitCheck)), 2)

// 	fmt.Printf("%d бит установлен -> %s, число %d\n", nthBitCheck, NewBitDigitOne, NumberInt64)
// 	fmt.Printf("%d бит установлен -> %s, число %d\n", nthBitCheck, NewBitDigitZero, NumberInt64)

// 	// fmt.Printf("xon %b\n", *(*uint64)(unsafe.Pointer(&p)))
// 	fmt.Printf("%d бит установлен -> %s, число %d\n", nthBitCheck, NewBitDigitOne, NumberInt64|(1<<nthBitCheck-1))

// }
