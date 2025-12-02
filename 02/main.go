package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func palin(num int) bool {
	str := strconv.Itoa(num)
	len := len(str)

	if len%2 != 0 {
		return false
	}
	for i := 0; i < len/2; i++ {
		if str[i] != str[len/2+i] {
			return false
		}
	}
	return true
}

func rep(num int) bool {
	str := strconv.Itoa(num)
	len := len(str)

	for pLen := 1; pLen <= len/2; pLen++ {
		if len%pLen != 0 {
			continue
		}
		pattern := str[:pLen]
		rep := true
		for i := pLen; i < len; i += pLen {
			if str[i:i+pLen] != pattern {
				rep = false
				break
			}
		}
		if rep && len/pLen >= 2 {
			return true
		}
	}
	return false
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
		ids := strings.SplitSeq(line, ",")

		for id := range ids {
			id1, _ := strconv.Atoi(strings.Split(id, "-")[0])
			id2, _ := strconv.Atoi(strings.Split(id, "-")[1])
			for i := id1; i <= id2; i++ {
				if palin(i) {
					total1 += i
				}
				if rep(i) {
					total2 += i
				}
			}
		}
	}
	fmt.Println("Part 1: ", total1)
	fmt.Println("Part 2: ", total2)
}
