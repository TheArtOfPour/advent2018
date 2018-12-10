package main

func advent17A(test int) int {
	var spinlock []int
	currentIndex := 0
	spinlock = append(spinlock, 0)
	for i := 1; i <= 2017; i++ {
		currentIndex += test
		for currentIndex > len(spinlock) {
			currentIndex -= len(spinlock)
		}
		var temp []int
		for _, j := range spinlock[:currentIndex] {
			temp = append(temp, j)
		}
		temp = append(temp, i)
		for _, j := range spinlock[currentIndex:] {
			temp = append(temp, j)
		}
		//fmt.Printf("%v + %d + %v = %v\n", spinlock[:currentIndex], i, spinlock[currentIndex:], temp)
		spinlock = temp
		//fmt.Printf("%d @ %d -> %v\n", i, currentIndex, spinlock)
		currentIndex++
		for currentIndex > len(spinlock) {
			currentIndex -= len(spinlock)
		}
	}
	return spinlock[currentIndex]
}

func advent17B(test int) int {
	currentIndex := 0
	result := 0
	for i := 1; i <= 50000000; i++ {
		currentIndex += test + 1
		for currentIndex > i {
			currentIndex -= i
		}
		if currentIndex == 1 {
			result = i
		}
	}
	return result
}
