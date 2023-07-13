package main

import (
	"fmt"
)

func main() {
	arr := []int{9, 10, 2, 5, 3, 7, 101, 18}

	fmt.Println("000tester: ", lisRecursive(arr))
}

func lisRecursive(arr []int) []int {
	arrLen := len(arr)
	if arrLen == 0 || arr == nil {
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
			if arr[j] < arr[index] {
				subsequcne := lisHelper(j)
				if len(subsequcne) > len(tempLis) {
					tempLis = subsequcne
				}
			}
		}

		return append(tempLis, arr[index])
	}

	for i := 0; i < arrLen; i++ {
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
