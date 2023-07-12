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
	if arrLen == 0 {
		return arr
	}

	dp := make([][]int, arrLen)
	for arrIndex, arrEle := range arr {
		dp[arrIndex] = []int{arrEle}
	}

	var lisHelper func(index int) []int
	lisHelper = func(index int) []int {
		tempLis := []int{}
		for i := index - 1; i >= 0; i-- {
			if arr[i] < arr[index] {
				subsequence := lisHelper(i)
				if len(subsequence) > len(tempLis) {
					tempLis = subsequence
				}
			}
		}

		return append(tempLis, arr[index])
	}

	for k := 1; k < arrLen; k++ {
		subsequence := lisHelper(k)
		if len(subsequence) > len(dp[k]) {
			dp[k] = subsequence
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
