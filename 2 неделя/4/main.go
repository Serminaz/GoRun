 /* Реализовать бинарный поиск встроенными методами языка.*/

package main

import (
	"fmt"
	"sort"
)

func binarySearch(numbers []int, leftBound, rightBound, numberToFind int) int {
	if rightBound >= leftBound {
		midPoint := leftBound + (rightBound-leftBound)/2

		if numbers[midPoint] == numberToFind {
			return midPoint
		}

		if numbers[midPoint] > numberToFind {
			return binarySearch(numbers, leftBound, midPoint-1, numberToFind)
		}

		return binarySearch(numbers, midPoint+1, rightBound, numberToFind)
	}

	return -1
}

func main() {
	numbers := []int{5, 3, 99, -4, 2, -1, 9, 0, 6}
	const numberToFind = 0

	sort.Ints(numbers)
	fmt.Println(numbers)

	if i := binarySearch(numbers, 0, len(numbers)-1, numberToFind); i != -1 {
		fmt.Printf("Found number %d at index %d.\n", numberToFind, i)
	}
}