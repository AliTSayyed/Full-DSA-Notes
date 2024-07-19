// binary search algorithim
// run in terminal using go run main.go
package main

import "fmt"

func binarySearch(nums []int, target int) (index int) {
	high := len(nums) - 1
	low := 0

	for low <= high {
		mid := (high + low) / 2
		if nums[mid] < target {
			low = mid + 1
		} else if nums[mid] > target {
			high = mid - 1
		} else if nums[mid] == target {
			return mid
		}
	}
	return -1 // return -1 if not found
}

func main() {
	array := []int{1, 3, 5, 7, 8, 10}
	target := binarySearch(array, 10)
	fmt.Println("The number 10 is found at index:", target)
}
