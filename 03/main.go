package main

import (
	"fmt"
	"os"
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

	matches := GetMatches(lines)
	score := GetScoreFromMatches(matches)

	groupMatches := GetGroupMatches(lines)
	groupScore := GetScoreFromMatches(groupMatches)

	fmt.Println(score)
	fmt.Println(groupScore)
}

func GetScoreFromMatches(matches []rune) int {
	score := 0
	for _, m := range matches {
		isUppercase := m >= 'A' && m <= 'Z'

		if isUppercase {
			score += int(m - 38)
		} else {
			score += int(m - 96)
		}
	}

	return score
}

func GetMatches(lines []string) []rune {
	matches := []rune{}

	for _, line := range lines {
		halfSize := len(line) / 2
		first := line[:halfSize]
		second := line[halfSize:]

		matches = append(matches, GetMatchInTarget(first, []string{second}))
	}

	return matches
}

func GetGroupMatches(lines []string) []rune {
	matches := []rune{}

	// split lines in groups of 3
	for i := 0; i < len(lines); i += 3 {
		first := lines[i]
		second := lines[i+1]
		third := lines[i+2]

		matches = append(matches, GetMatchInTarget(first, []string{second, third}))
	}

	return matches
}

func GetMatchInTarget(source string, targets []string) rune {
	matches := []rune(source)

	for _, target := range targets {
		newMatches := []rune{}
		for _, char := range matches {
			if strings.Contains(target, string(char)) {
				newMatches = append(newMatches, char)
			}
		}
		matches = newMatches
	}

	return matches[0]
}
