// 792_Number_of_Matching_Subsequences
// https://leetcode.com/problems/number-of-matching-subsequences/

func numMatchingSubseq(s string, words []string) int {
    waiting := make(map[byte][]string)
	res := 0
	for i := 0; i < len(words); i++ {
        if _, ok := waiting[words[i][0]]; !ok {
            waiting[words[i][0]] = make([]string, 0)
        }
		waiting[words[i][0]] = append(waiting[words[i][0]], words[i])
	}

	for j := 0; j < len(s); j++ {
		cur := s[j]
		if _, ok := waiting[cur]; ok {
            tmp := waiting[cur]
            waiting[cur] = nil
            for _, v := range tmp {
                if len(v) == 1 { 
					res++
                } else if len(v) > 1 {
    				substring := v[1:]
                    waiting[substring[0]] = append(waiting[substring[0]] , substring)           
                }
			}
		}
	}
	return res
}