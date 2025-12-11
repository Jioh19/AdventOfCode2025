package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Mach struct {
	light int
	wire  []int
	jolt  string
}

type state struct {
	current int
	steps   int
	path    []int
}

func strToBin(s string) int {
	b := 0
	strNums := strings.Split(s, ",")

	for _, strNum := range strNums {
		num, _ := strconv.Atoi(strNum)
		b = b | 1<<num
	}

	return b
}

func lightToBin(s string) int {
	var sb strings.Builder

	for i := len(s) - 1; i >= 0; i-- {
		c := s[i]
		switch c {
		case '.':
			sb.WriteString("0")
		case '#':
			sb.WriteString("1")
		}
	}
	aux, _ := strconv.ParseInt(sb.String(), 2, 64)

	return int(aux)
}

func shortXor(target int, available []int) int {
	if target == 0 {
		return 0
	}

	if slices.Contains(available, target) {
		return 1
	}

	queue := []state{{current: 0, steps: 0, path: []int{}}}
	visited := make(map[int]int)
	visited[0] = 0
	maxSteps := 20

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		if curr.steps >= maxSteps {
			continue
		}

		for _, val := range available {
			next := curr.current ^ val
			newSteps := curr.steps + 1

			if next == target {
				finalPath := make([]int, len(curr.path))
				copy(finalPath, curr.path)
				finalPath = append(finalPath, val)
				return newSteps
			}

			if prevSteps, seen := visited[next]; !seen || newSteps <= prevSteps {
				visited[next] = newSteps
				newPath := make([]int, len(curr.path))
				copy(newPath, curr.path)
				newPath = append(newPath, val)
				queue = append(queue, state{
					current: next,
					steps:   newSteps,
					path:    newPath,
				})
			}
		}
	}

	return 0
}

func part1(machs []Mach) int {
	total := 0

	for _, mach := range machs {
		steps := shortXor(mach.light, mach.wire)
		total += steps
	}
	return total
}

func main() {
	filename := "input.txt"
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var machs []Mach
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		var mach Mach

		lightStr := strings.Replace(line[0], "[", "", -1)
		lightStr = strings.Replace(lightStr, "]", "", -1)
		mach.light = lightToBin(lightStr)
		for _, numStr := range line[1 : len(line)-1] {
			numStr = strings.Replace(numStr, "(", "", -1)
			numStr = strings.Replace(numStr, ")", "", -1)
			mach.wire = append(mach.wire, strToBin(numStr))
		}
		joltStr := strings.Replace(line[len(line)-1], "{", "", -1)
		joltStr = strings.Replace(joltStr, "}", "", -1)
		mach.jolt = strings.TrimSpace(joltStr)
		machs = append(machs, mach)
	}

	fmt.Println("Part 1:", part1(machs))
	//
	// fmt.Println("Part 2:", part2(points))
}
