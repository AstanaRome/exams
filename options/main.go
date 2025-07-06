package main

import (
	"fmt"
	"os"
)

func printHelp() {
	fmt.Println("options: abcdefghijklmnopqrstuvwxyz")
}

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		printHelp()
		return
	}

	var options uint32 = 0
	for _, arg := range args {
		if arg == "-" {
			fmt.Println("Invalid Option")
			return
		}
		if len(arg) > 1 && arg[0] == '-' {
			for i, c := range arg[1:] {
				if c == 'h' && i == 0 {
					printHelp()
					return
				}
				if c < 'a' || c > 'z' {
					fmt.Println("Invalid Option")
					return
				}
				options |= 1 << (c - 'a')
			}
		}
	}

	// Print as 4 groups of 8 bits
	for i := 3; i >= 0; i-- {
		for j := 7; j >= 0; j-- {
			bit := (options >> uint(i*8+j)) & 1
			fmt.Print(bit)
		}
		if i > 0 {
			fmt.Print(" ")
		}
	}
	fmt.Println()
}
