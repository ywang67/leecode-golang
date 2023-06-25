package main

import "fmt"

func main() {
	fmt.Println("000tester: ", fiboRecursive(5))
}

func fiboRecursive(n int) int {
	if n < 0 {
		return 0
	}
	if n <= 1 {
		return 1
	}

	return fiboRecursive(n-1) + fiboRecursive(n-2)
}
