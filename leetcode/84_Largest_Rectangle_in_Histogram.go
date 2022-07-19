// 84_Largest_Rectangle_in_Histogram
// https://leetcode.com/problems/largest-rectangle-in-histogram/

func largestRectangleArea(heights []int) int {
	stack := make([]int, 0)
	res := 0
	heights = append(heights, -1)
	for i:=0; i<len(heights);i++ {
			for len(stack) != 0 && heights[i] < height[stack[len(stack)-1]] {
				cur := stack[len(stack) - 1]
				stack = stack[:len(stack)-1]
				if len(stack) == 0 {
					res = max(res, heights[cur] * i)
				} else {
					res = max(ans, heights[cur] * (i - height[stack[len(stack)-1] - 1))
				}
			}
			stack = append(stack, i)
		}
	}
	return res
}


func max(a int, b int) int {
	if a > b {
			return a
	} else {
			return b
	}
}