package main

import "fmt"

func getLastComputer(a int, b int) (result int) {
	a = a % 10
	result = a
	if a == 0 {
		return 10
	}
	for i := 1; i < b; i++ {
		result = (result * a) % 10
	}
	return result
}

func main() {
	var cnt, a, b int
	fmt.Scanln(&cnt)
	for i := 0; i < cnt; i++ {
		fmt.Scanf("%d %d\n", &a, &b)
		fmt.Println(getLastComputer(a, b))
	}

}
