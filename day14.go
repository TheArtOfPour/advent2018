package main

import (
	"fmt"
	"strconv"
)

// HexToBin binary goodness
func HexToBin(hex string) (string, error) {
	ui, err := strconv.ParseUint(hex, 16, 64)
	if err != nil {
		return "", err
	}

	format := fmt.Sprintf("%%0%db", len(hex)*4)
	return fmt.Sprintf(format, ui), nil
}

func advent14A(test string) int {
	total := 0
	return total
}

func printDisk(disk [128][128]bool) {
	fmt.Printf("\n")
	diskString := ""
	for i := range disk {
		for j := range disk[i] {
			if disk[i][j] {
				diskString += "#"
			} else {
				diskString += "."
			}
		}
		diskString += "\n"
	}
	fmt.Printf("%s\n", diskString)
}

func defrag(disk *[128][128]bool, row int, col int) {
	disk[row][col] = false
	if row > 0 && disk[row-1][col] {
		defrag(disk, row-1, col)
	}
	if row < 127 && disk[row+1][col] {
		defrag(disk, row+1, col)
	}
	if col > 0 && disk[row][col-1] {
		defrag(disk, row, col-1)
	}
	if col < 127 && disk[row][col+1] {
		defrag(disk, row, col+1)
	}
}

func advent14B(test string) int {
	return 0
}
