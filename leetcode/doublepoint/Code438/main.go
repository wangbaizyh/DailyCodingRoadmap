package Code438

/**
* @author G.E.E.K.
* @create 2022-04-16 18:40
 */

func main() {

}

func findAnagrams(s string, p string) []int {
	record := make([]int, 26)
	for _, c := range p {
		record[c-'a']++
	}

	result := make([]int, 0)
	i, j := 0, 0

	for j < len(s) {
		if record[s[j]-'a'] > 0 {
			record[s[j]-'a']--
			j++

			if j-i == len(p) {
				result = append(result, i)
			}
		} else {
			record[s[i]-'a']++
			i++
		}
	}

	return result
}
