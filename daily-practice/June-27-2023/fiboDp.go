package main

import "fmt"

func main() {
	fmt.Println("000tester: ", fiboDp(5))
}

func fiboDp(n int) []int {
	dp := make([]int, 0, n)
	for i := 0; i <= n; i++ {
		if i == 0 || i == 1 {
			dp = append(dp, 1)
		} else {
			dp = append(dp, dp[i-2]+dp[i-1])
		}
	}

	return dp
}
