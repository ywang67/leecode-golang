package main

import "fmt"

type Element struct {
	Val  int
	StartIndex int
	EndIndex int
}

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	res, resArr := maxSubArray(nums)
	fmt.Println("000tester: ", res, resArr)
}

func maxSubArray(arr []int) (int, []int) {
	arrLen := len(arr)
	if arrLen == 0 {
		return 0, arr
	}

	dp := make([]Element, arrLen)

	for arrIndex, arrEle := range arr {
		dp[arrIndex] = Element{
			Val: arrEle,
			StartIndex: arrIndex,
			EndIndex: arrIndex,
		}
	}

	for i := 0; i < arrLen; i++ {
		for j := i + 1; j < arrLen; j++ {
			tempSum := sum(arr[i:j+1])
			if tempSum > dp[i].Val {
				dp[i].EndIndex = j
				dp[i].Val = tempSum
			}
		}
	}

	maxEle := Element{}
	for _, dpEle := range dp {
		if maxEle.Val < dpEle.Val {
			maxEle = dpEle
		}
	}

	return maxEle.Val, arr[maxEle.StartIndex:maxEle.EndIndex+1]
}

func sum(arr []int) int {
	res := 0
	for _, arrEle := range arr {
		res += arrEle
	}

	return res
}