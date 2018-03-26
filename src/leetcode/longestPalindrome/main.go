package main

import "fmt"

// 最长回文字符串
func longestPalindrome(s string) (res string)  {
	lenString := len(s)
	var maxRes string = string(s[0])
	for i := 0; i < lenString - 1; i ++ {
		if ! (s[i] == s[i + 1] || (i>0 && s[i-1] == s[i+1])){
			continue
		}

		if s[i] == s[i+1] {
			res = ""
			for j := i; j >= 0; j -- {
				var left, right int
				left = j
				right = i + (i - j) + 1
				if right >= lenString || s[left] != s[right]{
					break
				}

				res = string(s[left]) + res + string(s[right])
			}

			if len(res) > len(maxRes) {
				maxRes = res
			}
		}

		if i > 0 && s[i-1] == s[i+1]{
			res = string(s[i])
			for j := i; j >= 0; j -- {
				var left, right int
				left = j - 1
				if left < 0 {
					break
				}

				right = i + (i - j) + 1
				if right >= lenString || s[left] != s[right]{
					break
				}

				res = string(s[left]) + res + string(s[right])
			}

			if len(res) > len(maxRes) {
				maxRes = res
			}
		}

	}

	res = maxRes

	return
}

func main() {
	//fmt.Println(longestPalindrome("babad"))
	//fmt.Println(longestPalindrome("cbbd"))
	fmt.Println(longestPalindrome("aaaa"))
	fmt.Println(longestPalindrome("aaa"))
	fmt.Println(longestPalindrome("abcda"))
}