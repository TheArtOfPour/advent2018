package main

import (
	"strconv"
	"strings"
)

func advent16A(test string, dancers string) string {
	commands := strings.Split(test, ",")
	for _, command := range commands {
		for i, rune := range command {
			if i == 0 {
				if rune == 's' {
					amount, _ := strconv.Atoi(string(command[1:]))
					dancers = dancers[len(dancers)-amount:] + dancers[:len(dancers)-amount]
					break
				}
				if rune == 'x' {
					parts := strings.Split(command[1:], "/")
					p1, _ := strconv.Atoi(parts[0])
					p2, _ := strconv.Atoi(parts[1])
					byteDancers := []byte(dancers)
					temp := byteDancers[p1]
					byteDancers[p1] = byteDancers[p2]
					byteDancers[p2] = temp
					dancers = string(byteDancers[:len(dancers)])
					break
				}
				if rune == 'p' {
					parts := strings.Split(command[1:], "/")
					p1 := strings.Index(dancers, parts[0])
					p2 := strings.Index(dancers, parts[1])
					byteDancers := []byte(dancers)
					temp := byteDancers[p1]
					byteDancers[p1] = byteDancers[p2]
					byteDancers[p2] = temp
					dancers = string(byteDancers[:len(dancers)])
					break
				}
			}
		}
	}
	return dancers
}

func advent16B(test string, dancerString string) string {
	dancers := []byte(dancerString)
	commands := strings.Split(test, ",")
	p1 := 0
	p2 := 0
	for b := 0; b < 100; b++ {
		for _, command := range commands {
			if command[0] == 's' {
				dancerString = string(dancers[:len(dancerString)])
				amount, _ := strconv.Atoi(string(command[1:]))
				dancerString = dancerString[len(dancerString)-amount:] + dancerString[:len(dancerString)-amount]
				dancers = []byte(dancerString)
			} else if command[0] == 'x' {
				parts := strings.Split(command[1:], "/")
				p1, _ = strconv.Atoi(parts[0])
				p2, _ = strconv.Atoi(parts[1])
				temp := dancers[p1]
				dancers[p1] = dancers[p2]
				dancers[p2] = temp
			} else if command[0] == 'p' {
				parts := strings.Split(command[1:], "/")
				for i, b := range dancers {
					if b == byte(parts[0][0]) {
						p1 = i
					} else if b == byte(parts[1][0]) {
						p2 = i
					}
				}
				temp := dancers[p1]
				dancers[p1] = dancers[p2]
				dancers[p2] = temp
			}
		}
	}
	dancerString = string(dancers[:len(dancerString)])
	return dancerString
}
