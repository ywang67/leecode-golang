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

func maxSubArrayDp(arr []int) (int, []int) {
	arrLen := len(arr)
	dp := make([]Element, arrLen)
	for arrIndex, arrEle := range arr {
		dp[arrIndex] = Element{
			Value:      arrEle,
			StartIndex: arrIndex,
			EndIndex:   arrIndex,
		}
	}

	for i := 0; i < arrLen; i++ {
		for j := i; j < arrLen; j++ {
			if dp[i].Value < sum(arr[i:j+1]) {
				dp[i].Value = sum(arr[i : j+1])
				dp[i].EndIndex = j
			}
		}
	}

	var res Element
	for _, dpEle := range dp {
		if res.Value < dpEle.Value {
			res = dpEle
		}
	}

	return res.Value, arr[res.StartIndex : res.EndIndex+1]
}

func sum(arr []int) int {
	res := 0
	for _, ele := range arr {
		res += ele
	}
	return res
}
