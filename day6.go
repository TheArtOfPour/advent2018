package main

import (
	"fmt"
	"strconv"
	"strings"
)

// type coord struct {
// 	X int
// 	Y int
// }

// func absint(n int) int {
// 	if n < 0 {
// 		return -n
// 	}
// 	return n
// }

func printGrid(grid [][]int) {
	for i := range grid {
		for j := range grid[i] {
			fmt.Printf("%d ", grid[i][j])
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
}

func getUnbound(coords map[int]coord) map[int]bool {
	unbound := make(map[int]bool)
	minX := 100000
	maxX := 0
	minY := 100000
	maxY := 0
	for _, c := range coords {
		if c.X < minX {
			minX = c.X
		}
		if c.X > maxX {
			maxX = c.X
		}
		if c.Y < minY {
			minY = c.Y
		}
		if c.Y > maxY {
			maxY = c.Y
		}
	}
	for key, c := range coords {
		if c.X == minX || c.X == maxX || c.Y == minY || c.Y == maxY {
			unbound[key] = true
		}
	}
	return unbound
}

func convertToCoords(s string) map[int]coord {
	coords := make(map[int]coord)
	cs := strings.Split(s, "\n")
	for i, c := range cs {
		coordParts := strings.Split(c, ", ")
		x, _ := strconv.Atoi(coordParts[0])
		y, _ := strconv.Atoi(strings.TrimRight(coordParts[1], "\r"))
		coords[i] = coord{X: x, Y: y}
	}
	return coords
}

func distance(c1 coord, c2 coord) int {
	return absint(c1.X-c2.X) + absint(c1.Y-c2.Y)
}

func getArea(c coord, coords map[int]coord) int {
	// do circles until full circle without increment
	area := 1
	layer := 1
	for true {
		lastArea := area
		// top line
		for i := c.X - layer; i < c.X+layer; i++ {
			d := distance(coord{X: i, Y: c.Y + layer}, c)
			closest := true
			for _, otherCoord := range coords {
				if distance(coord{X: i, Y: c.Y + layer}, otherCoord) <= d && c.X != otherCoord.X && c.Y != otherCoord.Y {
					closest = false
					break
				}
			}
			if closest {
				area++
			}
		}
		// right line
		for i := c.Y + layer; i > c.Y-layer; i-- {
			d := distance(coord{X: c.X + layer, Y: i}, c)
			closest := true
			for _, otherCoord := range coords {
				if distance(coord{X: c.X + layer, Y: i}, otherCoord) <= d && c.X != otherCoord.X && c.Y != otherCoord.Y {
					closest = false
					break
				}
			}
			if closest {
				area++
			}
		}
		// bottom line
		for i := c.X + layer; i > c.X-layer; i-- {
			d := distance(coord{X: i, Y: c.Y - layer}, c)
			closest := true
			for _, otherCoord := range coords {
				if distance(coord{X: i, Y: c.Y - layer}, otherCoord) <= d && c.X != otherCoord.X && c.Y != otherCoord.Y {
					closest = false
					break
				}
			}
			if closest {
				area++
			}
		}
		// left line
		for i := c.Y - layer; i < c.Y+layer; i++ {
			d := distance(coord{X: c.X - layer, Y: i}, c)
			closest := true
			for _, otherCoord := range coords {
				if distance(coord{X: c.X - layer, Y: i}, otherCoord) <= d && c.X != otherCoord.X && c.Y != otherCoord.Y {
					closest = false
					break
				}
			}
			if closest {
				area++
			}
		}
		if lastArea == area {
			break
		}
		layer++
	}
	return area
}

func advent6A(test string) (int, error) {
	coords := convertToCoords(test)
	unbound := getUnbound(coords)
	//fmt.Printf("%v %v \n", coords, unbound)
	maxArea := 0
	for i, coord := range coords {
		isUnbound := false
		for unboundCoord := range unbound {
			if i == unboundCoord {
				isUnbound = true
				break
			}
		}
		if isUnbound {
			continue
		}
		area := getArea(coord, coords)
		fmt.Printf("%v %v \n", coord, area)
		if area > maxArea {
			maxArea = area
		}
	}
	return maxArea, nil
}

func advent6B(test string) (int, error) {

	return 0, nil
}
