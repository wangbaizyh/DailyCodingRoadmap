package main

import (
	"sort"
)

/**
* @author G.E.E.K.
* @create 2022-04-24 14:46
 */

func main() {

}

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	var res [][]int
	for i, _ := range nums {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		for j := i + 1; j < len(nums); j++ {

			if j-i > 1 && nums[j] == nums[j-1] {
				continue
			}

			left := j + 1
			right := len(nums) - 1
			for left < right {
				sum := nums[i] + nums[j] + nums[left] + nums[right]
				if sum > target {
					right--
				} else if sum < target {
					left++
				} else {
					res = append(res, []int{nums[i], nums[j], nums[left], nums[right]})

					for left < right && nums[right] == nums[right-1] {
						right--
					}

					for left < right && nums[left] == nums[left+1] {
						left++
					}

					right--
					left++
				}
			}
		}
	}

	return res
}
