// 85_Maximal_Rectangle
// https://leetcode.com/problems/maximal-rectangle/

func maximalRectangle(matrix [][]byte) int {
    rows := len(matrix)
    cols := len(matrix[0])
    
	heigths := make([]int, cols)
	for i:=0; i<cols; i++ {
		heights[i] = 0
	}
	heights = append(heights, -1)
    for i:=0; i<rows; i++ {
        for j:=0; j<cols; j++ {
			if matrix[i][j] == 0 {
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
					res = max(res, heights[cur]*(i - cur - 1))
				}
			}
			val := heights[i]
			if len(stack) == 0 {
				stack = append(stack, heigths[i])
			}
			
			if heights[stack[len(stack)-1]] > val {
				

				res := max(res, heights[cur] * (i-cur))
			}
		}

    }
}