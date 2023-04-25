package main

import "fmt"

func main() {
	var cnt int
	var result, inputBuffer string
	fmt.Scanln(&cnt)
	fmt.Scanln(&inputBuffer)
	result = inputBuffer
	for i := 1; i < cnt; i++ {
		fmt.Scanln(&inputBuffer)
		for idx, char := range inputBuffer {
			if result[idx] != byte(char) {
				result = result[:idx] + "?" + result[idx+1:]
			}
		}
	}
	fmt.Println(result)
}
