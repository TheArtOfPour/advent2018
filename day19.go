package main

import (
	"bufio"
	"fmt"
	"strings"
)

type pos struct {
	row int
	col int
	dir rune
}

func advent19A(test string) string {
	// build maze
	var maze [][]rune
	scanner := bufio.NewScanner(strings.NewReader(test))
	for scanner.Scan() {
		maze = append(maze, []rune(scanner.Text()))
	}

	// find start position
	pos := pos{0, 0, 'd'}
	for col := range maze[0] {
		if maze[0][col] == '|' {
			pos.col = col
		}
	}

	// traverse maze
	result := ""
	for {
		if pos.col < 0 ||
			pos.row < 0 ||
			pos.row > len(maze) ||
			pos.col > len(maze[0]) ||
			maze[pos.row][pos.col] == ' ' {
			break
		}
		if maze[pos.row][pos.col] != '|' &&
			maze[pos.row][pos.col] != '-' &&
			maze[pos.row][pos.col] != '+' {
			result += string(maze[pos.row][pos.col])
		}
		fmt.Printf("%s", string(maze[pos.row][pos.col]))
		if maze[pos.row][pos.col] == '+' {
			if pos.dir == 'd' || pos.dir == 'u' {
				if pos.col+1 < len(maze[pos.row]) && maze[pos.row][pos.col+1] != ' ' {
					pos.dir = 'r'
				} else if pos.col-1 > 0 && maze[pos.row][pos.col-1] != ' ' {
					pos.dir = 'l'
				}
			} else {
				if pos.row+1 < len(maze) && maze[pos.row+1][pos.col] != ' ' {
					pos.dir = 'd'
				} else if pos.row-1 > 0 && maze[pos.row-1][pos.col] != ' ' {
					pos.dir = 'u'
				}
			}
		}
		switch pos.dir {
		case 'd':
			pos.row++
		case 'u':
			pos.row--
		case 'l':
			pos.col--
		case 'r':
			pos.col++
		default:
			panic(fmt.Sprintf("invalid direction %s", string(pos.dir)))
		}
	}

	// draw maze
	// for row := range maze {
	// 	for col := range maze[row] {
	// 		fmt.Printf("%s ", string(maze[row][col]))
	// 	}
	// 	fmt.Print("\n")
	// }
	return result
}

func advent19B(test string) int {
	// build maze
	var maze [][]rune
	scanner := bufio.NewScanner(strings.NewReader(test))
	for scanner.Scan() {
		maze = append(maze, []rune(scanner.Text()))
	}

	// find start position
	pos := pos{0, 0, 'd'}
	for col := range maze[0] {
		if maze[0][col] == '|' {
			pos.col = col
		}
	}

	// traverse maze
	result := 0
	for {
		if pos.col < 0 ||
			pos.row < 0 ||
			pos.row > len(maze) ||
			pos.col > len(maze[0]) ||
			maze[pos.row][pos.col] == ' ' {
			break
		}
		result++
		fmt.Printf("%s", string(maze[pos.row][pos.col]))
		if maze[pos.row][pos.col] == '+' {
			if pos.dir == 'd' || pos.dir == 'u' {
				if pos.col+1 < len(maze[pos.row]) && maze[pos.row][pos.col+1] != ' ' {
					pos.dir = 'r'
				} else if pos.col-1 > 0 && maze[pos.row][pos.col-1] != ' ' {
					pos.dir = 'l'
				}
			} else {
				if pos.row+1 < len(maze) && maze[pos.row+1][pos.col] != ' ' {
					pos.dir = 'd'
				} else if pos.row-1 > 0 && maze[pos.row-1][pos.col] != ' ' {
					pos.dir = 'u'
				}
			}
		}
		switch pos.dir {
		case 'd':
			pos.row++
		case 'u':
			pos.row--
		case 'l':
			pos.col--
		case 'r':
			pos.col++
		default:
			panic(fmt.Sprintf("invalid direction %s", string(pos.dir)))
		}
	}
	return result
}
