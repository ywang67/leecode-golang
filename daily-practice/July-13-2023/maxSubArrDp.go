package main

import "fmt"

type Element struct {
	Value      int
	StartIndex int
	EndIndex   int
}

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}

	resArr, res := maxSubArrayDp(nums)
	fmt.Println("000tester: ", resArr, res)
}

func maxSubArrayDp(nums []int) ([]int, int) {
	numsLen := len(nums)
	if numsLen == 0 || nums == nil {
		return nums, 0
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
		for j := i + 1; j < numsLen; j++ {
			tempSum := sum(nums[i : j+1])
			if tempSum > dp[i].Value {
				dp[i].Value = tempSum
				dp[i].EndIndex = j
			}
		}
	}

	var res Element

	for _, dpEle := range dp {
		if dpEle.Value > res.Value {
			res = dpEle
		}
	}

	return nums[res.StartIndex : res.EndIndex+1], res.Value
}

func sum(arr []int) int {
	res := 0
	for _, ele := range arr {
		res += ele
	}

	return res
}
