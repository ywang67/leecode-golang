package main

import "fmt"

func maxSubArray(nums []int) (int, []int) {
	n := len(nums)
	if n == 0 {
		return 0, nil
	}

	dp := make([]int, n) // dp[i] 表示以 nums[i] 结尾的最大子数组和
	dp[0] = nums[0]
	maxSum := dp[0]
	maxEnd := 0

	for i := 1; i < n; i++ {
		if dp[i-1] > 0 {
			dp[i] = dp[i-1] + nums[i]
		} else {
			dp[i] = nums[i]
		}

		if dp[i] > maxSum {
			maxSum = dp[i]
			maxEnd = i
		}
	}

	// 构造最大子数组
	subarray := make([]int, 0, maxEnd+1)
	currSum := maxSum

	for i := maxEnd; i >= 0; i-- {
		subarray = append([]int{nums[i]}, subarray...)
		currSum -= nums[i]
		if currSum == 0 {
			break
		}
	}

	return maxSum, subarray
}

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	maxSum, subarray := maxSubArray(nums)
	fmt.Println("最大和:", maxSum)
	fmt.Println("最大子数组:", subarray)
}
