package word

// 判断一个字符串是否从前往后和从后往前读是一致的
func IsPalindrome(s string) bool {
	for i := range s {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}
