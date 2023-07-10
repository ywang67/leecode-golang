package main

import "fmt"

func main() {
	nums := []int{9, 10, 2, 5, 3, 7, 101, 18}

	fmt.Println("000tester: ", lisDp(nums))
}

func lisDp(nums []int) []int {
	numsLen := len(nums)
	if numsLen <= 1 {
		return nums
	}

	dp := make([][]int, numsLen)
	for numsIndex, numsEle := range nums {
		dp[numsIndex] = []int{numsEle}
	}

	for i := 1; i < numsLen; i++ {
		tempArrEle := dp[i]
		for j := i - 1; j >= 0; j-- {
			if nums[j] < tempArrEle[0] {
				tempArrEle = append([]int{nums[j]}, tempArrEle...)
			}
		}

		if len(tempArrEle) > len(dp[i]) {
			dp[i] = tempArrEle
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
