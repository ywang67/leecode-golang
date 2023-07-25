package main

import "fmt"

func main() {
	arr := []int{9, 10, 2, 5, 3, 7, 101, 18}

	fmt.Println("00tester: ", lisDp(arr))
}

func lisDp(arr []int) []int {
	arrLen := len(arr)

	dp := make([][]int, arrLen)
	for arrIndex, arrEle := range arr {
		dp[arrIndex] = []int{arrEle}
	}

	for i := 1; i < arrLen; i++ {
		tempArrMax := dp[i]
		for j := i - 1; j >= 0; j-- {
			if arr[j] < tempArrMax[0] {
				tempArrMax = append([]int{arr[j]}, tempArrMax...)
			}
		}

		if len(tempArrMax) > len(dp[i]) {
			dp[i] = tempArrMax
		}
	}

	res := []int{}

	for _, dpEle := range dp {
		if len(dpEle) > len(res) {
			res = dpEle
		}
	}

	return res
}
