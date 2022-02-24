func isIsomorphic(s string, t string) bool {
    ht1 := make([]int, 256)
    ht2 := make([]int, 256)
    
    if len(s) != len(t) {
        return false
    }
    
    for i := 0; i < len(s); i++ {
		if ht1[int(s[i])] == 0 {
			ht1[int(s[i])] = int(t[i])
		}
		if ht2[int(t[i])] == 0 {
			ht2[int(t[i])] = int(s[i])
		}
		if ht1[int(s[i])] != int(t[i]) {
			return false
		}
		if ht2[int(t[i])] != int(s[i]) {
			return false
		}
    }
    return true
}
