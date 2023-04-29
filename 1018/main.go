package main

import "fmt"

func main() {
	var m, n int
	var arr [][]bool = [][]bool{}
	var inputBuffer rune
	fmt.Scanf("%d %d\n", &m, &n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			fmt.Scanf("%c", inputBuffer)
			if inputBuffer == 'W' {
				arr[i][j] = true
			} else {
				arr[i][j] = false
			}
		}
	}
	fmt.Println(arr)
}
