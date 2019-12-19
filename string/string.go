package string

// 最长子串
func lengthOfLongestSubstring(s string) int {
	n, ans := len(s), 0
	i := 0
	m := make(map[byte]int)
	for j := 0; j < n; j++ {
		k := s[j]
		if _, ok := m[k]; ok {
			if m[k] > i {
				i = m[k]
			}
		}
		if j-i+1 > ans {
			ans = j - i + 1
		}
		m[k] = j + 1
	}
	return ans
}

// 有效的括号
func isValid(s string) bool {
	if len(s) == 0 {
		return true
	}
	if len(s)%2 != 0 {
		return false
	}
	var arr []byte
	var m = map[byte]byte{')': '(', ']': '[', '}': '{'}
	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			arr = append(arr, s[i])
			continue
		}
		if len(arr) == 0 || m[s[i]] != arr[len(arr)-1] {
			return false
		}
		arr = arr[:len(arr)-1]
	}
	if len(arr) == 0 {
		return true
	}
	return false
}

// 最长有效
func longestValidParentheses(s string) int {
	maxans := 0
	var stack []int
	stack = append(stack, -1)
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, i)
			continue
		}
		stack = stack[:len(stack)-1]
		if len(stack) == 0 {
			stack = append(stack, i)
			continue
		}
		if i-stack[len(stack)-1] > maxans {
			maxans = i - stack[len(stack)-1]
		}
	}
	return maxans
}
