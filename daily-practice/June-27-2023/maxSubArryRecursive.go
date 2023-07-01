package main

import "fmt"

func maxSubArray(nums []int) (int, []int) {
	if len(nums) == 0 {
		return 0, nil
	}

	maxSum := nums[0]
	maxEnd := 0

	maxSubArrayRecursive(nums, len(nums)-1, &maxSum, &maxEnd)

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

func maxSubArrayRecursive(nums []int, idx int, maxSum *int, maxEnd *int) int {
	if idx == 0 {
		*maxSum = nums[0]
		*maxEnd = 0
		return nums[0]
	}

	prevSum := maxSubArrayRecursive(nums, idx-1, maxSum, maxEnd)
	currSum := max(prevSum+nums[idx], nums[idx])

	if currSum > *maxSum {
		*maxSum = currSum
		*maxEnd = idx
	}

	return currSum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	maxSum, subarray := maxSubArray(nums)
	fmt.Println("最大和:", maxSum)
	fmt.Println("最大子数组:", subarray)
}
