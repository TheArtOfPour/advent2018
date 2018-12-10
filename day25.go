package main

import (
	"bufio"
	"regexp"
	"strconv"
	"strings"
)

type instruction struct {
	write bool
	move  int
	next  rune
}

type instructions struct {
	start  rune
	steps  int
	states map[rune]map[bool]instruction
}

func getInstructions(input string) instructions {
	var instructions instructions
	instructions.states = make(map[rune]map[bool]instruction)
	var getState = regexp.MustCompile(`state ([A-Z]).`)
	var getStateAlt = regexp.MustCompile(`state ([A-Z]):`)
	var getNumber = regexp.MustCompile(`\d+`)
	var currentState rune
	var currentValue bool
	var currentInstruction instruction
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		s := scanner.Text()
		if strings.Contains(s, "Begin") {
			instructions.start = rune(strings.Split(getState.FindString(s), " ")[1][0])
		} else if strings.Contains(s, "Perform") {
			instructions.steps, _ = strconv.Atoi(getNumber.FindString(s))
		} else if strings.Contains(s, "In") {
			currentState = rune(strings.Split(getStateAlt.FindString(s), " ")[1][0])
			instructions.states[currentState] = make(map[bool]instruction)
		} else if strings.Contains(s, "current") {
			temp, _ := strconv.Atoi(getNumber.FindString(s))
			currentValue = temp == 1
			instructions.states[currentState][currentValue] = instruction{false, 0, ' '}
		} else if strings.Contains(s, "Write") {
			temp, _ := strconv.Atoi(getNumber.FindString(s))
			currentInstruction.write = temp == 1
		} else if strings.Contains(s, "Move") {
			if strings.Contains(s, "left") {
				currentInstruction.move = -1
			} else {
				currentInstruction.move = 1
			}
		} else if strings.Contains(s, "Continue") {
			currentInstruction.next = rune(strings.Split(getState.FindString(s), " ")[1][0])
			instructions.states[currentState][currentValue] = currentInstruction
		}
	}
	return instructions
}

func advent25A(test string) int {
	instructions := getInstructions(test)
	//fmt.Printf("%v", instructions)
	var tape []bool
	pos := 0
	state := instructions.start
	for i := 0; i < instructions.steps; i++ {
		if pos < 0 {
			tape = append([]bool{false}, tape...)
			pos = 0
		} else if pos >= len(tape) {
			tape = append(tape, false)
		}
		instruction := instructions.states[state][tape[pos]]
		tape[pos] = instruction.write
		pos += instruction.move
		state = instruction.next
		//fmt.Printf("%v\n", tape)
	}
	total := 0
	for _, pos := range tape {
		if pos {
			total++
		}
	}
	return total
}

func advent25B(test string) int {
	return 0
}
