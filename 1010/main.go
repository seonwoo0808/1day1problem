package main

import (
	"fmt"
)

func nCr(n, r int) int {
	if n < r {
		return 0
	}
	if r > n/2 {
		r = n - r
	}
	res := 1
	for i := 0; i < r; i++ {
		res *= n - i
		res /= i + 1
	}
	return res
}

func main() {
	var cnt int
	fmt.Scanf("%d/n", &cnt)
	for i := 0; i < cnt; i++ {
		var m, n int
		fmt.Scanf("%d %d", &n, &m)
		fmt.Println(nCr(m, n))
	}
}
