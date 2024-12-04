package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {

	var sum int
	var matches []string

	var removeDont string

	pattern := `mul\((([1-9][0-9]*)),(([1-9][0-9]*))\)`
	// Open the file
	data, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Print the contents as a string
	//	fmt.Println(string(data))

	parts := strings.Split(string(data), "do()")
	// drop don't parts
	for _, part := range parts {
		removeDont = removeDont + (strings.Split(string(part), "don't()")[0])
	}

	r := regexp.MustCompile(pattern)

	matches = r.FindAllString(string(removeDont), -1)

	//	fmt.Printf("Matches: %v\n", matches[0])

	for _, match := range matches {

		result := r.ReplaceAllString(string(match), "$2,$3")
		split := strings.Split(result, ",")

		item1, err := strconv.Atoi(string(split[0]))
		if err != nil {
			fmt.Println("Error converting string to int:", err)
			return
		}
		item2, err := strconv.Atoi(string(split[1]))
		if err != nil {
			fmt.Println("Error converting string to int:", err)
			return
		}
		sum = sum + (item1 * item2)
		fmt.Printf("Sum: %v\n", sum)
	}

	fmt.Println("Sum of all matches:", sum)
}
