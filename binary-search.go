package main

import (
	"fmt"
	"time"
)


func main() {
	array := []int { 1, 2, 9, 11, 14, 15, 23, 37, 45, 50, 51, 67, 78, 83, 99, 100 }
	numbersToFind := []int { 11, 2, 37, 51, 83, 67 }

	for index, value := range numbersToFind {
		fmt.Println("### starting test ", index)
	
		startTime := time.Now()
		found, operations := binarySearch(value, array)
		endTime := time.Now()
		delta := endTime.Sub(startTime).Nanoseconds()


		fmt.Println("start time:", startTime)
		fmt.Println("end time:", endTime)
		fmt.Println("delta (Nanoseconds):", delta)
		fmt.Println("operations:", operations)
		fmt.Println("found:", found)
		fmt.Println("")
	}
}

func binarySearch(value int, array []int) (found bool, operations int){
	low, high := 0, len(array) - 1
	var mid int

	for low <= high {
		operations += 1
		mid = (low + high) / 2

		if array[mid] == value {
			found = true
			return 
		}

		if value < array[mid] {
			high = mid - 1
		} else if value > mid {
			low = mid + 1
		}
	}

	return
}
