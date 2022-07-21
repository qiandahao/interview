// 1422 https://leetcode.com/problems/maximum-score-after-splitting-a-string
package main

func maxScore(s string) int {
	left, right := make([]int, len(s)), make([]int, len(s))
	zeros, ones := 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] == '0' {
			zeros++
		}
		if s[len(s)-1-i] == '1' {
			ones++
		}
		left[i] = zeros
		right[len(s)-1-i] = ones
	}
	res := 0
	for i := 0; i < len(s)-1; i++ {
		temp := left[i] + right[i+1]
		if temp > res {
			res = temp
		}
	}
	return res
}
