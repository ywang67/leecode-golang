package main

import "fmt"

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	res, resArr := maxSubArrayDp(nums)
	fmt.Println("000tester: ", res, resArr)
}

func maxSubArrayDp(nums []int) (int, []int) {
	res := make(map[int]map[int][]int, len(nums))
	for numsIndex, numsEle := range nums {
		if res[numsIndex] == nil {
			res[numsIndex] = map[int][]int{}
		}
		for m := 0; m < len(nums)-numsIndex; m++ {
			res[numsIndex][m] = []int{numsEle}
		}
	}

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			for k := 0; k < j; k++ {
				res[i][k] = append(res[i][k], nums[j])
			}
		}
	}

	max := 0
	maxArr := []int{}
	for _, resMap := range res {
		for _, resArr := range resMap {
			tempSum := sum(resArr)
			if tempSum > max {
				max = tempSum
				maxArr = resArr
			}
		}
	}

	return max, maxArr
}

func sum(arr []int) int {
	sumRes := 0
	for _, arrEle := range arr {
		sumRes += arrEle
	}
	return sumRes
}
