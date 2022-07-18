// 84. Largest Rectangle in Histogram
// 2,1,5,6,2,3
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
// 2, 2 1 -- 1

int n = heights.size(), ans = 0;
        heights.push_back(-1);
        // 为了算法书写方便，在数组末尾添加高度 -1
        // 这会使得栈中所有数字在最后出栈。
        stack<int> st;
        for (int i = 0; i <= n; i++) {
            while (!st.empty() && heights[i] < heights[st.top()]) {
                int cur = st.top();
                st.pop();
                if (st.empty())
                    ans = max(ans, heights[cur] * i);
                else
                    ans = max(ans, heights[cur] 
                            * (i - st.top() - 1));
            }
            st.push(i);
        }
        return ans;