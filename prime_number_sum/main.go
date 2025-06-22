package main

import "fmt"

func main() {
	n := 5
	res := 0
	if n < 2 {
		return
	}

	isPrime := make([]bool, n+1)
	for i := 2; i <= n; i++ {
		isPrime[i] = true
	}

	for i := 2; i*i <= n; i++ {
		for j := i * i; j <= n; j = j + i {
			isPrime[j] = false
		}
	}

	for i := 2; i <= n; i++ {
		if isPrime[i] {
			res += i
		}
	}

	fmt.Println("000tester: ", res)
}
