package main

import "fmt"

func main() {
	arr := []int{10, 9, 2, 5, 3, 7, 101, 18}

	fmt.Println("res: ", lis(arr))
}

func lis(arr []int) []int {
	arrLen := len(arr)
	dp := make([][]int, arrLen)

	for arrIndex, arrVal := range arr {
		dp[arrIndex] = []int{arrVal}
	}

	var lisHelper func(index int) []int
	lisHelper = func(index int) []int {
		tempMaxArr := []int{}
		if index > 0 {
			for j := index - 1; j >= 0; j-- {
				if arr[index] > arr[j] {
					subsequence := lisHelper(j)
					if len(subsequence) > len(tempMaxArr) {
						tempMaxArr = subsequence
					}
				}
			}
		}

		return append(tempMaxArr, arr[index])
	}

	for i := 1; i < arrLen; i++ {
		subsequence := lisHelper(i)
		if len(subsequence) > len(dp[i]) {
			dp[i] = subsequence
		}
	}

	res := []int{}
	for dpIndex, dpEle := range dp {
		fmt.Println("000tester: ", dpIndex, dpEle)
		if len(res) < len(dpEle) {
			res = dpEle
		}
	}

	return res
}
