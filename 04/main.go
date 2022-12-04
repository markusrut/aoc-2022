package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	bytes, error := os.ReadFile("input.txt")
	if error != nil {
		fmt.Println(error)
		return
	}

	input := string(bytes)
	lines := strings.Split(input, "\n")

	pairs := GetPairs(lines)

	fullOverlapCount, anyOverlapCount := CountOverlap(pairs)

	fmt.Println(fullOverlapCount)
	fmt.Println(anyOverlapCount)
}

func CountOverlap(pairs []pair) (int, int) {
	fullOverlap := 0
	anyOverlap := 0

	for _, pair := range pairs {
		if CheckFullOverlap(pair.first, pair.second) {
			fullOverlap++
		}
		if CheckAnyOverlap(pair.first, pair.second) {
			anyOverlap++
		}
	}

	return fullOverlap, anyOverlap
}

func CheckFullOverlap(assignment1 assignment, assignment2 assignment) bool {
	firstDoesEverything := assignment1.from >= assignment2.from && assignment1.to <= assignment2.to
	secondDoesEverything := assignment2.from >= assignment1.from && assignment2.to <= assignment1.to
	return firstDoesEverything || secondDoesEverything
}

func CheckAnyOverlap(assignment1 assignment, assignment2 assignment) bool {
	return assignment1.to >= assignment2.from && assignment1.from <= assignment2.to
}

func GetPairs(lines []string) []pair {
	pairs := []pair{}

	for _, line := range lines {
		elfs := strings.Split(line, ",")

		pairs = append(pairs, pair{
			first:  GetAssignmentFromElf(elfs[0]),
			second: GetAssignmentFromElf(elfs[1]),
		})
	}

	return pairs
}

func GetAssignmentFromElf(elf string) assignment {
	elfParts := strings.Split(elf, "-")
	from, _ := strconv.Atoi(elfParts[0])
	to, _ := strconv.Atoi(elfParts[1])

	return assignment{
		from: from,
		to:   to,
	}
}

type assignment struct {
	from int
	to   int
}

type pair struct {
	first  assignment
	second assignment
}
