package main

import (
	"fmt"
)

func main() {
	arr := []int{10, 9, 2, 5, 3, 7, 101, 18}

	fmt.Println("000tester: ", lisRecursive(arr))
}

func lisRecursive(arr []int) []int {
	arrLen := len(arr)
	if arrLen <= 1 {
		return arr
	}

	dp := make([][]int, arrLen)
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
		subseuqnce := lisHelper(i)
		if len(subseuqnce) > len(dp[i]) {
			dp[i] = subseuqnce
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
