package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Define the file path
	filePath := "path/to/your/file.txt"

	// Define the minimum word length
	var x int
	fmt.Print("Enter the minimum word length: ")
	_, err := fmt.Scanf("%d", &x)
	if err != nil {
		fmt.Println("Invalid input. Please enter a valid integer.")
		return
	}

	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Failed to open the file: %v\n", err)
		return
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	// Iterate over each line
	for scanner.Scan() {
		// Split the line into words
		words := strings.Fields(scanner.Text())

		// Iterate over each word
		for _, word := range words {
			// Check if the word length is greater than x
			if len(word) > x {
				fmt.Println(word)
			}
		}
	}

	// Check for any errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Printf("Failed to scan the file: %v\n", err)
		return
	}
}
