package main

import (
	"strconv"
	"strings"
)

func advent8A(test string) int {
	// push pop off stack
	s := make([]int, 0)
	nodes := strings.Split(test, " ")
	// mode := 0 // 0 node, 1 metadata
	s = append(s, 0)

	skip := false
	total := 0
	for i := range nodes {
		if skip {
			skip = false
			continue
		}
		l := len(s)
		operation := s[l-1]
		s = s[:l-1]
		//fmt.Printf("op %d, stack %v\n", operation, s)

		if operation == 0 {
			numberOfChildren, _ := strconv.Atoi(nodes[i])
			numberOfMetadata, _ := strconv.Atoi(nodes[i+1])
			//fmt.Printf("#C %d, #M %d\n", numberOfChildren, numberOfMetadata)
			for j := 0; j < numberOfMetadata; j++ {
				s = append(s, 1)
			}
			for j := 0; j < numberOfChildren; j++ {
				s = append(s, 0)
			}
			skip = true
		}
		if operation == 1 {
			metadata, _ := strconv.Atoi(nodes[i])
			//fmt.Printf("M %d\n", metadata)
			total += metadata
		}
	}
	return total
}

type stackEntry struct {
	operation int
	node      int
}

type nodeEntry struct {
	metadata []int
	children []int
}

func advent8B(test string) int {
	// push pop off stack
	s := make([]stackEntry, 0)
	nodes := strings.Split(test, " ")
	// mode := 0 // 0 node, 1 metadata
	se := stackEntry{operation: 0, node: 0}
	s = append(s, se)

	nodeEntries := make(map[int]nodeEntry)

	skip := false
	total := 0
	nodeNum := 0
	for i := range nodes {
		if skip {
			skip = false
			continue
		}
		l := len(s)
		se := s[l-1]
		s = s[:l-1]

		if se.operation == 0 {
			if se.node != 0 {
				if _, ok := nodeEntries[se.node]; !ok {
					nodeEntries[se.node] = nodeEntry{metadata: make([]int, 0), children: make([]int, 0)}
				}
				//nodeEntries[se.node].children = append(nodeEntries[se.node].children, nodeNum)
			}
			numberOfChildren, _ := strconv.Atoi(nodes[i])
			numberOfMetadata, _ := strconv.Atoi(nodes[i+1])
			// add node shell
			ne := nodeEntry{metadata: make([]int, 0), children: make([]int, 0)}
			nodeEntries[nodeNum] = ne
			for j := 0; j < numberOfMetadata; j++ {
				entry := stackEntry{operation: 1, node: nodeNum}
				s = append(s, entry)
			}
			for j := 0; j < numberOfChildren; j++ {
				entry := stackEntry{operation: 0, node: nodeNum}
				s = append(s, entry)
			}
			skip = true
			nodeNum++
		}
		if se.operation == 1 {
			metadata, _ := strconv.Atoi(nodes[i])
			total += metadata
		}
	}
	return total
}
