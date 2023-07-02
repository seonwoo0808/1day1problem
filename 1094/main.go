package main

import (
	"fmt"
	"math"
)

func main() {
	var x, cnt int = 0, 0
	var stick float64
	fmt.Scan(&x)
	for i := 6.0; i >= 0; i-- {
		stick = math.Pow(2, i)
		if x-int(stick) >= 0 {
			x -= int(stick)
			cnt++
		}
	}
	fmt.Println(cnt)
}
