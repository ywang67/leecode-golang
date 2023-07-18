package main

import "fmt"

func main() {
	arr := []int{9, 10 ,2 , 5, 7, 101, 18}

	fmt.Println("000tester: ", lisDp(arr))
}

func lisDp(arr []int) []int {
	arrLen := len(arr)
	if arrLen == 0 {
		return arr
	}

	dp := make([][]int, arrLen)
	for arrIndex, arrEle := range arr {
		dp[arrIndex] = []int{arrEle}
	}

	for i := 1; i < arrLen; i++ {
		tempMaxArr := dp[i]
		for j := i - 1; j >= 0; j-- {
			if arr[j] < tempMaxArr[0] {
				tempMaxArr = append([]int{arr[j]}, tempMaxArr...)
			}
		}

		if len(tempMaxArr) > len(dp[i]) {
			dp[i] = tempMaxArr
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
