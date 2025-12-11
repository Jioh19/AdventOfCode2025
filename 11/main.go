package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Node struct {
	id       string
	children []*Node
}

func (n *Node) Print() {
	if n == nil {
		return
	}
	fmt.Println(n.id)
	for _, child := range n.children {
		fmt.Print(child.id)
	}
}

func part1(you *Node) int {
	return travel(you)
}

func travel(n *Node) int {
	if n == nil {
		return 0
	}
	if n.id == "out" {
		return 1
	}
	sum := 0
	for _, child := range n.children {
		sum += travel(child)
	}
	return sum
}

func part2(svr *Node) int {
	m := make(map[string]int)
	return path(svr, false, false, m)
}

func path(n *Node, fft, dac bool, m map[string]int) int {
	if n == nil {
		return 0
	}
	if n.id == "out" {
		if fft && dac {
			return 1
		} else {
			return 0
		}
	}
	if n.id == "fft" {
		fft = true
	}
	if n.id == "dac" {
		dac = true
	}

	key := fmt.Sprintf("%s-%t-%t", n.id, fft, dac)
	if val, exists := m[key]; exists {
		return val
	}

	sum := 0
	for _, child := range n.children {
		sum += path(child, fft, dac, m)
	}

	m[key] = sum
	return sum
}
func initNode(input []string, start string) *Node {

	m := make(map[string]*Node)
	for _, line := range input {
		parts := strings.Split(line, " ")
		parts[0] = strings.TrimSuffix(parts[0], ":")
		if _, exists := m[parts[0]]; !exists {
			m[parts[0]] = &Node{id: parts[0]}
		}
		parent := m[parts[0]]
		for _, childName := range parts[1:] {
			if _, exists := m[childName]; !exists {
				m[childName] = &Node{id: childName}
			}
			child := m[childName]
			parent.children = append(parent.children, child)
		}
	}
	return m[start]
}
func main() {
	filename := "input.txt"
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var input []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	fmt.Println("Part 1:", part1(initNode(input, "you")))
	fmt.Println("Part 2:", part2(initNode(input, "svr")))
}
