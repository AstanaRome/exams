package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Print(FifthAndSkip("abcdefghijklmnopqrstuwxyz"))
	fmt.Print(FifthAndSkip("This is a short sentence"))
	fmt.Print(FifthAndSkip("1234"))
}

func FifthAndSkip(str string) string {
	// Удаление пробелов и табов
	cleanRunes := make([]rune, 0, len(str))
	for _, r := range str {
		if r != ' ' && r != '\t' {
			cleanRunes = append(cleanRunes, r)
		}
	}

	if len(cleanRunes) == 0 {
		return "\n"
	}
	if len(cleanRunes) < 5 {
		return "Invalid Input\n"
	}

	var builder strings.Builder
	for i, r := range cleanRunes {
		if (i+1)%6 == 0 {
			builder.WriteByte(' ')
		} else {
			builder.WriteRune(r)
		}
	}
	builder.WriteByte('\n')
	return builder.String()
}
