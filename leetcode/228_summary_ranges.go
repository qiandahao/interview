func summaryRanges(nums []int) []string {
    if nums == nil || len(nums) == 0 {
        return nil
    }
    if len(nums) == 1 {
        return []string{
            strconv.Itoa(nums[0]),
        }    
    }
    // 0, 1, 2, 4, 5, 7
    // i:1 , start = 0, curr=0
    // i:2 , start = 0, curr=1
    // i:3, start = 0 --> 4, curr =2
    // i:4, start =4, curr = 4 --> curr = 5
    // i:5, start= 7, curr = 7
    start, curr := nums[0], nums[0]
    res := make([]string, 0)
    for i:=1; i<len(nums); i++ {
        if curr+1 != nums[i] {
            if start != curr {
                temp := strconv.Itoa(start) + "->" + strconv.Itoa(curr)
                res = append(res, temp)
            } else {
                res = append(res, strconv.Itoa(start))
            }
            start = nums[i]
            curr = start
        } else {
            curr = nums[i]
        }
    }

    if start != curr {
        temp := strconv.Itoa(start) + "->" + strconv.Itoa(curr)
        res = append(res, temp)
    } else {
        res = append(res, strconv.Itoa(start))
    }
    
    return res
}