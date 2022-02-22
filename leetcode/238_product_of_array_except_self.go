package main

import (
	"fmt"
	"strconv"
)
// Given an integer array nums, return an array answer such that answer[i] is equal to the product of all the elements of nums except nums[i].

// The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.

// You must write an algorithm that runs in O(n) time and without using the division operation.

 

// Example 1:

// Input: nums = [1,2,3,4]
// Output: [24,12,8,6]
// Example 2:

// Input: nums = [-1,1,0,-3,3]
// Output: [0,0,9,0,0]
 

// Constraints:

// 2 <= nums.length <= 105
// -30 <= nums[i] <= 30
// The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.
 

// Follow up: Can you solve the problem in O(1) extra space complexity? (The output array does not count as extra space for space complexity analysis.)

func main() {
	fmt.Println(strconv.Itoa(0))
}

func productExceptSelf(nums []int) []int {
    prefix := make([]int, len(nums))
	prefix[0] = 1
	for i := 1; i < len(nums); i++ {
		prefix[i] = prefix[i-1] * nums[i-1]
	}
	// [24,12,8,6]
	// 1,2,3,4
	// 1,1,2,6
	//   ,12,4,1  
	n := len(nums) - 1
	temp := 1
	for i := n; i >= 0; i-- {
		if i == 0 {
			prefix[i] = temp
		} else {
			prefix[i] = temp * prefix[i] 
			temp = temp * nums[i]
		}
	}
	return prefix
}