package main

import (
	"fmt"
	"os"
)

func contains(s, substr string) bool {
	ls, lsub := len(s), len(substr)
	if lsub == 0 {
		return true
	}
	for i := 0; i <= ls-lsub; i++ {
		if s[i:i+lsub] == substr {
			return true
		}
	}
	return false
}

func trimNonLetters(s string) string {
	start, end := 0, len(s)-1
	for start <= end && !isLetter(s[start]) {
		start++
	}
	for end >= start && !isLetter(s[end]) {
		end--
	}
	if start > end {
		return ""
	}
	return s[start : end+1]
}

func isLetter(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z')
}

func main() {
	if len(os.Args) != 3 {
		return
	}
	pattern := os.Args[1]
	input := os.Args[2]
	if len(input) == 0 {
		return
	}
	if len(pattern) < 3 {
		return
	}
	open, close := pattern[0], pattern[len(pattern)-1]
	if !((open == '[' && close == ']') || (open == '(' && close == ')')) {
		return
	}
	inside := pattern[1 : len(pattern)-1]
	if len(inside) == 0 {
		return
	}

	// Split alternatives by |
	patterns := []string{}
	curr := ""
	for _, c := range inside {
		if c == '|' {
			patterns = append(patterns, curr)
			curr = ""
		} else {
			curr += string(c)
		}
	}
	patterns = append(patterns, curr)
	for _, p := range patterns {
		if p == "" {
			return
		}
	}

	// Split input into words (by spaces)
	words := []string{}
	word := ""
	for _, c := range input {
		if c == ' ' {
			if word != "" {
				words = append(words, word)
				word = ""
			}
		} else {
			word += string(c)
		}
	}
	if word != "" {
		words = append(words, word)
	}

	count := 0
	for _, w := range words {
		trimmed := trimNonLetters(w)
		if trimmed == "" {
			continue
		}
		for _, pat := range patterns {
			if contains(trimmed, pat) {
				count++
				fmt.Printf("%d: %s\n", count, trimmed)
			}
		}
	}
}
