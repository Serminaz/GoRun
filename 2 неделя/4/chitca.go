 /*Реализовать быструю сортировку массива (quicksort)
  встроенными методами языка.*/



 package main

import (
	"fmt"
)

func main() {
	ar := []int{3, 4, 1, 2, 5, 7, -1, 0}
	Quicksort(ar)
	fmt.Println(ar)
}

func Quicksort(ar []int) {
	if len(ar) <= 1 {
		return
	}

	split := partition(ar)

	Quicksort(ar[:split]) 
	Quicksort(ar[split:])
}

func partition(ar []int) int {
	pivot := ar[len(ar)/2]   //  опорная точка

	left := 0
	right := len(ar) - 1

	for {
		for ; ar[left] < pivot; left++ {
		}

		for ; ar[right] > pivot; right-- { 
		}

		if left >= right {
			return right
		}

		swap(ar, left, right)
	}
}

func swap(ar []int, i, j int) {
	tmp := ar[i]
	ar[i] = ar[j]
	ar[j] = tmp
}
