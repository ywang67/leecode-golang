package main

import "fmt"

func main() {
	arr := []int{10, 9, 2, 5, 3, 7, 101, 18}
	fmt.Println("000tester: ", listRecursive(arr))
}

func listRecursive(arr []int) []int {
	arrLen := len(arr)
	if arrLen <= 1 {
		return arr
	}
	dp := make([][]int, len(arr))
	for arrIndex, arrEle := range arr {
		dp[arrIndex] = []int{arrEle}
	}

	var lisHelper func(indexArg int) []int
	lisHelper = func(indexArg int) []int {
		tempLis := []int{}
		for k := indexArg - 1; k >= 0; k-- {
			if arr[indexArg] > arr[k] {
				subsequence := lisHelper(k)
				if len(subsequence) > len(tempLis) {
					tempLis = subsequence
				}
			}
		}

		return append(tempLis, arr[indexArg])
	}

	for i := 1; i < arrLen; i++ {
		subsequence := lisHelper(i)
		if len(subsequence) > len(dp[i]) {
			dp[i] = subsequence
		}
	}

	res := []int{}

	for _, dpEle := range dp {
		fmt.Println("111tester: ", dpEle)
		if len(dpEle) > len(res) {
			res = dpEle
		}
	}

	return res
}
