func minimizedMaximum(n int, quantities []int) int {
    sum := 0
    max := 0
    for i:=0; i<len(quantities); i++ {
        max = maxInt(max, quantities[i])
        sum += quantities[i]
    }
    //3, 11
    left := (sum + n - 1) / n
    right := max
    
    mid := left + (right - left) /2 
    
    for left < right {
        count := checkAvailable(mid, quantities)
        if count > n {
            left = mid + 1
        } else {
            right = mid
        }
        mid = left + (right - left) /2 
    }
    return left
}

func maxInt(i, j int) int {
    if i > j {
        return i
    }
    return j
}

func checkAvailable(num int, quan []int) int {
    count := 0
    for i:=0; i<len(quan); i++ {
        count += (quan[i] + num - 1) / num
    }
    return count
}