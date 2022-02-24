func isAnagram(s string, t string) bool {
    ht1 := make([]int, 26)
    ht2 := make([]int, 26)
    
    if len(s) != len(t) {
        return false
    }
    
    for i:=0; i < len(s); i++ {
        idx1 := int(s[i] - 'a')
        idx2 := int(t[i] - 'a')
        ht1[idx1]++
        ht2[idx2]++
    }
    
    for i:=0; i < 26; i++ {
        if ht1[i] != ht2[i] {
            return false
        }
    }
    return true
}