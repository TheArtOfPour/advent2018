package main

import (
	"strconv"
	"strings"
)

func judge(a int, b int) int {
	if int16(a) == int16(b) {
		return 1
	}
	return 0
}

func advent15A(test string) int {
	seeds := strings.Split(test, ",")
	seedA, _ := strconv.Atoi(seeds[0])
	seedB, _ := strconv.Atoi(seeds[1])
	factorA := 16807
	factorB := 48271
	divisor := 2147483647
	matches := judge(seedA, seedB)
	for i := 0; i < 40000000; i++ {
		seedA = (seedA * factorA) % divisor
		seedB = (seedB * factorB) % divisor
		matches += judge(seedA, seedB)
	}
	return matches
}

func advent15B(test string) int {
	seeds := strings.Split(test, ",")
	seedA, _ := strconv.Atoi(seeds[0])
	seedB, _ := strconv.Atoi(seeds[1])
	factorA := 16807
	factorB := 48271
	divisor := 2147483647
	matches := judge(seedA, seedB)
	for i := 0; i < 5000000; i++ {
		seedA = (seedA * factorA) % divisor
		for seedA%4 != 0 {
			seedA = (seedA * factorA) % divisor
		}
		seedB = (seedB * factorB) % divisor
		for seedB%8 != 0 {
			seedB = (seedB * factorB) % divisor
		}
		matches += judge(seedA, seedB)
	}
	return matches
}
