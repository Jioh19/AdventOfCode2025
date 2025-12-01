package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func posMod(value, mod int) int {
	return ((value % mod) + mod) % mod
}

func main() {
	filename := "input.txt"
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var numbers []int
	var abs []int
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}
		direction := line[0]
		numStr := line[1:]
		num, err := strconv.Atoi(numStr)
		if err != nil {
			fmt.Println("Error parsing number:", err)
			continue
		}
		if direction == 'L' {
			numbers = append(numbers, -num)
		} else {
			numbers = append(numbers, num)
		}
		abs = append(abs, num)
	}

	dial := 50
	count := 0
	for _, number := range numbers {
		dial += number
		dial %= 100
		if dial == 0 {
			count++
		}
	}
	fmt.Println("Part 1: ", count)

	dial = 50
	count = 0
	for _, num := range numbers {
		pos := dial
		if num > 0 {
			if pos == 0 {
				count += num / 100
			} else {
				d := 100 - pos
				if num >= d {
					count++
					count += (num - d) / 100
				}
			}
			dial = (pos + num) % 100
		} else {
			n := -num
			if pos == 0 {
				count += n / 100
			} else {
				if n >= pos {
					count++
					count += (n - pos) / 100
				}
			}
			dial = posMod(pos-n, 100)
		}
	}
	fmt.Println("Part 2: ", count)
}
