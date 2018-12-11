package main

import (
	"fmt"
	"strconv"
)

func calculatePower(x int, y int, serialNumber int) int {
	rackID := x + 10
	powerLevel := rackID * y
	powerLevel += serialNumber
	powerLevel *= rackID
	// get hundreds place
	hundredsPlace := 0
	powerLevel %= 1000
	if powerLevel >= 100 {
		slice := strconv.Itoa(powerLevel)
		hundredsPlace, _ = strconv.Atoi(string(slice[0]))
	}
	powerLevel = hundredsPlace - 5

	return powerLevel
}

func advent11A(test string) int {
	serialNumber, _ := strconv.Atoi(test)
	// for each cell (300x300) calculate power
	grid := make([][]int, 300)
	for i := range grid {
		grid[i] = make([]int, 300)
		for j := range grid[i] {
			grid[i][j] = calculatePower(i, j, serialNumber)
			//fmt.Printf("%d\t", grid[i][j])
		}
		//fmt.Printf("\n")
	}
	//fmt.Printf("\n")

	maxPower := 0
	maxX := 0
	maxY := 0
	for i := 0; i <= len(grid)-3; i++ {
		for j := 0; j <= len(grid[i])-3; j++ {
			clusterPower :=
				grid[i][j] + grid[i][j+1] + grid[i][j+2] +
					grid[i+1][j] + grid[i+1][j+1] + grid[i+1][j+2] +
					grid[i+2][j] + grid[i+2][j+1] + grid[i+2][j+2]
			if clusterPower > maxPower {
				maxPower = clusterPower
				maxX = i // shift for array index
				maxY = j
			}
		}
	}
	fmt.Printf("(%d,%d) %d \n", maxX, maxY, maxPower)

	// crawl through with 3x3 frames summing up and storing max

	return 0
}

func advent11B(test string) int {
	serialNumber, _ := strconv.Atoi(test)
	grid := make([][]int, 300)
	for i := range grid {
		grid[i] = make([]int, 300)
		for j := range grid[i] {
			grid[i][j] = calculatePower(i, j, serialNumber)
		}
	}

	maxPower := 0
	maxX := 0
	maxY := 0
	maxSize := 0

	for size := 0; size <= len(grid); size++ { // grid size
		for startX := 0; startX <= len(grid)-size; startX++ { // frame
			for startY := 0; startY <= len(grid)-size; startY++ { // frame
				clusterPower := 0
				for i := startX; i < startX+size; i++ {
					for j := startY; j < startY+size; j++ {
						clusterPower += grid[i][j]
					}
				}
				if clusterPower > maxPower {
					maxPower = clusterPower
					maxX = startX // shift for array index
					maxY = startY
					maxSize = size
				}
			}
		}
	}
	fmt.Printf("(%d,%d,%d) %d \n", maxX, maxY, maxSize, maxPower)

	// crawl through with 3x3 frames summing up and storing max

	return 0
}
