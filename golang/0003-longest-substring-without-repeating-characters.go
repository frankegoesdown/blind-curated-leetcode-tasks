package main

import "fmt"

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	fmt.Println(lengthOfLongestSubstring2("abcabcbb"))
}

func lengthOfLongestSubstring(s string) int {
	fmt.Println("=========================")
	var subs [128]bool
	i, j, result := 0, 0, 0
	max := len(s)
	for i < max && j < max {
		fmt.Println(i)
		fmt.Println(j)
		if !subs[s[j]] {
			subs[int8(s[j])] = true
			j++
			if result > j-i {
				result = result
			} else {
				result = j - i
			}
		} else {
			subs[int8(s[i])] = false
			i++
		}
		fmt.Println(subs)
	}
	fmt.Println(result)
	fmt.Println("=========================")

	return result
}

func lengthOfLongestSubstring2(s string) int {
	fmt.Println("=========================")
	maxr := 0
	m := [128]int{}
	for i := 0; i < 128; i++ {
		m[i] = -1
	}
	lastIndex := 0
	length := 0
	for k := 0; k < len(s); k++ {
		fmt.Println(k)
		fmt.Println(m[s[k]])
		fmt.Println(lastIndex)
		fmt.Println(length)
		if m[s[k]] >= 0 {
			for i := lastIndex; i < m[s[k]]; i++ {
				m[s[i]] = -1
			}
			length -= m[s[k]] - lastIndex
			lastIndex = m[s[k]] + 1
		} else {
			length += 1
		}
		m[s[k]] = k
		if maxr < length {
			maxr = length
		}
		fmt.Println(m)
		fmt.Println(maxr)
	}
	fmt.Println(maxr)
	fmt.Println("=========================")
	return maxr
}
