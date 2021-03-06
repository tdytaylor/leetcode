package main

import "fmt"

func main() {
	nums := []int{1, 2, 3}
	fmt.Println(trySubsets2(nums))
	fmt.Println("-------")
	slice := []int{0, 1, 2, 3}
	fmt.Printf("slice: %v slice addr %p \n", slice, &slice)

	ret := changeSlice(slice)
	fmt.Printf("slice: %v ret: %v slice addr %p \n", slice, &slice, ret)
}

// 第一次尝试
func trySubsets1(nums []int) [][]int {
	var result = make([][]int, 20)

	for start, mid, len := 0, len(nums)>>1, len(nums); start <= mid; start++ {
		result = append(result, nums[0:start])
		result = append(result, nums[len-start:len])
	}
	return result
}

// 第二次尝试：使用双层for循环遍历，只能找出元素都是挨着的子元组。
func trySubsets2(nums []int) [][]int {
	var result [][]int
	for i, len := 0, len(nums); i < len; i++ {
		for j := i + 1; j <= len; j++ {
			result = append(result, nums[i:j])
		}
	}
	return result
}

// 使用回溯算法求解所有解
func subsets(nums []int) [][]int {
	//backtrack := func(index int, nums []int, result [][]int, cache []int) {
	//
	//}
	//
	//var cache []int

	for i, len := 0, len(nums); i < len; i++ {

	}

	return nil
}

func changeSlice(nums []int) []int {
	fmt.Printf("func: %p \n", &nums)
	nums[1] = 11
	return nums
}
