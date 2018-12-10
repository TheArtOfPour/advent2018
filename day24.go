package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

type component struct {
	l int
	r int
}

func splice(i int, components []component) []component {
	var remaining []component
	for j, component := range components {
		if i != j {
			remaining = append(remaining, component)
		}
	}
	return remaining
}

func bestBridge(seed component, remaining []component) int {
	if len(remaining) == 0 {
		return seed.l + seed.r
	}
	// check for branches
	max := 0
	for i, branch := range remaining {
		if seed.r == branch.l {
			best := bestBridge(branch, splice(i, remaining))
			if best > max {
				max = best
			}
		} else if seed.r == branch.r {
			temp := branch.l
			branch.l = branch.r
			branch.r = temp
			best := bestBridge(branch, splice(i, remaining))
			if best > max {
				max = best
			}
		}
	}
	return seed.l + seed.r + max
}

type performance struct {
	length   int
	strength int
}

func longestBridge(seed component, remaining []component, length int) performance {
	length++
	if len(remaining) == 0 {
		perf := performance{length, seed.l + seed.r}
		return perf
	}
	// check for branches
	maxPerf := performance{length, 0}
	for i, branch := range remaining {
		if seed.r == branch.l {
			perf := longestBridge(branch, splice(i, remaining), length)
			if perf.length > maxPerf.length {
				maxPerf = perf
			} else if perf.length == maxPerf.length && perf.strength > maxPerf.strength {
				maxPerf = perf
			}
		} else if seed.r == branch.r {
			temp := branch.l
			branch.l = branch.r
			branch.r = temp
			perf := longestBridge(branch, splice(i, remaining), length)
			if perf.length > maxPerf.length {
				maxPerf = perf
			} else if perf.length == maxPerf.length && perf.strength > maxPerf.strength {
				maxPerf = perf
			}
		}
	}
	maxPerf.strength += seed.l + seed.r
	return maxPerf
}

func getComponents(input string) []component {
	var components []component
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		var c component
		s := scanner.Text()
		parts := strings.Split(s, "/")
		c.l, _ = strconv.Atoi(parts[0])
		c.r, _ = strconv.Atoi(parts[1])
		components = append(components, c)
	}
	return components
}

func advent24A(test string) int {
	components := getComponents(test)
	max := 0
	for i, component := range components {
		if component.l != 0 && component.r != 0 {
			continue
		}
		if component.r == 0 {
			temp := component.l
			component.l = component.r
			component.r = temp
		}
		remaining := splice(i, components)
		best := bestBridge(component, remaining)
		if best > max {
			max = best
		}
	}
	return max
}

func advent24B(test string) int {
	components := getComponents(test)
	fmt.Printf("%v\n", components)
	maxPerf := performance{0, 0}
	for i, component := range components {
		if component.l != 0 && component.r != 0 {
			continue
		}
		if component.r == 0 {
			temp := component.l
			component.l = component.r
			component.r = temp
		}
		// fmt.Printf("i: %d, comp: %v\n", i, component)
		remaining := splice(i, components)
		perf := longestBridge(component, remaining, 0)
		fmt.Printf("perf: %v\n", perf)
		if perf.length > maxPerf.length {
			maxPerf = perf
		} else if perf.length == maxPerf.length && perf.strength > maxPerf.strength {
			maxPerf = perf
		}
	}
	return maxPerf.strength
}
