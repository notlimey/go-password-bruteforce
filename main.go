package main

import (
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Error: Please provide a parameter")
		fmt.Println("Usage: go run main.go <parameter>")
		os.Exit(1)
	}

	parameter := os.Args[1]
	fmt.Printf("Received parameter: %s\n", parameter)

	tryPasswords(utf8.RuneCountInString(parameter), parameter)
}

func tryPasswords(length int, correctPassword string) {
	chars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	current := make([]byte, length)
	var attemptInt int64 = 0

	// Initialize the password with the first character
	for i := 0; i < length; i++ {
		current[i] = chars[0]
	}

	for {
		// Process current password attempt
		attempt := string(current)
		fmt.Printf("Trying: %s (Attempt: %d)\n", attempt, attemptInt)

		attemptInt++

		// Check if we found the password
		if attempt == correctPassword {
			fmt.Printf("Password found: %s\n", attempt)
			return // Exit the function when password is found
		}

		// Generate next password
		i := length - 1
		for i >= 0 {
			pos := strings.IndexByte(chars, current[i])
			if pos < utf8.RuneCountInString(chars)-1 {
				current[i] = chars[pos+1]
				break
			}
			current[i] = chars[0]
			i--
		}

		// If we've gone through all possibilities without finding it
		if i < 0 {
			fmt.Println("Password not found!")
			return
		}
	}
}
