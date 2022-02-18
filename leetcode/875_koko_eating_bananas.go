func minEatingSpeed(piles []int, h int) int {
	max := 0
	for i := 0; i < len(piles); i++ {
		max = maxInt(max, piles[i])
	}

	left, right := 1, max
	res := 0
	for left <= right {
		mid := left + (right-left)/2
		hours := consume(piles, mid)
		if hours <= h {
			res = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return res
}

func maxInt(i, j int) int {
	if i > j {
		return i
	}
	return j
}

// [3, 3], 4 : 2
// [6, 3], 4 : 3
func consume(piles []int, es int) int {
	hours := 0
	for i := 0; i < len(piles); i++ {
		hours += (piles[i] + es - 1) / es
	}
	return hours
}