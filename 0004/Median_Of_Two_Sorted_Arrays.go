package main

import "fmt"

func main() {
	var arr1, arr2 = []int{1, 3}, []int{2}
	fmt.Println(findMedianSortedArrays(arr1, arr2))
}

// O(m+n)
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) == 0 && len(nums2) == 0 {
		return 0
	}

	i, j, ilen, jlen, length, index := 0, 0, len(nums1), len(nums2), len(nums1)+len(nums2), 0
	var arr = make([]int, ilen+jlen, ilen+jlen)

	for i < ilen || j < jlen {
		if nums1[i] < nums2[j] {
			//arr = append(arr, nums1[i])
			arr[index] = nums1[i]
			i++
			index++
		} else {
			// arr = append(arr, nums2[j])
			arr[index] = nums2[j]
			j++
			index++
		}
	}

	if length&1 == 1 {
		return float64(arr[length/2] + arr[length/2+1])
	} else {
		return float64(arr[length/2])
	}
}

// O(log (m+n))
func findMedianSortedArrays2(nums1 []int, nums2 []int) float64 {
	if len(nums1) == 0 && len(nums2) == 0 {
		return 0
	}
	return 0
}
