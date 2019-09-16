package main

import "fmt"

func main() {
	var arr1, arr2 = []int{1, 2, 3, 6, 7}, []int{2, 7, 9}
	fmt.Println(findMedianSortedArrays(arr1, arr2))
	fmt.Println(findMedianSortedArrays2(arr1, arr2))
	fmt.Println(merge(arr1, arr2))
}

// 时间复杂度 O(m+n) 空间复杂度 O(m+n) 合并两个待排序的数组，然后找出中位数
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	i, j, ilen, jlen, length, index := 0, 0, len(nums1), len(nums2), len(nums1)+len(nums2), 0
	var arr = make([]int, ilen+jlen, ilen+jlen)

	for i < ilen || j < jlen { // for index := 0; index < ilen+jlen; index++ {
		if i == ilen || (i < length && j < jlen && nums1[i] >= nums2[j]) {
			arr[index] = nums2[j]
			j++
			index++
			continue
		}

		if j == jlen || (i < length && j < jlen && nums1[i] < nums2[j]) {
			arr[index] = nums1[i]
			i++
			index++
		}
	}

	if length&1 == 0 {
		return float64(arr[length/2]+arr[length/2-1]) / 2.0
	} else {
		return float64(arr[length/2])
	}
}

// O(log (m+n))  时间复杂度 O(1)
// 用 len 表示合并后数组的长度，如果是奇数，我们需要知道第 （len+1）/2 个数就可以了，如果遍历的话需要遍历 int(len/2 ) + 1 次。如果是偶数，
//我们需要知道第 len/2和 len/2+1 个数，也是需要遍历 len/2+1 次。所以遍历的话，奇数和偶数都是 len/2+1 次。
//
// 返回中位数的话，奇数需要最后一次遍历的结果就可以了，偶数需要最后一次和上一次遍历的结果。
//所以我们用两个变量 left 和 right，right 保存当前循环的结果，在每次循环前将 right 的值赋给 left。
//这样在最后一次循环的时候，left 将得到 right 的值，也就是上一次循环的结果，接下来 right 更新为最后一次的结果。
func findMedianSortedArrays2(nums1 []int, nums2 []int) float64 {
	// 如果两个
	i, j, las, cur, ilen, jlen, length := 0, 0, -1, -1, len(nums1), len(nums2), (len(nums1)+len(nums2))/2+1

	for h := 0; h < length; h++ {
		las = cur
		if i < ilen && (j >= jlen || nums1[i] <= nums2[j]) {
			cur = nums1[i]
			i++
		} else {
			cur = nums2[j]
			j++
		}
	}
	if ((len(nums1) + len(nums2)) & 1) == 0 {
		return float64(las+cur) / 2.0
	} else {
		return float64(cur)
	}
}

// 合并两个排序好的数组
func merge(nums1 []int, nums2 []int) []int {
	i, j, ilen, jlen, length, index := 0, 0, len(nums1), len(nums2), len(nums1)+len(nums2), 0
	var arr = make([]int, ilen+jlen, ilen+jlen)

	for i < ilen || j < jlen {
		if i == ilen || (i < length && j < jlen && nums1[i] >= nums2[j]) {
			arr[index] = nums2[j]
			j++
			index++
			continue
		}

		if j == jlen || (i < length && j < jlen && nums1[i] < nums2[j]) {
			arr[index] = nums1[i]
			i++
			index++
		}
	}
	return arr
}
