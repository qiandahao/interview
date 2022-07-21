// 392. Is Subsequence
// https://leetcode.com/problems/is-subsequence/

func isSubsequence(s string, t string) bool {
    sdx, tdx := 0, 0
	for sdx < len(s) && tdx < len(t) {
		if s[sdx] == t[tdx] {
			sdx++
			tdx++
		} else {
			tdx++
		}
	}
	if sdx < len(s) {
		return false
	} 
	return true
}