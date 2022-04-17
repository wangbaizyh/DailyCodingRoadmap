package Code349

/**
* @author G.E.E.K.
* @create 2022-04-17 12:31
 */

func main() {

}

func intersection(nums1 []int, nums2 []int) []int {
	set := make(map[int]bool)
	for _, v := range nums1 {
		set[v] = true
	}

	result := make(map[int]bool)
	for _, v := range nums2 {
		if set[v] == true {
			result[v] = true
		}
	}

	ans := make([]int, 0)
	for k := range result {
		ans = append(ans, k)
	}

	return ans
}

func intersection01(nums1 []int, nums2 []int) []int {
	set := make(map[int]bool)
	for _, v := range nums1 {
		set[v] = true
	}

	ans := make([]int, 0)
	for _, v := range nums2 {
		if set[v] == true {
			ans = append(ans, v)
			set[v] = false
		}
	}

	return ans
}

func intersection02(nums1 []int, nums2 []int) []int {
	set := make(map[int]struct{})
	for _, v := range nums1 {
		if _, ok := set[v]; !ok {
			set[v] = struct{}{}
		}
	}

	ans := make([]int, 0)
	for _, v := range nums2 {
		if _, ok := set[v]; ok {
			ans = append(ans, v)
			delete(set, v)
		}
	}

	return ans
}
