package main

import (
	"fmt"
	"os"
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
	rounds := getRounds(input)

	opponentRock, opponentPaper, opponentScissors := "A", "B", "C"
	selfRock, selfPaper, selfScissors := "X", "Y", "Z"
	selfLose, selfDraw, selfWin := "X", "Y", "Z"

	scoreRock, scorePaper, scoreScissors, scoreLost, scoreDraw, scoreWin := 1, 2, 3, 0, 3, 6

	score := 0
	score2 := 0

	for _, round := range rounds {
		if round.opponent == opponentRock && round.self == selfRock {
			score += scoreDraw
			score += scoreRock
		} else if round.opponent == opponentRock && round.self == selfPaper {
			score += scoreWin
			score += scorePaper
		} else if round.opponent == opponentRock && round.self == selfScissors {
			score += scoreLost
			score += scoreScissors
		} else if round.opponent == opponentPaper && round.self == selfRock {
			score += scoreLost
			score += scoreRock
		} else if round.opponent == opponentPaper && round.self == selfPaper {
			score += scoreDraw
			score += scorePaper
		} else if round.opponent == opponentPaper && round.self == selfScissors {
			score += scoreWin
			score += scoreScissors
		} else if round.opponent == opponentScissors && round.self == selfRock {
			score += scoreWin
			score += scoreRock
		} else if round.opponent == opponentScissors && round.self == selfPaper {
			score += scoreLost
			score += scorePaper
		} else if round.opponent == opponentScissors && round.self == selfScissors {
			score += scoreDraw
			score += scoreScissors
		}

		if round.opponent == opponentRock && round.self == selfWin {
			score2 += scoreWin
			score2 += scorePaper
		} else if round.opponent == opponentRock && round.self == selfDraw {
			score2 += scoreDraw
			score2 += scoreRock
		} else if round.opponent == opponentRock && round.self == selfLose {
			score2 += scoreLost
			score2 += scoreScissors
		} else if round.opponent == opponentPaper && round.self == selfWin {
			score2 += scoreWin
			score2 += scoreScissors
		} else if round.opponent == opponentPaper && round.self == selfDraw {
			score2 += scoreDraw
			score2 += scorePaper
		} else if round.opponent == opponentPaper && round.self == selfLose {
			score2 += scoreLost
			score2 += scoreRock
		} else if round.opponent == opponentScissors && round.self == selfWin {
			score2 += scoreWin
			score2 += scoreRock
		} else if round.opponent == opponentScissors && round.self == selfDraw {
			score2 += scoreDraw
			score2 += scoreScissors
		} else if round.opponent == opponentScissors && round.self == selfLose {
			score2 += scoreLost
			score2 += scorePaper
		}
	}

	fmt.Println(score)
	fmt.Println(score2)
}

func getRounds(input string) []round {
	lines := strings.Split(input, "\n")

	var rounds []round

	for _, line := range lines {
		parts := strings.Split(line, " ")
		round := round{opponent: parts[0], self: parts[1]}
		rounds = append(rounds, round)
	}

	return rounds
}

type round struct {
	opponent string
	self     string
}
