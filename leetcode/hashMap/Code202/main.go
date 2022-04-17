package main

import "fmt"

/**
* @author G.E.E.K.
* @create 2022-04-17 20:22
 */

func main() {
	happy := isHappy(2)
	fmt.Println(happy)
}

func isHappy(n int) bool {
	set := make(map[int]struct{})

	getNextNumber := func(n int) int {
		res := 0
		for n > 0 {
			tmp := n % 10
			res += tmp * tmp
			n /= 10
		}

		return res
	}

	for _, ok := set[n]; n != 1 && !ok; {
		set[n] = struct{}{}
		n = getNextNumber(n)

		_, ok = set[n]
	}

	return n == 1
}
