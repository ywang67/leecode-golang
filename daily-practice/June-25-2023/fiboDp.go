package main

import "fmt"

func main() {
	fibo := fiboDp(5)
	fmt.Println("000tester: ", fibo)
}

func fiboDp(n int) []int {
	dp := make([]int, n)

	for i := 0; i < n; i++ {
		if i == 0 || i == 1 {
			dp[i] = 1
		} else {
			dp[i] = dp[i-1] + dp[i-2]
		}
	}

	return dp
}
