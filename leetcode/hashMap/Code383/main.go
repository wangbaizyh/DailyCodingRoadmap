package main

/**
* @author G.E.E.K.
* @create 2022-04-23 20:02
 */

func main() {

}

func canConstruct(ransomNote string, magazine string) bool {
	m := make([]int, 26)
	for _, value := range magazine {
		m[value-'a']++
	}

	for _, value := range ransomNote {
		if m[value-'a'] == 0 {
			return false
		} else {
			m[value-'a']--
		}
	}

	return true
}
