package Code49

import "sort"

/**
* @author G.E.E.K.
* @create 2022-04-16 18:54
 */

func main() {

}

func groupAnagrams(strs []string) [][]string {
	strMap := make(map[string][]string)

	for _, str := range strs {
		s := []byte(str)
		// 对slice进行排序
		sort.Slice(s, func(i, j int) bool {
			return s[i] < s[j]
		})
		sortedStr := string(s)
		strMap[sortedStr] = append(strMap[sortedStr], str)
	}

	ans := make([][]string, 0, len(strMap))
	for _, v := range strMap {
		ans = append(ans, v)
	}

	return ans
}

func groupAnagrams02(strs []string) [][]string {
	// 使用[]int作为k []string作为v
	strMap := make(map[[26]int][]string)

	for _, str := range strs {
		// 计数字母异位词数组
		cnt := [26]int{}
		for _, v := range str {
			cnt[v-'a']++
		}

		// map append操作
		strMap[cnt] = append(strMap[cnt], str)
	}

	ans := make([][]string, 0, len(strMap))
	for _, v := range strMap {
		ans = append(ans, v)
	}

	return ans
}
