package main

import "fmt"

func main() {
	fmt.Println("000tester: ", dpFibo(6))
}

func dpFibo(n int) []int {
	dp := make([]int, n, n)

	for i := 0; i < n; i++ {
		if i == 0 || i == 1 {
			fmt.Println("111tester: ", i)
			dp[i] = 1
		} else {
			dp[i] = dp[i-1] + dp[i-2]
		}
	}

	return dp
}
