package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func advent4A(test string) (int, error) {
	resultGuard := 0
	resultMinute := 0
	guardNum := 0
	sleepStart := 0
	minutesPerGuard := make(map[int]int)
	stringSlice := strings.Split(test, "\n")
	sort.Strings(stringSlice)
	for _, s := range stringSlice {
		s = strings.TrimLeft(s, "[")
		logParts := strings.Split(s, "] ")
		timestamp := logParts[0]
		minutes, _ := strconv.Atoi(strings.Split(strings.Split(timestamp, " ")[1], ":")[1])
		logEntry := logParts[1]
		entryParts := strings.Split(logEntry, " ")
		if entryParts[0] == "Guard" {
			guardNum, _ = strconv.Atoi(strings.TrimLeft(entryParts[1], "#"))
			continue
		}
		if entryParts[0] == "falls" {
			sleepStart = minutes
			continue
		}
		if entryParts[0] == "wakes" {
			minutesPerGuard[guardNum] += minutes - sleepStart
			continue
		}
	}
	fmt.Printf("%v\n", minutesPerGuard)
	longestSleep := 0
	for guard, sleep := range minutesPerGuard {
		if sleep > longestSleep {
			resultGuard = guard
			longestSleep = sleep
		}
	}
	timesPerGuard := make(map[int]int)
	for _, s := range stringSlice {
		s = strings.TrimLeft(s, "[")
		logParts := strings.Split(s, "] ")
		timestamp := logParts[0]
		minutes, _ := strconv.Atoi(strings.Split(strings.Split(timestamp, " ")[1], ":")[1])
		logEntry := logParts[1]
		entryParts := strings.Split(logEntry, " ")
		if entryParts[0] == "Guard" {
			guardNum, _ = strconv.Atoi(strings.TrimLeft(entryParts[1], "#"))
			continue
		}
		if guardNum != resultGuard {
			continue
		}
		if entryParts[0] == "falls" {
			sleepStart = minutes
			continue
		}
		if entryParts[0] == "wakes" {
			for i := sleepStart; i < minutes; i++ {
				timesPerGuard[i]++
			}
			continue
		}
	}
	fmt.Printf("%v\n", timesPerGuard)
	longestSleep = 0
	for minute, sleep := range timesPerGuard {
		if sleep > longestSleep {
			resultMinute = minute
			longestSleep = sleep
		}
	}
	fmt.Printf("%d %d\n", resultGuard, resultMinute)
	return resultGuard * resultMinute, nil
}

func advent4B(test string) (int, error) {
	resultGuard := 0
	resultMinute := 0
	guardNum := 0
	sleepStart := 0
	minutesPerGuard := make(map[int]map[int]int)
	stringSlice := strings.Split(test, "\n")
	sort.Strings(stringSlice)
	for _, s := range stringSlice {
		s = strings.TrimLeft(s, "[")
		logParts := strings.Split(s, "] ")
		timestamp := logParts[0]
		minutes, _ := strconv.Atoi(strings.Split(strings.Split(timestamp, " ")[1], ":")[1])
		logEntry := logParts[1]
		entryParts := strings.Split(logEntry, " ")
		if entryParts[0] == "Guard" {
			guardNum, _ = strconv.Atoi(strings.TrimLeft(entryParts[1], "#"))
			continue
		}
		if entryParts[0] == "falls" {
			sleepStart = minutes
			continue
		}
		if entryParts[0] == "wakes" {
			for i := sleepStart; i < minutes; i++ {
				if minutesPerGuard[guardNum] == nil {
					minutesPerGuard[guardNum] = map[int]int{}
				}
				minutesPerGuard[guardNum][i]++
			}
			continue
		}
	}
	fmt.Printf("%v\n", minutesPerGuard)
	longestSleep := 0
	for guard, times := range minutesPerGuard {
		for minute, sleep := range times {
			if sleep > longestSleep {
				resultGuard = guard
				resultMinute = minute
				longestSleep = sleep
			}
		}
	}
	fmt.Printf("%d %d\n", resultGuard, resultMinute)
	return resultGuard * resultMinute, nil
}
