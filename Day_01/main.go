package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var leftList []int
var rightList []int
var value int
var difs int
var simScore int

func readTextFile(filepath string) ([]int, []int) {

	fmt.Println("Reading text file")
	file, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		if len(parts) == 4 {
			value, err = strconv.Atoi(parts[0])
			if err != nil {
				panic(err)
			}
			leftList = append(leftList, value)
			value, err = strconv.Atoi(parts[3])
			if err != nil {
				panic(err)
			}
			rightList = append(rightList, value)
		}
	}

	sort.Ints(leftList)
	sort.Ints(rightList)
	return leftList, rightList

}

func sumDifs(leftList []int, rightList []int) int {
	var dif int
	for i := 0; i < len(leftList); i++ {
		if leftList[i] > rightList[i] {
			dif = leftList[i] - rightList[i]
		} else {
			dif = rightList[i] - leftList[i]
		}
		difs = difs + dif
	}
	return difs
}

func countInArray(array []int, value int) int {
	count := 0
	for i := 0; i < len(array); i++ {
		if array[i] == value {
			count++
		}
	}
	return count
}

func findSimScore(leftList []int, rightList []int) int {

	simScore := 0

	for i := 0; i < len(leftList); i++ {
		count := countInArray(rightList, leftList[i])
		simScore = simScore + (leftList[i] * count)
	}
	return simScore

}

func main() {

	readTextFile("input.txt")

	sumDifs(leftList, rightList)
	simScore = findSimScore(leftList, rightList)

	fmt.Printf("The sum of the differences is: %v\n", difs)
	fmt.Printf("The similarity score is: %v\n", simScore)
}
