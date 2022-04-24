package main

/**
* @author G.E.E.K.
* @create 2022-04-23 19:38
 */

func main() {

}

func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	m := make(map[int]int)
	tmp := 0
	for _, i := range nums1 {
		for _, j := range nums2 {
			tmp = i + j
			if _, ok := m[tmp]; ok {
				m[tmp]++
			} else {
				m[tmp]++
			}
		}
	}

	count := 0
	for _, i := range nums3 {
		for _, j := range nums4 {
			tmp = i + j
			if _, ok := m[-tmp]; ok {
				count += m[-tmp]
			}
		}
	}

	return count
}
