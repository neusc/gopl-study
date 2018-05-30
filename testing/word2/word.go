package word

import "unicode"

// 判断一个字符串是否从前往后和从后往前读是一致的
// 忽略非字母字符
func IsPalindrome(s string) bool {
	letters := make([]rune, 0, len(s)) // 预先分配足够大的数组，避免后续append多次内存重新分配
	for _, r := range s {
		if unicode.IsLetter(r) {
			letters = append(letters, unicode.ToLower(r))
		}
	}

	n := len(letters) / 2 // 避免每个比较进行两次
	for i := 0; i < n; i++ {
		if letters[i] != letters[len(letters)-1-i] {
			return false
		}
	}
	return true
}
