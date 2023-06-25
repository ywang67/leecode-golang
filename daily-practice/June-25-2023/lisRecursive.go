package main

import "fmt"

func main() {
	arr := []int{10, 9, 2, 5, 3, 7, 101, 18}
	fmt.Println("000tester: ", findLis(arr))
}

func findLis(arr []int) []int {
	if len(arr) == 0 {
		return arr[0:]
	}
	dp := make([][]int, len(arr))

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

	for i := 1; i < len(arr); i++ {
		subsequence := lisHelper(i)
		if len(subsequence) > len(dp[i]) {
			dp[i] = subsequence
		}
	}

	res := []int{}

	for dpIndex, dpEle := range dp {
		fmt.Println("111tester: ", dpIndex, dpEle)
		if len(dpEle) > len(res) {
			res = dpEle
		}
	}

	return res
}
