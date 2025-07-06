package main

import (
	"fmt"
	"os"
)

func isBracketed(s string) bool {
	stack := []rune{}
	bracketMap := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}
	for _, c := range s {
		switch c {
		case '(', '[', '{':
			stack = append(stack, c)
		case ')', ']', '}':
			if len(stack) == 0 || stack[len(stack)-1] != bracketMap[c] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println()
		return
	}
	for _, arg := range args {
		if isBracketed(arg) {
			fmt.Println("OK")
		} else {
			fmt.Println("Error")
		}
	}
}
