package main

import (
	"fmt"
	"strconv"
	"strings"
)

func advent9A(test string) int {
	stringSlice := strings.Split(test, " ")
	numPlayers, _ := strconv.Atoi(stringSlice[0])
	lastMarble, _ := strconv.Atoi(stringSlice[6])
	playerScores := make(map[int]int, numPlayers)
	board := make([]int, 0)
	board = append(board, 0)
	currentPlayer := 0
	currentIndex := 0
	for currentMarble := 1; currentMarble <= lastMarble; currentMarble++ {
		// mult of 23, keep marble + remove and keep marble 7 back
		if currentMarble%23 == 0 {
			// add current marble to score
			playerScores[currentPlayer] += currentMarble
			// calculate 7 marbles back
			marbleToRemoveIndex := currentIndex - 7
			if marbleToRemoveIndex < 0 {
				marbleToRemoveIndex = len(board) + marbleToRemoveIndex
			}
			// add to score
			playerScores[currentPlayer] += board[marbleToRemoveIndex]
			// remove from board
			board = append(board[:marbleToRemoveIndex], board[marbleToRemoveIndex+1:]...)
			// print board
			// fmt.Printf("[%d] (%d) %v\n", currentPlayer+1, board[marbleToRemoveIndex], board)
			// reset index
			currentIndex = marbleToRemoveIndex
			// iterate player
			currentPlayer++
			if currentPlayer >= numPlayers {
				currentPlayer = 0
			}
			continue
		}

		nextMarbleIndex := currentIndex + 2 // calculate next marble
		if currentMarble == 1 {
			nextMarbleIndex = 1
		}
		if nextMarbleIndex > len(board) {
			nextMarbleIndex -= len(board)
		}
		// insert into board
		board = append(board, 0)                                 // enlarge
		copy(board[nextMarbleIndex+1:], board[nextMarbleIndex:]) // shift
		board[nextMarbleIndex] = currentMarble                   // insert

		// print board
		// fmt.Printf("[%d] (%d) %v\n", currentPlayer+1, board[nextMarbleIndex], board)

		currentIndex = nextMarbleIndex

		// iterate player
		currentPlayer++
		if currentPlayer >= numPlayers {
			currentPlayer = 0
		}
		if currentMarble%10000 == 0 {
			fmt.Printf("%d\n", currentMarble)
		}
	}

	// fmt.Printf("\n%v\n\n", playerScores)

	maxScore := 0
	for _, score := range playerScores {
		if score > maxScore {
			maxScore = score
		}
	}

	return maxScore
}

func advent9B(test string) int {
	return 0
}
