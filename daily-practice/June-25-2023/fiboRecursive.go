package main

import "fmt"

func main() {
	fmt.Println("000tester: ", fiboRecursive(1))
}

func fiboRecursive(n int) int {
	if n == 0 || n == 1 {
		return 1
	}

	return fiboRecursive(n-1) + fiboRecursive(n-2)
}
