package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Level struct {
	Value int
}

type Report struct {
	Levels []Level
}

var report []Report

func readTextFile(filepath string) []Report {

	fmt.Println("Reading text file")
	file, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		tmpreport := Report{
			Levels: []Level{},
		}

		line := scanner.Text()
		parts := strings.Split(line, " ")

		for i := 0; i < len(parts); i++ {
			//fmt.Println(parts[i])
			value, err := strconv.Atoi(parts[i])
			if err != nil {
				panic(err)
			}
			tmpreport.Levels = append(tmpreport.Levels, Level{Value: value})
		}

		report = append(report, tmpreport)
	}

	return report

}

func checkIncreasing(slice []int) bool {
	check := false
	for i := 1; i < len(slice); i++ {
		if slice[i-1] >= slice[i] {
			check = false
			return check
		} else if slice[i]-slice[i-1] > 3 {

			check = false
			return check
		} else {
			check = true
		}
	}
	return check

}

func checkDecreasing(slice []int) bool {
	check := false
	for i := 1; i < len(slice); i++ {
		if slice[i-1] <= slice[i] {
			check = false
			return check
		} else if slice[i-1]-slice[i] > 3 {

			check = false
			return check
		} else {
			check = true
		}
	}
	return check

}

func main() {

	count := 0
	report = readTextFile("input.txt")

	for _, item := range report {
		intLevels := make([]int, len(item.Levels))
		for i, level := range item.Levels {
			intLevels[i] = level.Value
		}
		safeInc := checkIncreasing(intLevels)
		//fmt.Println(safeInc)

		safeDec := checkDecreasing(intLevels)
		//fmt.Println(safeDec)

		if safeInc {
			count++
		}

		if safeDec {
			count++
		}
	}

	fmt.Println(count)
}
