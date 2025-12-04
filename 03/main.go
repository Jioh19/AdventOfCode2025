package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func big(line string, keep int) int {
	n := len(line)
	res := ""
	startId := 0

	for i := 0; i < keep; i++ {
		maxNum := line[startId]
		maxId := startId

		endIdx := n - (keep - i) + 1
		for j := startId; j < endIdx; j++ {
			if line[j] > maxNum {
				maxNum = line[j]
				maxId = j
			}
		}
		res += string(maxNum)
		startId = maxId + 1
	}

	num, _ := strconv.Atoi(res)
	return num
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
	total1 := 0
	total2 := 0

	for scanner.Scan() {
		line := scanner.Text()
		first := line[0]
		index := 0
		for i := 1; i < len(line)-1; i++ {
			if first < line[i] {
				first = line[i]
				index = i
			}
		}
		second := line[index+1]
		for i := index + 1; i < len(line); i++ {
			if second < line[i] {
				second = line[i]
			}
		}
		num, _ := strconv.Atoi(string(first) + string(second))
		total1 += num
		res := big(line, 12)
		total2 += res
	}

	fmt.Println("Part 1:", total1)
	fmt.Println("Part 2:", total2)
}
