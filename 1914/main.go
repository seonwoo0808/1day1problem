package main

import (
	"fmt"
	"math/big"
)

func hanoi(n, start, mid, end int64) {
	if n == 1 {
		fmt.Println(start, end)
		return
	}
	hanoi(n-1, start, end, mid)
	fmt.Println(start, end)
	hanoi(n-1, mid, start, end)
}

func main() {
	var n int64
	fmt.Scan(&n)
	nn := big.NewInt(n)
	two := big.NewInt(2)

	exp := new(big.Int).Exp(two, nn, nil)
	result := new(big.Int).Sub(exp, big.NewInt(1))
	fmt.Println(result.String())

	if n <= 20 {
		hanoi(n, 1, 2, 3)
	}
}
