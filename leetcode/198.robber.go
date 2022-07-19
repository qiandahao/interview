package main
// 100， 1， 1， 100， 1

import "fmt"

func main() {
	fmt.Println(rob([]int{
		1,2,3,1,
	}))
}
func rob(nums []int) int {
    n := len(nums)+2
    dp := make([]int, n)
	dp[0], dp[1] = 0, 0
    
    for i:=2;i<n; i++ {
        dp[i] = maxInt(dp[i-2] + nums[i-2], dp[i-1])
    }

    return dp[n-1]
}

func maxInt(i, j int) int {
	if i > j {
		return i
	} 
	return j
}