package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func part1(nums [][]int, op []string) int {
	total := 0
	for j := 0; j < len(nums[0]); j++ {
		sub := 0
		for i := 0; i < len(nums); i++ {
			switch op[j] {
			case "+":
				sub += nums[i][j]
			case "*":
				if sub == 0 {
					sub = 1
				}
				sub *= nums[i][j]
			}
		}
		total += sub
	}
	return total
}

func part2(input []string) int {
	var pos []int
	total := 0

	for i := 0; i < len(input[len(input)-1]); i++ {
		if input[len(input)-1][i] != ' ' {
			pos = append(pos, i)
		}
	}
	for i := 0; i < len(pos); i++ {
		end := len(input[0])
		if i != len(pos)-1 {
			end = pos[i+1] - 1
		}
		total += num(input, pos[i], end, rune(input[len(input)-1][pos[i]]))
	}
	return total
}

func num(input []string, start, end int, op rune) int {
	var total int
	switch op {
	case '+':
		total = 0
	case '*':
		total = 1
	}
	for j := end - 1; j >= start; j-- {
		val := 0
		pow := 0
		for i := len(input) - 2; i >= 0; i-- {
			if input[i][j] == ' ' {
				continue
			}
			aux, _ := strconv.Atoi(string(input[i][j]))
			aux = aux * int(math.Pow10(pow))
			pow++
			val += aux
		}
		switch op {
		case '+':
			total += val
		case '*':
			total *= val
		}
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

	scanner := bufio.NewScanner(file)
	var input []string
	for scanner.Scan() {
		line := scanner.Text()
		input = append(input, line)
	}
	var nums [][]int
	var ops []string
	for i := 0; i < len(input)-1; i++ {
		row := strings.Fields(input[i])
		var aux []int
		for j := 0; j < len(row); j++ {
			val, _ := strconv.Atoi(row[j])
			aux = append(aux, val)
		}
		nums = append(nums, aux)
	}
	ops = strings.Fields(input[len(input)-1])

	fmt.Println("Total 1:", part1(nums, ops))
	fmt.Println("Total 2:", part2(input))
}
