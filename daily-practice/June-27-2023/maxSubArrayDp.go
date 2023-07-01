package main

import "fmt"

type MaxArrayVal struct {
	Val  int
	List []int
}

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	res, resArr := maxSubArrayDp(nums)
	fmt.Println("000tester: ", res, resArr)
}

func maxSubArrayDp(nums []int) (int, []int) {
	dp := make([][]MaxArrayVal, len(nums))

	for numsIndex, numsEle := range nums {
		temp := MaxArrayVal{
			Val:  numsEle,
			List: []int{numsEle},
		}

		dp[numsIndex] = append(dp[numsIndex], temp)
	}

	for i := 0; i < len(nums); i++ {
		tempLis := []int{}
		for j := i + 1; j < len(nums); j++ {
			tempLis = append(tempLis, nums[j])
			tempVal := nums[i]
			for _, tempLisVal := range tempLis {
				tempVal += tempLisVal
			}
			tempMaxArrayVal := MaxArrayVal{
				Val:  tempVal,
				List: append(dp[i][0].List, tempLis...),
			}
			dp[i] = append(dp[i], tempMaxArrayVal)
		}
	}

	res := nums[0]
	resArr := []int{}
	for _, dpEle := range dp {
		fmt.Println("111tester: ", dpEle)
		for _, subDpEle := range dpEle {
			if subDpEle.Val > res {
				res = subDpEle.Val
				resArr = subDpEle.List
			}
		}
	}

	return res, resArr
}

func sum(nums []int) int {
	res := 0
	for _, numsEle := range nums {
		res += numsEle
	}
	return res
}
