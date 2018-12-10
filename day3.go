package main

import (
	"strconv"
	"strings"
)

func buildMatrix(source string) [1000][1000]int {
	var twoD [1000][1000]int
	stringSlice := strings.Split(source, "\n")
	for _, s := range stringSlice {
		claimParts := strings.Split(s, " ")
		coords := strings.Split(claimParts[2], ",")
		x, _ := strconv.Atoi(coords[0])
		y, _ := strconv.Atoi(strings.TrimRight(coords[1], ":"))
		size := strings.Split(claimParts[3], "x")
		sizeY, _ := strconv.Atoi(strings.TrimRight(size[1], "\r"))
		sizeX, _ := strconv.Atoi(size[0])
		for i := x; i < x+sizeX; i++ {
			for j := y; j < y+sizeY; j++ {
				twoD[i][j]++
			}
		}
	}
	return twoD
}

func advent3A(test string) (int, error) {
	total := 0
	twoD := buildMatrix(test)
	for i := 0; i < 1000; i++ {
		for j := 0; j < 1000; j++ {
			if twoD[i][j] > 1 {
				total++
			}
		}
	}

	return total, nil
}

func advent3B(test string) (int, error) {
	claim := 0
	twoD := buildMatrix(test)
	stringSlice := strings.Split(test, "\n")
	for _, s := range stringSlice {
		claimParts := strings.Split(s, " ")
		claimNum, _ := strconv.Atoi(strings.TrimLeft(claimParts[0], "#"))
		coords := strings.Split(claimParts[2], ",")
		x, _ := strconv.Atoi(coords[0])
		y, _ := strconv.Atoi(strings.TrimRight(coords[1], ":"))
		size := strings.Split(claimParts[3], "x")
		sizeY, _ := strconv.Atoi(strings.TrimRight(size[1], "\r"))
		sizeX, _ := strconv.Atoi(size[0])
		isolated := true
		for i := x; i < x+sizeX; i++ {
			for j := y; j < y+sizeY; j++ {
				if twoD[i][j] != 1 {
					isolated = false
					break
				}
			}
			if !isolated {
				break
			}
		}
		if isolated {
			return claimNum, nil
		}
	}

	return claim, nil
}
