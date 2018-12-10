package main

import (
	"fmt"
	"strconv"
	"strings"
)

type coord struct {
	X int
	Y int
}

type light struct {
	Position coord
	Velocity coord
}

func getLights(s string) ([]light, coord) {
	lights := make([]light, 0)
	lightStrings := strings.Split(s, "\n")
	sumX := 0
	sumY := 0
	for _, ls := range lightStrings {
		ls = strings.TrimSuffix(ls, "\r")
		ls = strings.TrimLeft(ls, "position=<")
		ls = strings.TrimRight(ls, ">")
		parts := strings.Split(ls, "> velocity=<")
		positionParts := strings.Split(parts[0], ", ")
		x, _ := strconv.Atoi(strings.TrimLeft(positionParts[0], " "))
		y, _ := strconv.Atoi(strings.TrimLeft(positionParts[1], " "))
		sumX += x
		sumY += y
		position := coord{X: x, Y: y}
		velocityParts := strings.Split(parts[1], ", ")
		x, _ = strconv.Atoi(strings.TrimLeft(velocityParts[0], " "))
		y, _ = strconv.Atoi(strings.TrimLeft(velocityParts[1], " "))
		velocity := coord{X: x, Y: y}
		light := light{Position: position, Velocity: velocity}
		lights = append(lights, light)
	}
	count := len(lightStrings)
	meanCoord := coord{X: sumX / count, Y: sumY / count}
	return lights, meanCoord
}

func absint(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func getError(c1 coord, c2 coord) int {
	return absint(c1.X-c2.X) + absint(c1.Y-c2.Y)
}

func printLights(lights []light) {
	// get min/max x/y bounds and crop
	minX := 10000000
	maxX := 0
	minY := 10000000
	maxY := 0
	for _, l := range lights {
		if l.Position.X < minX {
			minX = l.Position.X
		}
		if l.Position.X > maxX {
			maxX = l.Position.X
		}
		if l.Position.Y < minY {
			minY = l.Position.Y
		}
		if l.Position.Y > maxY {
			maxY = l.Position.Y
		}
	}
	fmt.Printf("Range X[%d-%d] Y[%d-%d]\n", minX, maxX, minY, maxY)
	sky := make([][]rune, absint(maxY)+absint(minY)+1)
	for i := range sky {
		sky[i] = make([]rune, absint(maxX)+absint(minX)+1)
	}
	// add lights to sky
	for _, l := range lights {
		sky[l.Position.Y-minY][l.Position.X-minX] = '#'
	}
	// print sky
	fmt.Printf("\n")
	for i := range sky {
		for j := range sky[i] {
			if sky[i][j] != '#' {
				sky[i][j] = ' '
			}
			fmt.Printf("%s", string(sky[i][j]))
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
}

func advent10A(test string) int {
	// or maybe calculate line intersections?
	lights, meanCoord := getLights(test)
	count := len(lights)
	epsilon := 99999999
	closestLights := make([]light, len(lights))
	// iterate until error stops decreasing
	numberOfLoops := 0
	for true {
		numberOfLoops++
		// iterate over lights updating positions and calculating error
		newError := 0
		sumX := 0
		sumY := 0
		copy(closestLights, lights)
		for i := range lights {
			lights[i].Position.X = lights[i].Position.X + lights[i].Velocity.X
			lights[i].Position.Y = lights[i].Position.Y + lights[i].Velocity.Y
			sumX += lights[i].Position.X
			sumY += lights[i].Position.Y
			newError += getError(meanCoord, lights[i].Position)
		}
		meanCoord = coord{X: sumX / count, Y: sumY / count}
		//fmt.Printf("Error: %d\n", newError)
		if newError >= epsilon {
			fmt.Printf("Error: %d\n", newError)
			break
		}
		epsilon = newError
	}

	fmt.Printf("%v\n", numberOfLoops)

	// display "sky" result
	printLights(closestLights)

	return 0
}

func advent10B(test string) int {
	// or maybe calculate line intersections?
	lights, meanCoord := getLights(test)
	count := len(lights)
	epsilon := 99999999
	// iterate until error stops decreasing
	numberOfLoops := 0
	for true {
		numberOfLoops++
		// iterate over lights updating positions and calculating error
		newError := 0
		sumX := 0
		sumY := 0
		for i := range lights {
			lights[i].Position.X = lights[i].Position.X + lights[i].Velocity.X
			lights[i].Position.Y = lights[i].Position.Y + lights[i].Velocity.Y
			sumX += lights[i].Position.X
			sumY += lights[i].Position.Y
			newError += getError(meanCoord, lights[i].Position)
		}
		meanCoord = coord{X: sumX / count, Y: sumY / count}
		if newError >= epsilon {
			break
		}
		epsilon = newError
	}

	return numberOfLoops - 1
}
