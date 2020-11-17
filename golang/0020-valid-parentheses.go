package main

import "fmt"

func main() {
	fmt.Println(isValid("()[]{}"))
}
func isValid(s string) bool {
	size := len(s)
	stack := make([]byte, size)
	top := 0

	for i := 0; i < size; i++ {
		// s := "{{[]}}()"
		// rune int32
		// brackets, a-zA-Z, digits := 1 byte
		// russian letters := 2 byte
		// chineese letters := 4 byte
		c := s[i] // rune
		// [41, 52, 0, 0, 0, 0,]
		switch c {
		case '(':
			stack[top] = c + 1
			top++
		case '[', '{':
			stack[top] = c + 2
			top++
		case ')', ']', '}':
			if top > 0 && stack[top-1] == c {
				top--
			} else {
				return false
			}
		}
	}
	return top == 0
}
