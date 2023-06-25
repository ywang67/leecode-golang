package main

import "fmt"

func main() {
	arr := []int{10, 9, 2, 5, 3, 7, 101, 18}

	fmt.Println("res: ", lis(arr))
}

func lis(arr []int) []int {
	dp := make([][]int, len(arr))
	n := len(arr)

	for arrIndex, ele := range arr {
		dp[arrIndex] = []int{ele}
	}

	for i := 1; i < n; i++ {
		tempMaxArr := dp[i]
		for j := i - 1; j >= 0; j-- {
			if arr[i] > arr[j] {
				tempArr := append(dp[j], arr[i])
				if len(tempArr) > len(tempMaxArr) {
					tempMaxArr = tempArr
				}
			}
		}

		fmt.Println("111tester: ", i, tempMaxArr)
		dp[i] = tempMaxArr
	}

	maxSubArr := []int{}
	for dpIndex, dpEle := range dp {
		fmt.Println("000tester: ", dpIndex, dpEle)
		if len(maxSubArr) < len(dpEle) {
			maxSubArr = dpEle
		}
	}

	return maxSubArr
}
