package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type layer struct {
	pos  int
	size int
	up   bool
}

func scanFirewall(firewall map[int]layer) {
	for key, layer := range firewall {
		if layer.up && layer.pos == 0 {
			layer.up = false
		} else if !layer.up && layer.pos == layer.size-1 {
			layer.up = true
		}
		if layer.up {
			layer.pos--
		} else {
			layer.pos++
		}
		firewall[key] = layer
	}
}

func drawFirewall(firewall map[int]layer) {
	fmt.Printf("\n")
	for key, layer := range firewall {
		fmt.Printf("\n%d | ", key)
		for i := 0; i < layer.size; i++ {
			if i == layer.pos {
				fmt.Printf("#")
			} else {
				fmt.Printf(".")
			}
		}
	}
}

func resetFirewall(firewall map[int]layer) {
	for key, layer := range firewall {
		layer.up = false
		layer.pos = 0
		firewall[key] = layer
	}
}

func advent13A(test string) int {
	firewall := make(map[int]layer)
	scanner := bufio.NewScanner(strings.NewReader(test))
	severity := 0
	maxDepth := 0
	for scanner.Scan() {
		s := scanner.Text()
		stringParts := strings.Split(s, ": ")
		depth, _ := strconv.Atoi(stringParts[0])
		size, _ := strconv.Atoi(stringParts[1])
		firewall[depth] = layer{pos: 0, size: size}
		maxDepth = depth
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	for pico := 0; pico <= maxDepth; pico++ {
		scan, ok := firewall[pico]
		scanFirewall(firewall)
		if !ok || scan.pos > 0 {
			continue
		}
		severity += scan.size * pico
	}
	return severity
}

func advent13B(test string) int {
	firewall := make(map[int]layer)
	scanner := bufio.NewScanner(strings.NewReader(test))
	maxDepth := 0
	for scanner.Scan() {
		s := scanner.Text()
		stringParts := strings.Split(s, ": ")
		depth, _ := strconv.Atoi(stringParts[0])
		size, _ := strconv.Atoi(stringParts[1])
		firewall[depth] = layer{pos: 0, size: size}
		maxDepth = depth
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	delay := 0
	for {
		caught := false
		resetFirewall(firewall)
		for pico := 0; pico <= maxDepth; pico++ {
			scan, ok := firewall[pico]
			if !ok || (delay+pico)%((scan.size-1)*2) != 0 {
				continue
			}
			caught = true
			break
		}
		if !caught {
			break
		}
		delay++
	}
	return delay
}
