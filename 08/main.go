package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
)

type CalcRes struct {
	Index int
	Data  []int
}

type CalcData struct {
	B1, B2, Dist int
}
type Pair struct {
	V, C int
}

func part1(input []string) int {
	res := make(chan CalcRes, len(input))
	conn := make(map[int][]int)
	var listConn []CalcData
	mapConn := make(map[int]int)

	var wg sync.WaitGroup

	for i := range input {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			d := calc(input, index)
			res <- CalcRes{Index: index, Data: d}
		}(i)
	}

	go func() {
		wg.Wait()
		close(res)
	}()

	for result := range res {
		conn[result.Index] = result.Data
	}

	for i := 0; i < len(conn); i++ {
		for j := i + 1; j < len(conn); j++ {
			dist := conn[i][j]
			listConn = append(listConn, CalcData{B1: i, B2: j, Dist: dist})
		}
	}

	sort.Slice(listConn, func(i, j int) bool {
		return listConn[i].Dist < listConn[j].Dist
	})

	for i, v := range listConn[:1000] {

		v1, ex1 := mapConn[v.B1]
		v2, ex2 := mapConn[v.B2]

		if !ex1 && !ex2 {
			mapConn[v.B1] = i
			mapConn[v.B2] = i
		} else if ex1 && ex2 {
			if v1 != v2 {
				rMapV(mapConn, v1, v2)
			}
		} else if ex1 {
			mapConn[v.B2] = v1
		} else {
			mapConn[v.B1] = v2
		}
	}

	pairs := countMapV(mapConn)

	return pairs[0].C * pairs[1].C * pairs[2].C
}
func rMapV(m map[int]int, v1, v2 int) {
	for key, value := range m {
		if value == v1 {
			m[key] = v2
		}
	}
}

func countMapV(m map[int]int) []Pair {
	count := make(map[int]int)

	for _, v := range m {
		count[v]++
	}
	var pairs []Pair
	for value, cnt := range count {
		pairs = append(pairs, Pair{V: value, C: cnt})
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].C > pairs[j].C
	})

	return pairs
}

func pow(base, exp int) int {
	result := 1
	for i := 0; i < exp; i++ {
		result *= base
	}
	return result
}

func calc(input []string, i int) []int {
	var d []int
	pos := strings.Split(input[i], ",")
	posX, _ := strconv.Atoi(pos[0])
	posY, _ := strconv.Atoi(pos[1])
	posZ, _ := strconv.Atoi(pos[2])
	for _, line := range input {
		pos2 := strings.Split(line, ",")
		pos2X, _ := strconv.Atoi(pos2[0])
		pos2Y, _ := strconv.Atoi(pos2[1])
		pos2Z, _ := strconv.Atoi(pos2[2])
		d = append(d, pow((posX-pos2X), 2)+pow((posY-pos2Y), 2)+pow((posZ-pos2Z), 2))
	}
	return d
}

func part2(input []string) int {
	res := make(chan CalcRes, len(input))
	conn := make(map[int][]int)
	var listConn []CalcData
	mapConn := make(map[int]int)

	var wg sync.WaitGroup

	for i := range input {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			d := calc(input, index)
			res <- CalcRes{Index: index, Data: d}
		}(i)
	}

	go func() {
		wg.Wait()
		close(res)
	}()

	for result := range res {
		conn[result.Index] = result.Data
	}

	for i := 0; i < len(conn); i++ {
		for j := i + 1; j < len(conn); j++ {
			dist := conn[i][j]
			listConn = append(listConn, CalcData{B1: i, B2: j, Dist: dist})
		}
	}

	sort.Slice(listConn, func(i, j int) bool {
		return listConn[i].Dist < listConn[j].Dist
	})

	lastB1, lastB2 := -1, -1

	for _, v := range listConn {
		v1, ex1 := mapConn[v.B1]
		v2, ex2 := mapConn[v.B2]

		connMade := false

		if !ex1 && !ex2 {
			mapConn[v.B1] = v.B1
			mapConn[v.B2] = v.B1
			connMade = true
		} else if ex1 && ex2 {
			if v1 != v2 {
				rMapV(mapConn, v1, v2)
				connMade = true
			}
		} else if ex1 {
			mapConn[v.B2] = v1
			connMade = true
		} else {
			mapConn[v.B1] = v2
			connMade = true
		}

		if connMade {
			lastB1 = v.B1
			lastB2 = v.B2

			groups := countMapV(mapConn)
			if len(groups) == 1 && len(mapConn) == len(input) {
				break
			}
		}
	}

	pos1 := strings.Split(input[lastB1], ",")
	x1, _ := strconv.Atoi(pos1[0])

	pos2 := strings.Split(input[lastB2], ",")
	x2, _ := strconv.Atoi(pos2[0])

	return x1 * x2
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
