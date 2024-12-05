package main

import (
	"bufio"
	"fmt"
	"os"
)

var directions = [][]int{
	{-1, 0},  // Up
	{1, 0},   // Down
	{0, -1},  // Left
	{0, 1},   // Right
	{-1, -1}, // Diagonal Up-Left
	{-1, 1},  // Diagonal Up-Right
	{1, -1},  // Diagonal Down-Left
	{1, 1},   // Diagonal Down-Right
}

type Puzzle struct {
	Grid  [][]string
	Words []string
}

func findWords(puzzle Puzzle) ([]string, int) {
	var foundWords []string
	var foundCount int
	for _, word := range puzzle.Words {
		for i := 0; i < len(puzzle.Grid); i++ {
			for j := 0; j < len(puzzle.Grid[i]); j++ {
				if found, count := searchFrom(puzzle.Grid, i, j, word); found {
					foundWords = append(foundWords, word)
					foundCount = foundCount + count
				}
			}
		}
	}
	return foundWords, foundCount
}

func searchFrom(grid [][]string, row, col int, word string) (bool, int) {
	// Check in all 8 directions
	count := 0
	for _, dir := range directions {
		if found := searchDirection(grid, row, col, word, dir); found {
			if foundX := searchDirection(grid, row+2, col, word, []int{-1, 1}); foundX {
				count++
			}
			if foundY := searchDirection(grid, row, col+2, word, []int{-1, -1}); foundY {
				count++
			}
		}

		if count > 0 {
			return true, count
		}
	}
	return false, 0
}

func searchDirection(grid [][]string, row, col int, word string, dir []int) bool {
	// Check if word can fit in the grid in the given direction
	for i := 0; i < len(word); i++ {
		newRow := row + i*dir[0]
		newCol := col + i*dir[1]

		if newRow < 0 || newRow >= len(grid) || newCol < 0 || newCol >= len(grid[0]) {
			return false
		}
	}

	// Check if the word matches
	for i := 0; i < len(word); i++ {
		newRow := row + i*dir[0]
		newCol := col + i*dir[1]

		if grid[newRow][newCol] != string(word[i]) {
			return false
		}
	}

	return true
}

func readGridFromFile(filename string) ([][]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var grid [][]string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		var row []string
		for _, c := range line {
			row = append(row, string(c))
		}
		grid = append(grid, row)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return grid, nil
}

func main() {
	// Open the file
	grid, err := readGridFromFile("input.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for _, row := range grid {
		fmt.Println(row)
	}

	puzzle := Puzzle{
		Grid:  grid,
		Words: []string{"MAS"},
	}

	foundWords, foundCount := findWords(puzzle)
	for _, word := range foundWords {
		fmt.Println(word)
	}
	fmt.Printf("Found %d words\n", foundCount)

}
