package main

import (
	"fmt"
	"math"
	"os"
	"strings"
	"time"
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
	chars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*"
	current := make([]byte, length)
	var attemptInt int64 = 0
	startTime := time.Now()

	// Initialize the password with the first character
	for i := 0; i < length; i++ {
		current[i] = chars[0]
	}

	// Calculate total possible combinations
	totalCombinations := math.Pow(float64(len(chars)), float64(length))

	fmt.Printf("\nStarting password cracking...\n")
	fmt.Printf("Possible combinations: %.0f\n", totalCombinations)
	fmt.Printf("Character set size: %d\n", len(chars))
	fmt.Printf("Password length: %d\n\n", length)

	for {
		attemptInt++
		attempt := string(current)

		// Print progress every 10000 attempts or use \r to update same line
		if attemptInt%10000 == 0 || attempt == correctPassword {
			elapsed := time.Since(startTime)
			progress := (float64(attemptInt) / totalCombinations) * 100
			speed := float64(attemptInt) / elapsed.Seconds()

			fmt.Printf("\rProgress: %.2f%% | Attempts: %d | Speed: %.0f/s | Current: %s",
				progress, attemptInt, speed, attempt)
		}

		if attempt == correctPassword {
			elapsed := time.Since(startTime)
			fmt.Printf("\n\n🎉 Password found: %s\n", attempt)
			fmt.Printf("⏱️  Time taken: %s\n", elapsed.Round(time.Millisecond))
			fmt.Printf("🔢 Total attempts: %d\n", attemptInt)
			return
		}

		// Generate next password
		i := length - 1
		for i >= 0 {
			pos := strings.IndexByte(chars, current[i])
			if pos < len(chars)-1 {
				current[i] = chars[pos+1]
				break
			}
			current[i] = chars[0]
			i--
		}

		if i < 0 {
			elapsed := time.Since(startTime)
			fmt.Printf("\n❌ Password not found after %d attempts\n", attemptInt)
			fmt.Printf("⏱️  Time taken: %s\n", elapsed.Round(time.Millisecond))
			return
		}
	}
}
