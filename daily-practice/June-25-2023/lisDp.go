package main

import "fmt"

func main() {
	arr := []int{10, 9, 2, 5, 3, 7, 101, 18}
	fmt.Println("000tester: ", lisDp(arr))
}

func lisDp(nums []int) []int {
	dp := make([][]int, len(nums))

	for numsIndex, numsEle := range nums {
		dp[numsIndex] = []int{numsEle}
	}

	for i := 1; i < len(nums); i++ {
		tempMaxArr := dp[i]
		for k := i - 1; k >= 0; k-- {
			if tempMaxArr[0] > nums[k] {
				tempMaxArr = append([]int{nums[k]}, tempMaxArr...)
			}
		}

		if len(tempMaxArr) > len(dp[i]) {
			dp[i] = tempMaxArr
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
