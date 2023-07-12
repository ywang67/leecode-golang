package main

import "fmt"

func main() {
	nums := []int{10, 9, 2, 5, 3, 7, 101, 18}

	fmt.Println("000tester: ", lisDp(nums))
}

func lisDp(nums []int) []int {
	numsLen := len(nums)
	if numsLen == 0 {
		return nums
	}

	dp := make([][]int, numsLen)
	for numsIndex, numsEle := range nums {
		dp[numsIndex] = []int{numsEle}
	}

	for i := 1; i < numsLen; i++ {
		tempArrMax := dp[i]
		for j := i - 1; j >= 0; j-- {
			if nums[j] < tempArrMax[0] {
				tempArrMax = append([]int{nums[j]}, tempArrMax...)
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
