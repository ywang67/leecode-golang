package main

import (
	"fmt"
)

func main() {
	arr := []int{10, 9, 2, 5, 3, 7, 101, 18}

	fmt.Println("000tester: ", lisRecursive(arr))
}

func lisRecursive(arr []int) []int {
	dp := make([][]int, len(arr))
	arrLen := len(arr)
	for arrIndex, arrEle := range arr {
		dp[arrIndex] = []int{arrEle}
	}

	var lisHelper func(index int) []int
	lisHelper = func(index int) []int {
		tempLis := []int{}
		for j := index - 1; j >= 0; j-- {
			if arr[index] > arr[j] {
				subsequence := lisHelper(j)
				if len(subsequence) > len(tempLis) {
					tempLis = subsequence
				}
			}
		}
		return append(tempLis, arr[index])
	}

	for i := 1; i < arrLen; i++ {
		subsequence := lisHelper(i)
		if len(subsequence) > len(dp[i]) {
			dp[i] = subsequence
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
