package main

import "fmt"

func main() {
	// num := []int{2, 7, 11, 15}
	num := []int{1, 3, 4, 2}
	fmt.Print(twoSum2(num, 6))
}

// Brute Force
func twoSum(num []int, target int) []int {
	for i := 0; i < len(num); i++ {
		for j := i + 1; j < len(num); j++ {
			if num[i]+num[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

func twoSum2(nums []int, target int) []int {
	hashset := make(map[int]int, len(nums))
	for i, v := range nums {
		hashset[v] = i
	}
	for i, v := range nums {
		value := target - v
		index := hashset[value]
		if index != 0 && index != i {
			return []int{i, index}
		}
	}
	return nil
}
