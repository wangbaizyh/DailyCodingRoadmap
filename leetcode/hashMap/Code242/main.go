package Code242

/**
* @author G.E.E.K.
* @create 2022-04-16 18:50
 */

func main() {

}

func isAnagram(s string, t string) bool {
	record := make([]int, 26)
	for _, c := range s {
		record[c-'a']++
	}

	for _, c := range t {
		record[c-'a']--
	}

	for _, value := range record {
		if value != 0 {
			return false
		}
	}

	return true
}
