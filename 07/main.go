package main

import (
	"bufio"
	"fmt"
	"os"
)

type Point struct {
	X, Y int
}

func part1(input []string) int {
	set := make(map[Point]struct{})
	for i := range input {
		move(input, i, set)
	}
	return len(set)
}

func move(input []string, line int, set map[Point]struct{}) {
	if line == len(input)-1 {
		return
	}
	for i, c := range input[line] {
		switch c {
		case 'S':
			input[line+1] = input[line+1][:i] + "|" + input[line+1][i+1:]
		case '|':
			switch input[line+1][i] {
			case '.':
				input[line+1] = input[line+1][:i] + "|" + input[line+1][i+1:]
			case '^':
				input[line+1] = input[line+1][:i-1] + "|^|" + input[line+1][i+2:]
				set[Point{line + 1, i}] = struct{}{}
			}
		}
	}
}

func part2(input []string) int {
	memo := make(map[Point]int)

	startPos := 0
	for i, c := range input[0] {
		if c == 'S' {
			startPos = i
			break
		}
	}
	result := move2(input, 0, startPos, memo)

	return result
}

func move2(input []string, line int, pos int, memo map[Point]int) int {
	if line == len(input)-1 {
		return 1
	}

	key := Point{line, pos}
	if val, exist := memo[key]; exist {
		return val
	}

	line++
	result := 0

	switch input[line][pos] {
	case '.', '|':
		result = move2(input, line, pos, memo)
	case '^':
		result += move2(input, line, pos-1, memo)
		result += move2(input, line, pos+1, memo)
	}

	memo[key] = result
	return result
}

func main() {
	filename := "input.txt"
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var input []string
	for scanner.Scan() {
		line := scanner.Text()
		input = append(input, line)
	}
	fmt.Println("Part 1:", part1(input))
	fmt.Println("Part 2:", part2(input))
}
