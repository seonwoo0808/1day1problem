package main

import "fmt"

func main() {
	var S string
	var streak, prev int = 0, 2
	var zero, one int = 0, 0

	fmt.Scan(&S)
	for _, v := range S {
		if v == '0' {
			if prev == 0 {
				streak++
			} else {
				zero++
				streak = 1
				prev = 0
			}
		} else {
			if prev == 1 {
				streak++
			} else {
				one++
				streak = 1
				prev = 1
			}
		}
	}
	if zero > one {
		fmt.Println(one)
	} else {
		fmt.Println(zero)
	}
}
