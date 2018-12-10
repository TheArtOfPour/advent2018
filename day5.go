package main

import (
	"bytes"
	"strings"
	"unicode"
)

func advent5A(test string) (int, error) {
	lastRune := ' '
	modified := false
	skip := false
	var newString bytes.Buffer
	//fmt.Printf("%v\n", test)
	for i, r := range test {
		if i == 0 {
			lastRune = r
			continue
		}
		if skip {
			skip = false
			lastRune = r
			continue
		}
		if unicode.IsLower(lastRune) != unicode.IsLower(r) {
			// different cases
			if unicode.ToLower(lastRune) == unicode.ToLower(r) {
				// matched pair, remove
				//fmt.Printf("Match: %s %s\n", string(lastRune), string(r))
				modified = true
				skip = true
				lastRune = r
				continue
			}
		}
		newString.WriteString(string(lastRune))
		lastRune = r
	}
	newString.WriteString(string(lastRune))
	if !modified {
		return len(newString.String()), nil
	}
	return advent5A(newString.String())
}

func advent5B(test string) (int, error) {
	// remove all pairs for each letter then run A
	upper := `ABCDEFGHIJKLMNOPQRSTUVWXYZ`
	lower := `abcdefghijklmnopqrstuvwxyz`
	bestResult := -1
	for i := 0; i < 26; i++ {
		copy := test
		copy = strings.Replace(copy, string(upper[i]), "", -1)
		copy = strings.Replace(copy, string(lower[i]), "", -1)
		//fmt.Printf("TEST: %v\n", copy)
		result, _ := advent5A(copy)
		if bestResult == -1 {
			bestResult = result
			continue
		}
		if result < bestResult {
			bestResult = result
		}
	}
	return bestResult, nil
}
