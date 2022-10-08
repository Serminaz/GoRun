/*Реализовать пересечение двух неупорядоченных множеств.*/
//https://tproger.ru/articles/awesome-bits/  ( побитывые операции)
package main

import (
	"fmt"
	//"time"
	//"image"
)
// Вариант 1
// func main() {
// 	a := []int{3, 5, 1, 4, 2}
// 	b := []int{7, 6, 9, 3, 8, 5}
// 	fmt.Println(a)
// 	fmt.Println(b)

// 	m := make(map[int]uint8)
// 	for _, k := range a {
// 		m[k] |= (1 << 0)
// 	}
// 	fmt.Println("after", m)
// 	for _, k := range b {
// 		m[k] |= (1 << 1)
// 	}
// 	fmt.Println("after", m)

// 	var inAAndB []int
// 	for k, v := range m {
// 		// сравнивает мапу(сумму массивов ) и два массива
// 		// и при совпадении с мапой вручает true
// 		a := v&(1<<0) != 0 //Чтобы записать 1 в бит n(0):
// 		b := v&(1<<1) != 0
// 		switch {
// 		case a && b:
// 			inAAndB = append(inAAndB, k)
// 		}
// 	}
// 	fmt.Println(inAAndB)
//   }

//Вариант 2
func main() {
	first := []int{7, 6, 9, 3, 8, 5}
	second := []int{1, 2, 3, 4, 5}

	fmt.Println(intersection(first, second))
}

func intersection(first, second []int) []int{
	out := []int{}
	bucket := map[int]bool{}

	for _, i := range first {
		for _, j := range second {
			if i == j && !bucket[i] {
				out = append(out, i)
				bucket[i] = true
			}
		}
	}

	return out
}

