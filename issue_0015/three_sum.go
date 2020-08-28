package issue_0015

import "sort"

func threeSum(nums []int) [][]int {
	var res [][]int
	var counts = map[int]int{}

	for _, value := range nums {
		counts[value]++
	}

	var uniq []int
	for inx := range counts {
		uniq = append(uniq, inx)
	}

	sort.Ints(uniq)
	for i, size := 0, len(uniq); i < size; i++ {
		if uniq[i]*3 == 0 && counts[uniq[i]] > 2 {
			res = append(res, []int{uniq[i], uniq[i], uniq[i]})
		}

		for j := i + 1; j < size; j++ {
			if uniq[i]*2+uniq[j] == 0 && counts[uniq[i]] > 1 {
				res = append(res, []int{uniq[i], uniq[i], uniq[j]})
			}
			if uniq[j]*2+uniq[i] == 0 && counts[uniq[j]] > 1 {
				res = append(res, []int{uniq[i], uniq[j], uniq[j]})
			}

			c := 0 - uniq[i] - uniq[j]
			if c > uniq[j] && counts[c] > 0 {
				res = append(res, []int{uniq[i], uniq[j], c})
			}
		}
	}
	return res
}

func threeSum2(nums []int) [][]int {
	var res [][]int
	sort.Ints(nums)
	for i, size := 0, len(nums); i < size-2; i++ {
		if i == 0 || (i > 0 && nums[i] != nums[i-1]) {
			l, r, target := i+1, size-1, 0-nums[i]

			for l < r {
				if nums[l]+nums[r] == target {
					res = append(res, []int{nums[i], nums[l], nums[r]})
					for l < r && nums[l] == nums[l+1] {
						l++
					}
					for l < r && nums[r] == nums[r-1] {
						r--
					}
					l++
					r--
				} else if nums[l]+nums[r] < target {
					l++
				} else {
					r--
				}
			}
		}
	}
	return res
}
