package main

import (
	"fmt"
	"strings"
)

type grule struct {
	condition string
	result    rune
}

func advent12A(test string) int {
	result := 0

	stringSlice := strings.Split(test, "\n")

	// get initial state
	plants := strings.TrimLeft(stringSlice[0], "initial state: ")
	plants = strings.TrimSuffix(plants, "\r")

	// get rules
	grules := make([]grule, 0)
	for _, s := range stringSlice[2:] {
		s = strings.TrimSuffix(s, "\r")
		parts := strings.Split(s, " => ")
		r := grule{condition: parts[0], result: rune(parts[1][0])}
		grules = append(grules, r)
	}

	// iterate through generations
	indexOffset := 0 // subtract from this every time a plant is added to the front of the array
	//plants = "..." + plants
	for generation := 0; generation <= 19; generation++ {
		plants = ".." + plants
		if plants[2] == '#' {
			plants = "." + plants
			indexOffset--
		}
		plants = plants + "..."
		if generation < 10 {
			fmt.Print(" ")
		}
		fmt.Printf("%d: ", generation)
		for i := 0; i < 10+indexOffset; i++ {
			fmt.Print(" ")
		}
		fmt.Printf("%s\n", plants)
		newPlants := make([]rune, 0)
		for i := 2; i < len(plants)-2; i++ {
			frame := string([]byte{plants[i-2], plants[i-1], plants[i], plants[i+1], plants[i+2]})
			found := false
			for _, rule := range grules {
				if frame == rule.condition {
					newPlants = append(newPlants, rule.result)
					found = true
					break
				}
			}
			if !found {
				newPlants = append(newPlants, '.')
			}
		}
		plants = string(newPlants)
	}

	for i := 0; i < len(plants); i++ {
		if plants[i] == '#' {
			result += i + indexOffset
		}
	}

	return result
}

func advent12B(test string) int {
	result := 0

	stringSlice := strings.Split(test, "\n")

	// get initial state
	plants := strings.TrimLeft(stringSlice[0], "initial state: ")
	plants = strings.TrimSuffix(plants, "\r")
	// plantsBool := make([]bool, 0)
	// for i := 0; i < len(plants); i++ {
	// 	plantsBool = append(plantsBool, plants[i] == '#')
	// }

	// get rules
	grules := make([]grule, 0)
	for _, s := range stringSlice[2:] {
		s = strings.TrimSuffix(s, "\r")
		parts := strings.Split(s, " => ")
		r := grule{condition: parts[0], result: rune(parts[1][0])}
		grules = append(grules, r)
	}

	// iterate through generations
	indexOffset := 0
	for generation := 0; generation <= 252; generation++ {
		plants = ".." + plants
		//plantsBool = append([]bool{false, false}, plantsBool...)
		if plants[2] == '#' {
			plants = "." + plants
			//plantsBool = append([]bool{false}, plantsBool...)
			indexOffset--
		}
		if generation < 10 {
			fmt.Print(" ")
		}
		if generation < 100 {
			fmt.Print(" ")
		}
		fmt.Printf("%d: ", generation)
		for i := 0; i < 10+indexOffset; i++ {
			fmt.Print(" ")
		}
		fmt.Printf("%s\n", plants)
		plants = plants + "..."
		//plantsBool = append(plantsBool, []bool{false, false, false})
		newPlants := make([]rune, 0)
		for i := 2; i < len(plants)-2; i++ {
			frame := string([]byte{plants[i-2], plants[i-1], plants[i], plants[i+1], plants[i+2]})
			found := false
			for _, rule := range grules {
				if frame == rule.condition {
					newPlants = append(newPlants, rule.result)
					found = true
					break
				}
			}
			if !found {
				newPlants = append(newPlants, '.')
			}
		}
		plants = string(newPlants)
	}

	for i := 0; i < len(plants); i++ {
		if plants[i] == '#' {
			result += i + indexOffset
		}
	}

	return result
}
