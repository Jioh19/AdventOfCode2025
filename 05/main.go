package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func part1(rans [2][]int, ids []int) int {
	fresh := 0
	for _, id := range ids {
	IdLoop:
		for i := 0; i < len(rans[0]); i++ {
			if id >= rans[0][i] && id <= rans[1][i] {
				fresh++
				break IdLoop
			}
		}
	}
	return fresh
}

func part2(rans [2][]int) int {
	fresh := 0
	merged := mergeAll(rans)
	for i := 0; i < len(merged[0]); i++ {
		fresh += merged[1][i] - merged[0][i] + 1
	}
	return fresh
}

func merge(rans [2][]int) [2][]int {
	var merged [2][]int
	for i := 0; i < len(rans[0]); i++ {
		start := rans[0][i]
		end := rans[1][i]
		mergedFlag := false
		for j := 0; j < len(merged[0]); j++ {
			if (start >= merged[0][j] && start <= merged[1][j]) || (end >= merged[0][j] && end <= merged[1][j]) || (start <= merged[0][j] && end >= merged[1][j]) {
				if start < merged[0][j] {
					merged[0][j] = start
				}
				if end > merged[1][j] {
					merged[1][j] = end
				}
				mergedFlag = true
				break
			}
		}
		if !mergedFlag {
			merged[0] = append(merged[0], start)
			merged[1] = append(merged[1], end)
		}
	}
	return merged
}

func mergeAll(rans [2][]int) [2][]int {
	for {
		merged := merge(rans)
		if len(merged[0]) == len(rans[0]) {
			return merged
		}
		rans = merged
	}
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
	var rans [2][]int
	var ids []int

	check := false
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			check = true
			continue
		}
		if !check {
			ran := strings.Split(line, "-")
			start, _ := strconv.Atoi(ran[0])
			end, _ := strconv.Atoi(ran[1])
			rans[0] = append(rans[0], start)
			rans[1] = append(rans[1], end)
		} else {
			id, _ := strconv.Atoi(line)
			ids = append(ids, id)
		}
	}
	fmt.Println("Total 1:", part1(rans, ids))
	fmt.Println("Total 2:", part2(rans))
}
