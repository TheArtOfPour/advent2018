package main

import (
	"fmt"
	"strconv"
	"strings"
)

func advent1A(test string) (int, error) {
	total := 0
	stringSlice := strings.Split(test, "\n")
	for _, s := range stringSlice {
		fmt.Printf("%v\n", s)
		s = strings.TrimSuffix(s, "\r")
		i, err := strconv.Atoi(s)
		if err != nil {
			return 0, err
		}
		total += i
		fmt.Printf("%d %d\n", i, total)

	}
	return total, nil
}

func advent1B(test string) (int, error) {
	total := 0
	seen := map[int]bool{}
	stringSlice := strings.Split(test, "\n")
	for true {
		for _, s := range stringSlice {
			if seen[total] {
				return total, nil
			}
			seen[total] = true
			//fmt.Printf("%v\n", total)
			s = strings.TrimSuffix(s, "\r")
			i, err := strconv.Atoi(s)
			if err != nil {
				return 0, err
			}
			total += i
		}
	}
	return total, nil
}
