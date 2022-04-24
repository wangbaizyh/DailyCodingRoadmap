package main

import "sort"

/**
* @author G.E.E.K.
* @create 2022-04-24 12:20
 */

func main() {

}

func threeSum(nums []int) [][]int {
	result := make([][]int, 0)
	sort.Ints(nums)

	for i, num := range nums {
		if nums[i] > 0 {
			return result
		}

		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left := i + 1
		right := len(nums) - 1
		for left < right {
			sum := num + nums[left] + nums[right]
			if sum > 0 {
				right--
			} else if sum < 0 {
				left++
			} else {
				ans := []int{num, nums[left], nums[right]}
				result = append(result, ans)

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

	return result
}
