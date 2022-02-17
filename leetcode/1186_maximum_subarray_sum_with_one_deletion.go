func maximumSum(arr []int) int {
	leftMax := make([]int, 0)
	rightMax := make([]int, 0)
	for i := 0; i < len(arr); i++ {
		leftMax = append(leftMax, 0)
		rightMax = append(rightMax, 0)
	}

	curSum := 0
	resSum := ^int(^uint(8) >> 1)
	fmt.Println(resSum)
	// [1, -1, 0, 3]
	// [-1, -1, -1, -1]
	// [0, -5, -6, 5, 5, 0]
	for i := 0; i < len(arr); i++ {
		curSum = max(curSum+arr[i], arr[i])
		leftMax[i] = curSum
		resSum = max(resSum, curSum)
	}
	fmt.Println(resSum)
	curSum = 0
	// [2, 1, 3, 3]
	// [-1, -1, -1, -1]
	// [0, -6,-1,5,0,-5]
	for j := len(arr) - 1; j >= 0; j-- {
		curSum = max(curSum+arr[j], arr[j])
		rightMax[j] = curSum
		resSum = max(resSum, curSum)
	}
	fmt.Println(resSum)
	if len(arr) > 2 {
		for i := 1; i < len(arr)-1; i++ {
			resSum = max(resSum, leftMax[i-1]+rightMax[i+1])
		}
	}
	return resSum
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}