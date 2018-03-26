package main

import "fmt"

func main()  {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	fmt.Println(lengthOfLongestSubstring("bbbbbb"))
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
	fmt.Println(lengthOfLongestSubstring("c"))
	fmt.Println(lengthOfLongestSubstring("au"))
}

func lengthOfLongestSubstring(s string) (length int) {
	strLength := len(s)
	var subString string
	for i := 0; i < strLength; i ++ {
		subString = string(s[i])
		for j := i + 1; j < strLength; j ++ {
			lenSubString := len(subString)
			for k := 0; k < lenSubString; k ++ {
				if subString[k] == s[j] { // 相同
					goto exit
				}
			}
			subString += string(s[j])

		}
		exit:
			if len(subString) > length {
				length = len(subString)
			}
	}

	return
}
