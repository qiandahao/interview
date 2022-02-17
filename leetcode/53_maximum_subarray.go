// https://leetcode.com/problems/maximum-subarray/
// No. 53
const INT_MIN = ^int(^uint(0) >> 1)

func maxSubArray(nums []int) int {
	n := len(nums)

	resSum, curSum := INT_MIN, 0

	for i := 0; i < n; i++ {
		curSum = max(curSum+nums[i], nums[i])
		resSum = max(resSum, curSum)
	}
	return resSum
}

func max(i, j int) int {
	if i > j {
		return i
	} else {
		return j
	}
}