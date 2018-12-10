package main

import (
	"bytes"
	"fmt"
	"strings"
)

func advent2A(test string) (int, error) {
	stringSlice := strings.Split(test, "\n")
	twiceCount := 0
	thriceCount := 0
	for _, s := range stringSlice {
		m := map[rune]int{}
		for _, r := range s {
			m[r]++
		}
		foundTwice := false
		foundThrice := false
		for _, value := range m {
			if value == 2 && !foundTwice {
				twiceCount++
				foundTwice = true
			} else if value == 3 && !foundThrice {
				thriceCount++
				foundThrice = true
			}
		}
	}

	return twiceCount * thriceCount, nil
}

func runeDiffCountByPosition(s1 string, s2 string) int {
	diff := 0
	fmt.Printf("%v\n", s1)
	fmt.Printf("%v\n", s2)
	for i := 0; i < len(s2); i++ {
		if s1[i] != s2[i] {
			diff++
		}
	}
	return diff
}

func getSameRunes(s1 string, s2 string) string {
	var buffer bytes.Buffer
	for i := 0; i < len(s1); i++ {
		if s1[i] == s2[i] {
			buffer.WriteString(string(s1[i]))
		}
	}
	return buffer.String()
}

func advent2B(test string) (string, error) {
	stringSlice := strings.Split(test, "\n")
	for i := 0; i < len(stringSlice)-1; i++ {
		s1 := stringSlice[i]
		for j := i + 1; j < len(stringSlice); j++ {
			s2 := stringSlice[j]
			diff := runeDiffCountByPosition(s1, s2)
			if diff == 1 {
				fmt.Printf("%s:%s\n", s1, s2)
				return getSameRunes(s1, s2), nil
			}
		}
	}

	return "", nil
}
