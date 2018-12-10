package main

import (
	"fmt"
	"strconv"
	"strings"
)

func advent23A(test string) int {
	mulTime := 0
	registers := make(map[string]int)
	commands := strings.Split(test, "\r\n")
	i := 0
	for {
		if i < 0 || i >= len(commands) {
			break
		}
		command := commands[i]
		stringParts := strings.Split(command, " ")
		switch stringParts[0] {
		case "set":
			numeric, err := strconv.Atoi(stringParts[2])
			if err != nil {
				registers[stringParts[1]] = registers[stringParts[2]]
			} else {
				registers[stringParts[1]] = numeric
			}
		case "sub":
			numeric, err := strconv.Atoi(stringParts[2])
			if err != nil {
				registers[stringParts[1]] -= registers[stringParts[2]]
			} else {
				registers[stringParts[1]] -= numeric
			}
		case "mul":
			mulTime++
			numeric, err := strconv.Atoi(stringParts[2])
			if err != nil {
				registers[stringParts[1]] *= registers[stringParts[2]]
			} else {
				registers[stringParts[1]] *= numeric
			}
		case "jnz":
			numeric, err := strconv.Atoi(stringParts[1])
			if err != nil {
				numeric = registers[stringParts[1]]
			}
			if numeric != 0 {
				numeric, err := strconv.Atoi(stringParts[2])
				if err != nil {
					numeric = registers[stringParts[2]]
				}
				i += numeric - 1
			}
		default:
			panic(fmt.Sprintf("invalid command %s", stringParts[0]))
		}
		//fmt.Printf("%d: %s -> %v\n", i, command, registers)
		i++
	}
	return mulTime
}

func advent23B(test string) int {
	b := 84*100 + 100000
	c := b + 17000
	d := 0
	f := 0
	g := 0
	h := 0

	for {
		f = 1
		d = 2
		if b%d == 0 {
			f = 0
		}
		for d != b {
			if b%d == 0 {
				f = 0
			}
			d++
		}

		if f == 0 {
			h++
		}
		g = b - c
		if g == 0 {
			break
		}
		b += 17
	}
	return h
}
