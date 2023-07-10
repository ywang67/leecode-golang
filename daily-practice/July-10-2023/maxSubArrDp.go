package main

import "fmt"

type Element struct {
	Value      int
	StartIndex int
	EndIndex   int
}

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	// nums := []int{4, -1, 2, 1, -5, 4}

	res, resArr := maxSubArrayDp(nums)
	fmt.Println("000tester: ", res, resArr)
}

func maxSubArrayDp(nums []int) (int, []int) {
	numsLen := len(nums)
	if numsLen <= 0 {
		return sum(nums), nums
	}

	dp := make([]Element, numsLen)
	for numsIndex, numsEle := range nums {
		dp[numsIndex] = Element{
			Value:      numsEle,
			StartIndex: numsIndex,
			EndIndex:   numsIndex,
		}
	}

	for i := 0; i < numsLen; i++ {
		max := nums[i]
		endIndex := i
		for j := i + 1; j < numsLen; j++ {
			if sum(nums[i:j+1]) > max {
				max = sum(nums[i : j+1])
				endIndex = j
			}
		}

		if dp[i].Value < max {
			dp[i] = Element{
				Value:      max,
				StartIndex: i,
				EndIndex:   endIndex,
			}
		}
	}

	res := Element{}
	for _, dpEle := range dp {
		if res.Value < dpEle.Value {
			res = dpEle
		}
	}

	return res.Value, nums[res.StartIndex : res.EndIndex+1]
}

func sum(nums []int) int {
	res := 0
	for _, ele := range nums {
		res += ele
	}

	return res
}
