package main

import (
	"fmt"
	"io/ioutil"
	"sort"
)

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

// SortString sorts a string
func SortString(s string) string {
	r := []rune(s)
	sort.Sort(sortRunes(r))
	return string(r)
}

func main() {
	fileContents, _ := ioutil.ReadFile("./inputs/day12.txt")
	input := string(fileContents)
	out := advent12B(input)
	fmt.Printf("Result %v\n", out)
}
