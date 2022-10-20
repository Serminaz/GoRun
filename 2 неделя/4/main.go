/* Реализовать бинарный поиск встроенными методами языка.*/
// leftBound   равна 0, rightBound равна  len(numbers)-1
package main

import (
	"fmt"
	"sort"
)

func binarySearch(numbers []int, leftBound, rightBound, numberToFind int) int {
	if rightBound >= leftBound { // если пустой слайс
		midPoint := leftBound + (rightBound-leftBound)/2 // ищем середину массива

		if numbers[midPoint] == numberToFind {
			return midPoint
		}

		if numbers[midPoint] > numberToFind { // если центральное число > снова делим массив на 2
			return binarySearch(numbers, leftBound, midPoint-1, numberToFind)
		}

		return binarySearch(numbers, midPoint+1, rightBound, numberToFind)
	}

	return -1
}

func main() {
	numbers := []int{5, 3, 99, -4, 2, -1, 9, 0, 6}
	const numberToFind = 2 // какое число ищем
    //fmt.Println(len(numbers))
	sort.Ints(numbers)
	fmt.Println(numbers)
    
	if i := binarySearch(numbers, 0, len(numbers)-1, numberToFind); i != -1 {
		fmt.Printf("ищем число %d, \nнаходится под индексом %d.\n", numberToFind, i)
	}
}
