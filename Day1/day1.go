package main

import (
	"strconv"
	"strings"
)

func day1ExA(input string) int {
	lines := strings.Split(input, "\n")
	if len(lines) < 2 {
		return 0
	}
	prev, _ := strconv.Atoi(lines[0])
	result := 0
	for i := 1; i < len(lines); i++ {
		curr, _ := strconv.Atoi(lines[i])
		if curr > prev {
			result++
		}
		prev = curr
	}
	return result
}

func day1ExB(input string) int {
	windowSize := 3
	lines := strings.Split(input, "\n")
	if len(lines) < windowSize*2 {
		return 0
	}
	result := 0
	prevWindowSum := 0
	p := 0
	for p < windowSize {
		num, _ := strconv.Atoi(lines[p])
		prevWindowSum += num
		p++
	}
	for p < len(lines) {
		currNum, _ := strconv.Atoi(lines[p])
		startNum, _ := strconv.Atoi(lines[p-windowSize])
		currWindowSum := prevWindowSum - startNum + currNum
		if currWindowSum > prevWindowSum {
			result++
		}
		prevWindowSum = currWindowSum
		p++
	}
	return result
}
