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
	for i := 0; i < len(leftList); i++ {
		dif := rightList[i] - leftList[i]
		difs = difs + dif
	}
	return difs
}

func main() {

	readTextFile("input.txt")

	sumDifs(leftList, rightList)

	fmt.Printf("The sum of the differences is: %v\n", difs)
}
