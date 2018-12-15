package main

import (
	"bytes"
	"fmt"
	"strings"
)

// uni-directional graph vertex
type vertex struct {
	name     rune
	children []vertex
}

type instruct struct {
	reqStep rune
	step    rune
}

func buildSleighGraph(s string) vertex {
	instructionsRaw := strings.Split(s, "\n")
	var instructions []instruct
	reqSteps := make(map[rune]bool, len(instructionsRaw))
	steps := make(map[rune]bool, len(instructionsRaw))
	for _, instruction := range instructionsRaw {
		instructionParts := strings.Split(instruction, " ")
		reqStep := []rune(instructionParts[1])[0]
		reqSteps[reqStep] = true
		step := []rune(instructionParts[7])[0]
		steps[step] = true
		instructions = append(instructions, instruct{reqStep: reqStep, step: step})
	}
	head := ' '
	for reqStep := range reqSteps {
		found := false
		for step := range steps {
			if reqStep == step {
				found = true
				break
			}
		}
		if !found {
			head = reqStep
			break
		}
	}

	var emptyChildren []vertex
	headVertex := vertex{name: head, children: emptyChildren}
	headVertex = buildSleighGraphRecursive(headVertex, instructions)
	return headVertex
}

func buildSleighGraphRecursive(v vertex, instructions []instruct) vertex {
	found := false
	for _, instruction := range instructions {
		if instruction.reqStep == v.name {
			found = true
			var emptyChildren []vertex
			newVertex := vertex{name: instruction.step, children: emptyChildren}
			v.children = append(v.children, buildSleighGraphRecursive(newVertex, instructions))
		}
	}
	if !found {
		var baseCase []vertex
		return vertex{name: v.name, children: baseCase}
	}
	return v
}

func readGraph(v vertex) string {
	var newString bytes.Buffer
	var availableVertices []vertex
	availableVertices = append(availableVertices, v)
	for len(availableVertices) > 0 {
		//fmt.Printf("%v\n", availableVertices)
		var emptyChildren []vertex
		leastVertex := vertex{name: ' ', children: emptyChildren}
		leastVertexIndex := 0
		for i, potentialVertex := range availableVertices {
			tmp := make([]vertex, len(availableVertices))
			copy(tmp, availableVertices)
			excluded := append(tmp[:i], tmp[i+1:]...)
			if i == len(tmp)-1 {
				excluded = tmp[:i]
			}
			if potentialVertex.name < leastVertex.name || leastVertex.name == ' ' {
				if reqsMet(potentialVertex, excluded) {
					leastVertex = potentialVertex
					leastVertexIndex = i
				}
			}
			fmt.Printf("%v, %v \n", potentialVertex, reqsMet(potentialVertex, excluded))
		}
		//fmt.Printf("\n")
		if leastVertex.name != ' ' {
			newString.WriteString(string(leastVertex.name))
		}
		availableVertices[leastVertexIndex] = availableVertices[len(availableVertices)-1]
		availableVertices = availableVertices[:len(availableVertices)-1]
		availableVertices = append(availableVertices, leastVertex.children...)
		//fmt.Printf("%v\n\n\n", availableVertices)
	}

	return newString.String()
}

func reqsMet(v vertex, vertices []vertex) bool {
	found := true
	for _, vertexCheck := range vertices {
		if v.name == vertexCheck.name {
			return false
		}
		found = found && reqsMet(v, vertexCheck.children)
	}
	return found
}

func getStepPairs(s string) []instruct {
	instructionsRaw := strings.Split(s, "\n")
	var instructions []instruct
	reqSteps := make(map[rune]bool, len(instructionsRaw))
	steps := make(map[rune]bool, len(instructionsRaw))
	for _, instruction := range instructionsRaw {
		instructionParts := strings.Split(instruction, " ")
		reqStep := []rune(instructionParts[1])[0]
		reqSteps[reqStep] = true
		step := []rune(instructionParts[7])[0]
		steps[step] = true
		instructions = append(instructions, instruct{reqStep: reqStep, step: step})
	}
	return instructions
}

func advent7A(test string) (string, error) {
	instructions := getStepPairs(test)
	done := make([]rune, 0)
	lastStep := ' '

	for len(instructions) > 0 {
		steps := make([]rune, 0)
		reqSteps := make([]rune, 0)
		for _, instruction := range instructions {
			steps = append(steps, instruction.step)
			reqSteps = append(reqSteps, instruction.reqStep)
		}

		available := make([]rune, 0)
		// diff req steps to steps
		for _, rs := range reqSteps {
			found := false
			for _, s := range steps {
				lastStep = s
				if rs == s {
					found = true
					break
				}
			}
			if !found {
				// whichever letters aren't in steps and haven't been done are "available"
				alreadyDone := false
				for _, d := range done {
					if rs == d {
						alreadyDone = true
						break
					}
				}
				if !alreadyDone {
					available = append(available, rs)
				}
			}
		}

		// choose from available
		nextStep := ' '
		for _, a := range available {
			if a < nextStep || nextStep == ' ' {
				nextStep = a
			}
		}

		// add to done
		done = append(done, nextStep)
		//fmt.Printf("%s ", string(nextStep))

		// remove all instructions with that required step
		tmp := make([]instruct, 0)
		for _, instruction := range instructions {
			if instruction.reqStep != nextStep {
				tmp = append(tmp, instruction)
			}
		}
		instructions = tmp
	}
	done = append(done, lastStep)

	//fmt.Printf("\n%v\n", done)

	// convert done to string and return
	return string(done), nil
}

func advent7B(test string) (int, error) {

	// implement as heap?

	instructions := getStepPairs(test)
	done := make([]rune, 0)
	lastStep := ' '
	time := 0
	workerStep := ' '
	workerCooldown := 0
	extraWorkerStep := ' '
	extraWorkerCooldown := 0
	extraWorker2Step := ' '
	extraWorker2Cooldown := 0
	extraWorker3Step := ' '
	extraWorker3Cooldown := 0
	extraWorker4Step := ' '
	extraWorker4Cooldown := 0
	// 65 = A (r-64 = time) or int(r - '0')

	for len(instructions) > 0 {
		steps := make([]rune, 0)
		reqSteps := make([]rune, 0)
		for _, instruction := range instructions {
			steps = append(steps, instruction.step)
			reqSteps = append(reqSteps, instruction.reqStep)
		}

		available := make([]rune, 0)
		// diff req steps to steps
		for _, rs := range reqSteps {
			found := false
			for _, s := range steps {
				lastStep = s
				if rs == s {
					found = true
					break
				}
			}
			if !found && rs != workerStep &&
				rs != extraWorkerStep &&
				rs != extraWorker2Step &&
				rs != extraWorker3Step &&
				rs != extraWorker4Step {
				// whichever letters aren't in steps and haven't been done are "available"
				alreadyDone := false
				for _, d := range done {
					if rs == d {
						alreadyDone = true
						break
					}
				}
				if !alreadyDone {
					available = append(available, rs)
				}
			}
		}

		// jobs done
		if workerCooldown > 0 {
			workerCooldown--
			if workerCooldown == 0 {
				// add to done
				done = append(done, workerStep)
				// remove from instructions
				tmp := make([]instruct, 0)
				for _, instruction := range instructions {
					if instruction.reqStep != workerStep {
						tmp = append(tmp, instruction)
					}
				}
				instructions = tmp
			}
		}

		// more work?
		if workerCooldown == 0 && len(available) > 0 {
			workerStep = ' '
			for _, a := range available {
				if a < workerStep || workerStep == ' ' {
					workerStep = a
					workerCooldown = int(a-'0') - 16 + 60
				}
			}
			// if workerCooldown > 0 {
			// 	workerCooldown--
			// 	if workerCooldown == 0 {
			// 		// add to done
			// 		done = append(done, workerStep)
			// 		// remove from instructions
			// 		tmp := make([]instruct, 0)
			// 		for _, instruction := range instructions {
			// 			if instruction.reqStep != workerStep {
			// 				tmp = append(tmp, instruction)
			// 			}
			// 		}
			// 		instructions = tmp
			// 	}
			// }
		}

		// extra worker done
		if extraWorkerCooldown > 0 {
			extraWorkerCooldown--
			if extraWorkerCooldown == 0 {
				// add to done
				done = append(done, extraWorkerStep)
				// remove from instructions
				tmp := make([]instruct, 0)
				for _, instruction := range instructions {
					if instruction.reqStep != extraWorkerStep {
						tmp = append(tmp, instruction)
					}
				}
				instructions = tmp
			}
		}

		// check for second worker
		if extraWorkerCooldown == 0 && len(available) > 0 {
			extraWorkerStep = ' '
			for _, a := range available {
				if a == workerStep {
					continue // don't double up the elves
				}
				if a < extraWorkerStep || extraWorkerStep == ' ' {
					extraWorkerStep = a
					extraWorkerCooldown = int(a-'0') - 16 + 60
				}
			}
			// if extraWorkerCooldown > 0 {
			// 	extraWorkerCooldown--
			// 	if extraWorkerCooldown == 0 {
			// 		// add to done
			// 		done = append(done, extraWorkerStep)
			// 		// remove from instructions
			// 		tmp := make([]instruct, 0)
			// 		for _, instruction := range instructions {
			// 			if instruction.reqStep != extraWorkerStep {
			// 				tmp = append(tmp, instruction)
			// 			}
			// 		}
			// 		instructions = tmp
			// 	}
			// }
		}

		// extra worker 3

		// extra worker done
		if extraWorker2Cooldown > 0 {
			extraWorker2Cooldown--
			if extraWorker2Cooldown == 0 {
				// add to done
				done = append(done, extraWorker2Step)
				// remove from instructions
				tmp := make([]instruct, 0)
				for _, instruction := range instructions {
					if instruction.reqStep != extraWorker2Step {
						tmp = append(tmp, instruction)
					}
				}
				instructions = tmp
			}
		}

		// check for second worker
		if extraWorker2Cooldown == 0 && len(available) > 0 {
			extraWorker2Step = ' '
			for _, a := range available {
				if a == workerStep || a == extraWorkerStep {
					continue // don't double up the elves
				}
				if a < extraWorker2Step || extraWorker2Step == ' ' {
					extraWorker2Step = a
					extraWorker2Cooldown = int(a-'0') - 16 + 60
				}
			}
			// if extraWorker2Cooldown > 0 {
			// 	extraWorker2Cooldown--
			// 	if extraWorker2Cooldown == 0 {
			// 		// add to done
			// 		done = append(done, extraWorker2Step)
			// 		// remove from instructions
			// 		tmp := make([]instruct, 0)
			// 		for _, instruction := range instructions {
			// 			if instruction.reqStep != extraWorker2Step {
			// 				tmp = append(tmp, instruction)
			// 			}
			// 		}
			// 		instructions = tmp
			// 	}
			// }
		}

		// extra worker 4

		// extra worker done
		if extraWorker3Cooldown > 0 {
			extraWorker3Cooldown--
			if extraWorker3Cooldown == 0 {
				// add to done
				done = append(done, extraWorker3Step)
				// remove from instructions
				tmp := make([]instruct, 0)
				for _, instruction := range instructions {
					if instruction.reqStep != extraWorker3Step {
						tmp = append(tmp, instruction)
					}
				}
				instructions = tmp
			}
		}

		// check for second worker
		if extraWorker3Cooldown == 0 && len(available) > 0 {
			extraWorker3Step = ' '
			for _, a := range available {
				if a == workerStep || a == extraWorkerStep || a == extraWorker2Step {
					continue // don't double up the elves
				}
				if a < extraWorker3Step || extraWorker3Step == ' ' {
					extraWorker3Step = a
					extraWorker3Cooldown = int(a-'0') - 16 + 60
				}
			}
			// if extraWorker3Cooldown > 0 {
			// 	extraWorker3Cooldown--
			// 	if extraWorker3Cooldown == 0 {
			// 		// add to done
			// 		done = append(done, extraWorker3Step)
			// 		// remove from instructions
			// 		tmp := make([]instruct, 0)
			// 		for _, instruction := range instructions {
			// 			if instruction.reqStep != extraWorker3Step {
			// 				tmp = append(tmp, instruction)
			// 			}
			// 		}
			// 		instructions = tmp
			// 	}
			// }
		}

		// extra worker 5

		// extra worker done
		if extraWorker4Cooldown > 0 {
			extraWorker4Cooldown--
			if extraWorker4Cooldown == 0 {
				// add to done
				done = append(done, extraWorker4Step)
				// remove from instructions
				tmp := make([]instruct, 0)
				for _, instruction := range instructions {
					if instruction.reqStep != extraWorker4Step {
						tmp = append(tmp, instruction)
					}
				}
				instructions = tmp
			}
		}

		// check for second worker
		if extraWorker4Cooldown == 0 && len(available) > 0 {
			extraWorker4Step = ' '
			for _, a := range available {
				if a == workerStep || a == extraWorkerStep || a == extraWorker2Step || a == extraWorker3Step {
					continue // don't double up the elves
				}
				if a < extraWorker4Step || extraWorker4Step == ' ' {
					extraWorker4Step = a
					extraWorker4Cooldown = int(a-'0') - 16 + 60
				}
			}
			// if extraWorker4Cooldown > 0 {
			// 	extraWorker4Cooldown--
			// 	if extraWorker4Cooldown == 0 {
			// 		// add to done
			// 		done = append(done, extraWorker4Step)
			// 		// remove from instructions
			// 		tmp := make([]instruct, 0)
			// 		for _, instruction := range instructions {
			// 			if instruction.reqStep != extraWorker4Step {
			// 				tmp = append(tmp, instruction)
			// 			}
			// 		}
			// 		instructions = tmp
			// 	}
			// }
		}

		//fmt.Printf("%s ", string(nextStep))
		fmt.Printf("%ds W1:%s W2:%s W3:%s W4:%s W5:%s -> %s\n", time, string(workerStep), string(extraWorkerStep), string(extraWorker2Step), string(extraWorker3Step), string(extraWorker4Step), string(done))

		time++
	}
	done = append(done, lastStep)
	time += int(lastStep-'0') - 16 + 60
	time++

	fmt.Printf("\n%ds %s\n", time, string(done))

	return time, nil
}
