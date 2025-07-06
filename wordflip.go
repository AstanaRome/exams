package main

import (
	"strings"
)

// WordFlip takes a string, reverses the order of words, trims spaces, and returns the result with a newline.
// If the input is empty or only spaces, returns a newline.
func WordFlip(s string) string {
	s = strings.TrimSpace(s)
	if s == "" {
		return "\n"
	}
	words := strings.Split(s, " ")
	var filtered []string
	for _, w := range words {
		w = strings.TrimSpace(w)
		if w != "" {
			filtered = append(filtered, w)
		}
	}
	if len(filtered) == 0 {
		return "\n"
	}
	// Reverse the filtered words
	for i, j := 0, len(filtered)-1; i < j; i, j = i+1, j-1 {
		filtered[i], filtered[j] = filtered[j], filtered[i]
	}
	return strings.Join(filtered, " ") + "\n"
}
