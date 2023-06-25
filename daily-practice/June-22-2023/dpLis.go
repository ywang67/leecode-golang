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

	for i := 1; i < arrLen; i++ {
		tempMaxArr := dp[i]
		for j := i - 1; j >= 0; j-- {
			if arr[j] < tempMaxArr[0] {
				tempArr := []int{arr[j]}
				tempMaxArr = append(tempArr, tempMaxArr...)
			}
		}

		if len(tempMaxArr) > len(dp[i]) {
			dp[i] = tempMaxArr
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
