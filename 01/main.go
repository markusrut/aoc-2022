package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	// read all liens from input.txt
	bytes, error := os.ReadFile("input.txt")
	if error != nil {
		fmt.Println(error)
		return
	}

	input := string(bytes)

	// split input into groups by empty lines
	groups := splitByEmptyLines(input)

	var sums []int
	var currentSum int
	for _, group := range groups {
		currentSum = 0
		for _, str := range group {
			num, _ := strconv.Atoi(str)
			currentSum += num
		}
		sums = append(sums, currentSum)
	}

	sort.Ints(sums)
	// convert bytes to string

	fmt.Println(sums[len(sums)-1])

	// take last 3 from sums into separate list
	lastThree := sums[len(sums)-3:]

	// multiply them
	var result int
	for _, num := range lastThree {
		result += num
	}

	fmt.Println(result)
}

func splitByEmptyLines(input string) [][]string {
	list := strings.Split(input, "\n")

	var groups [][]string
	var currentGroup []string

	for _, line := range list {
		if line == "" {
			groups = append(groups, currentGroup)
			currentGroup = []string{}
		} else {
			currentGroup = append(currentGroup, line)
		}
	}

	return groups
}
