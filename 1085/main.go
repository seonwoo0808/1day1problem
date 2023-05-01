package main

import "fmt"

func main() {
	var x, y, w, h, min int
	fmt.Scanf("%d %d %d %d", &x, &y, &w, &h)
	min = x
	if min > y {
		min = y
	}
	if min > w-x {
		min = w - x
	}
	if min > h-y {
		min = h - y
	}
	fmt.Println(min)
}
