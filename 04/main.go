package main

import (
	"bufio"
	"fmt"
	"os"
)

func check(m []string, i, j int) bool {
	roll := 4
	for x := -1; x <= 1; x++ {
		for y := -1; y <= 1; y++ {
			if i+x < 0 || i+x >= len(m) || j+y < 0 || j+y >= len(m[i]) || (x == 0 && y == 0) {
				continue
			}
			if string(m[i+x][j+y]) == "@" {
				roll--
			}
		}
		if roll <= 0 {
			return false
		}
	}
	return true
}

func take(m []string, i, j int) bool {
	roll := 4
	for x := -1; x <= 1; x++ {
		for y := -1; y <= 1; y++ {
			if i+x < 0 || i+x >= len(m) || j+y < 0 || j+y >= len(m[i]) || (x == 0 && y == 0) {
				continue
			}
			if string(m[i+x][j+y]) == "@" {
				roll--
			}
		}
		if roll <= 0 {
			return false
		}
	}
	m[i] = m[i][:j] + "." + m[i][j+1:]
	return true
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
	var m []string
	for scanner.Scan() {
		line := scanner.Text()
		m = append(m, line)
	}

	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[i]); j++ {
			if string(m[i][j]) == "@" && check(m, i, j) {
				total1++
			}
		}
	}
	fmt.Println("Total 1:", total1)

	for {
		aux := 0
		for i := 0; i < len(m); i++ {
			for j := 0; j < len(m[i]); j++ {
				if string(m[i][j]) == "@" && take(m, i, j) {
					aux++
				}
			}
		}
		if aux == 0 {
			break
		}
		total2 += aux
	}
	fmt.Println("Total 2:", total2)
}
