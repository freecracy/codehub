package string

import (
	"errors"
	"math"
	"net"
)

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

// ip转整数
func ip2long(ip string) (uint, error) {
	ip4 := net.ParseIP(ip).To4()
	if ip4 == nil {
		return 0, errors.New("error")
	}
	return uint(ip4[3]) | uint(ip4[2])<<8 | uint(ip4[1])<<16 | uint(ip4[0])<<24, nil
}

// 整数转ip
func long2ip(i uint) (string, error) {
	if i > math.MaxUint32 {
		return "", errors.New("error")
	}
	ip := make(net.IP, net.IPv4len)
	ip[0] = byte(i >> 24)
	ip[1] = byte(i >> 16)
	ip[2] = byte(i >> 8)
	ip[3] = byte(i)
	return ip.String(), nil
}
