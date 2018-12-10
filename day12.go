package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Vertex bi-directional graph vertex
type Vertex struct {
	value int
	edges []*Vertex
}

func buildGraph(root *Vertex, inputs map[int][]int, seen map[int]bool) {
	_, ok := seen[root.value]
	if ok {
		return
	}
	seen[root.value] = true
	for _, linked := range inputs[root.value] {
		var temp Vertex
		temp.value = linked
		root.edges = append(root.edges, &temp)
		temp.edges = append(temp.edges, root)
		buildGraph(&temp, inputs, seen)
	}
}

func countGraphEdges(root *Vertex, seen map[int]bool) int {
	_, ok := seen[root.value]
	if ok {
		return 0
	}
	seen[root.value] = true
	sum := 0
	for _, edge := range root.edges {
		sum += countGraphEdges(edge, seen)
	}
	return sum + 1
}

func advent12A(test string) int {
	inputs := make(map[int][]int)
	scanner := bufio.NewScanner(strings.NewReader(test))
	for scanner.Scan() {
		s := scanner.Text()
		stringParts := strings.Split(s, " <-> ")
		program, _ := strconv.Atoi(stringParts[0])
		linkedParts := strings.Split(stringParts[1], ", ")
		var temp []int
		for _, linked := range linkedParts {
			ilinked, _ := strconv.Atoi(linked)
			temp = append(temp, ilinked)
		}
		inputs[program] = temp
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	var root Vertex
	root.value = 0
	seen := make(map[int]bool)
	buildGraph(&root, inputs, seen)

	seen = make(map[int]bool)
	return countGraphEdges(&root, seen)
}

func advent12B(test string) int {
	inputs := make(map[int][]int)
	scanner := bufio.NewScanner(strings.NewReader(test))
	for scanner.Scan() {
		s := scanner.Text()
		stringParts := strings.Split(s, " <-> ")
		program, _ := strconv.Atoi(stringParts[0])
		linkedParts := strings.Split(stringParts[1], ", ")
		var temp []int
		for _, linked := range linkedParts {
			ilinked, _ := strconv.Atoi(linked)
			temp = append(temp, ilinked)
		}
		inputs[program] = temp
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	groups := 0
	graphs := make(map[*Vertex][]int)
	for input := range inputs {
		haveSeen := false
		for _, graph := range graphs {
			for _, seenValue := range graph {
				if input == seenValue {
					haveSeen = true
					break
				}
			}
			if haveSeen {
				break
			}
		}
		if !haveSeen {
			var root Vertex
			root.value = input
			seen := make(map[int]bool)
			buildGraph(&root, inputs, seen)
			keys := make([]int, 0, len(seen))
			for k := range seen {
				keys = append(keys, k)
			}
			graphs[&root] = keys
			groups++
		}
	}
	return groups
}
