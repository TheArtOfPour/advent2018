package main

import (
	"fmt"
	"strings"
)

func abs(i int) int {
	if i < 0 {
		i *= -1
	}
	return i
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func getDistance(coords map[string]int) int {
	//  (x,  y, z)
	//+ (ne, n, nw)
	//- (sw, s, se)
	if coords["x"] > 0 && coords["y"] < 0 {
		min := min(coords["x"], abs(coords["y"]))
		coords["x"] -= min
		coords["y"] += min
		coords["z"] -= min
	} else if coords["y"] > 0 && coords["z"] < 0 {
		min := min(coords["y"], abs(coords["z"]))
		coords["y"] -= min
		coords["z"] += min
		coords["x"] += min
	} else if coords["z"] > 0 && coords["x"] > 0 {
		min := min(coords["z"], coords["x"])
		coords["z"] -= min
		coords["x"] -= min
		coords["y"] += min
	} else if coords["y"] > 0 && coords["x"] < 0 {
		min := min(coords["y"], abs(coords["x"]))
		coords["y"] -= min
		coords["x"] += min
		coords["z"] += min
	} else if coords["z"] > 0 && coords["y"] < 0 {
		min := min(coords["z"], abs(coords["y"]))
		coords["z"] -= min
		coords["y"] += min
		coords["x"] -= min
	} else if coords["x"] < 0 && coords["z"] < 0 {
		min := min(abs(coords["x"]), abs(coords["z"]))
		coords["x"] += min
		coords["z"] += min
		coords["y"] -= min
	}
	return abs(coords["x"]) + abs(coords["y"]) + abs(coords["z"])
}

func advent11A(test string) int {
	//  (x,  y, z)
	//+ (ne, n, nw)
	//- (sw, s, se)
	coords := make(map[string]int)
	coords["x"] = 0
	coords["y"] = 0
	coords["z"] = 0
	directions := strings.Split(test, ",")
	for _, direction := range directions {
		switch direction {
		case "ne":
			coords["x"]++
		case "n":
			coords["y"]++
		case "nw":
			coords["z"]++
		case "sw":
			coords["x"]--
		case "s":
			coords["y"]--
		case "se":
			coords["z"]--
		default:
			panic(fmt.Sprintf("invalid direction %s", direction))
		}
	}
	return getDistance(coords)
}

func advent11B(test string) int {
	coords := make(map[string]int)
	coords["x"] = 0
	coords["y"] = 0
	coords["z"] = 0
	directions := strings.Split(test, ",")
	max := 0
	for _, direction := range directions {
		switch direction {
		case "ne":
			coords["x"]++
		case "n":
			coords["y"]++
		case "nw":
			coords["z"]++
		case "sw":
			coords["x"]--
		case "s":
			coords["y"]--
		case "se":
			coords["z"]--
		default:
			panic(fmt.Sprintf("invalid direction %s", direction))
		}
		distance := getDistance(coords)
		if distance > max {
			max = distance
		}
	}
	fmt.Printf("%v\n", coords)
	return max
}
