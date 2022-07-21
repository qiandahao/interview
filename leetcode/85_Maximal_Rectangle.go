// 85_Maximal_Rectangle
// https://leetcode.com/problems/maximal-rectangle/


func maximalRectangle(matrix [][]byte) int {
    rows := len(matrix)
    cols := len(matrix[0])
    
	heights := make([]int, cols)
	for i:=0; i<cols; i++ {
		heights[i] = 0
	}

	heights = append(heights, -1)
    res := 0
    for i:=0; i<rows; i++ {
        for j:=0; j<cols; j++ {
			if matrix[i][j] == '0' {
				heights[j] = 0
			} else {
				heights[j]++
			}
		}
        
		stack := make([]int, 0)
		for i:=0; i<len(heights); i++{
			for len(stack) != 0 && heights[i] < heights[stack[len(stack)-1]] {
				cur := stack[len(stack)-1]
				stack = stack[:len(stack)-1]

				if len(stack) == 0 {
					res = max(res, heights[cur]*i)
				} else {
					res = max(res, heights[cur]*(i - stack[len(stack)-1] - 1))
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