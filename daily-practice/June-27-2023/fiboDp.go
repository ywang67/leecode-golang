package main

import "fmt"

func main() {
	fmt.Println("000tester: ", fiboDp(5))
}

func fiboDp(n int) []int {
	dp := make([]int, n+1)
	for i := 0; i <= n; i++ {
		if i == 0 || i == 1 {
			dp[i] = 1
		} else {
			dp[i] = dp[i-2] + dp[i-1]
		}
	}

	return dp
}
