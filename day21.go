package main

import (
	"bufio"
	"fmt"
	"strings"
)

type rule struct {
	pattern string
	output  string
}

var seed = `.#.
..#
###`

func sizeOf(s string) int {
	return len(strings.Split(s, "\n")[0])
}

func fract2(s string, rules []rule) string {
	rows := strings.Split(s, "\n")
	size := len(rows[0])
	var large []string
	var rebuilt []string
	count := 0
	for i := 0; i < size-1; i += 2 {
		rebuilt = append(rebuilt, "")
		rebuilt = append(rebuilt, "")
		rebuilt = append(rebuilt, "")
		rebuilt = append(rebuilt, "")
		for j := 0; j < size-1; j += 2 {
			var small []string
			small = append(small, rows[i][j:j+2])
			small = append(small, rows[i+1][j:j+2])
			large = applyRule(small, rules)
			rebuilt[count] += large[0]
			rebuilt[count+1] += large[1]
			rebuilt[count+2] += large[2]
		}
		count += 3
	}
	return strings.Join(rebuilt, "\n")
}

func fract3(s string, rules []rule) string {
	rows := strings.Split(s, "\n")
	size := len(rows[0])
	var large []string
	var rebuilt []string
	count := 0
	for i := 0; i < size-2; i += 3 {
		rebuilt = append(rebuilt, "")
		rebuilt = append(rebuilt, "")
		rebuilt = append(rebuilt, "")
		rebuilt = append(rebuilt, "")
		for j := 0; j < size-2; j += 3 {
			var small []string
			small = append(small, rows[i][j:j+3])
			small = append(small, rows[i+1][j:j+3])
			small = append(small, rows[i+2][j:j+3])
			large = applyRule(small, rules)
			rebuilt[count] += large[0]
			rebuilt[count+1] += large[1]
			rebuilt[count+2] += large[2]
			rebuilt[count+3] += large[3]
		}
		count += 4
	}
	return strings.Join(rebuilt, "\n")
}

func applyRule(s []string, rules []rule) []string {
	// for rotation, for flip, for rule
	for _, rule := range rules {
		testPattern := strings.Join(s, "/")
		if testPattern == rule.pattern {
			return strings.Split(rule.output, "/")
		}
	}
	return nil
}

func flip(rules []rule, pattern []string, output string) []rule {
	// flip horizontal
	var horiz []string
	temp := pattern
	if len(temp) == 2 {
		// 00 01  01 00
		// 10 11  11 10
		horiz = append(horiz, string(temp[0][1])+string(temp[0][0]))
		horiz = append(horiz, string(temp[1][1])+string(temp[1][0]))
		rules = append(rules, rule{strings.Join(horiz, "/"), output})
	} else {
		// 00 01 02  02 01 00
		// 10 11 12  12 11 10
		// 20 21 22  22 21 20
		horiz = append(horiz, string(temp[0][2])+string(temp[0][1])+string(temp[0][0]))
		horiz = append(horiz, string(temp[1][2])+string(temp[1][1])+string(temp[1][0]))
		horiz = append(horiz, string(temp[2][2])+string(temp[2][1])+string(temp[2][0]))
		rules = append(rules, rule{strings.Join(horiz, "/"), output})
	}

	// flip vertical
	var vert []string
	temp = pattern
	if len(temp) == 2 {
		// 00 01  10 11
		// 10 11  00 01
		vert = append(vert, string(temp[1][0])+string(temp[1][1]))
		vert = append(vert, string(temp[0][0])+string(temp[0][1]))
		rules = append(rules, rule{strings.Join(vert, "/"), output})
	} else {
		// 00 01 02  20 21 22
		// 10 11 12  10 11 12
		// 20 21 22  00 01 02
		vert = append(vert, string(temp[2][0])+string(temp[2][1])+string(temp[2][2]))
		vert = append(vert, string(temp[1][0])+string(temp[1][1])+string(temp[1][2]))
		vert = append(vert, string(temp[0][0])+string(temp[0][1])+string(temp[0][2]))
		rules = append(rules, rule{strings.Join(vert, "/"), output})
	}
	return rules
}

func advent21A(test string) int {
	iterations := 5
	var rules []rule
	scanner := bufio.NewScanner(strings.NewReader(test))
	for scanner.Scan() {
		s := scanner.Text()
		parts := strings.Split(s, " => ")
		rules = append(rules, rule{parts[0], parts[1]})
		temp := strings.Split(parts[0], "/")
		rules = flip(rules, temp, parts[1])
		// rotate clockwise 3x
		for i := 0; i < 2; i++ {
			var rotated []string
			if len(temp) == 2 {
				// 00 01  10 00
				// 10 11  11 01
				rotated = append(rotated, string(temp[1][0])+string(temp[0][0]))
				rotated = append(rotated, string(temp[1][1])+string(temp[0][1]))
				rules = append(rules, rule{strings.Join(rotated, "/"), parts[1]})
			} else {
				// 00 01 02  20 10 00
				// 10 11 12  21 11 01
				// 20 21 22  22 12 02
				rotated = append(rotated, string(temp[2][0])+string(temp[1][0])+string(temp[0][0]))
				rotated = append(rotated, string(temp[2][1])+string(temp[1][1])+string(temp[0][1]))
				rotated = append(rotated, string(temp[2][2])+string(temp[1][2])+string(temp[0][2]))
				rules = append(rules, rule{strings.Join(rotated, "/"), parts[1]})
			}
			temp = rotated
			rules = flip(rules, temp, parts[1])
		}
	}
	pattern := seed
	for i := 0; i < iterations; i++ {
		if sizeOf(pattern)%2 == 0 {
			pattern = fract2(pattern, rules)
		} else {
			pattern = fract3(pattern, rules)
		}
		fmt.Printf("\n-------------------------------\n%v\n", pattern)
	}
	result := 0
	for _, rune := range pattern {
		if rune == '#' {
			result++
		}
	}

	return result
}

func advent21B(test string) int {
	iterations := 18
	var rules []rule
	scanner := bufio.NewScanner(strings.NewReader(test))
	for scanner.Scan() {
		s := scanner.Text()
		parts := strings.Split(s, " => ")
		rules = append(rules, rule{parts[0], parts[1]})
		temp := strings.Split(parts[0], "/")
		rules = flip(rules, temp, parts[1])
		// rotate clockwise 3x
		for i := 0; i < 2; i++ {
			var rotated []string
			if len(temp) == 2 {
				// 00 01  10 00
				// 10 11  11 01
				rotated = append(rotated, string(temp[1][0])+string(temp[0][0]))
				rotated = append(rotated, string(temp[1][1])+string(temp[0][1]))
				rules = append(rules, rule{strings.Join(rotated, "/"), parts[1]})
			} else {
				// 00 01 02  20 10 00
				// 10 11 12  21 11 01
				// 20 21 22  22 12 02
				rotated = append(rotated, string(temp[2][0])+string(temp[1][0])+string(temp[0][0]))
				rotated = append(rotated, string(temp[2][1])+string(temp[1][1])+string(temp[0][1]))
				rotated = append(rotated, string(temp[2][2])+string(temp[1][2])+string(temp[0][2]))
				rules = append(rules, rule{strings.Join(rotated, "/"), parts[1]})
			}
			temp = rotated
			rules = flip(rules, temp, parts[1])
		}
	}
	pattern := seed
	for i := 0; i < iterations; i++ {
		if sizeOf(pattern)%2 == 0 {
			pattern = fract2(pattern, rules)
		} else {
			pattern = fract3(pattern, rules)
		}
		fmt.Printf("\n-------------------------------\n%v\n", pattern)
	}
	result := 0
	for _, rune := range pattern {
		if rune == '#' {
			result++
		}
	}
	return result
}
