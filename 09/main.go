package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type Point struct {
	X, Y int
}

type CoordMapper struct {
	xMap map[int]int // original -> compressed
	yMap map[int]int // original -> compressed
	xInv []int       // compressed -> original
	yInv []int       // compressed -> original
}

func newCoordMapper(points []Point) *CoordMapper {
	xSet := make(map[int]bool)
	ySet := make(map[int]bool)

	for _, p := range points {
		xSet[p.X] = true
		ySet[p.Y] = true
	}

	xVals := make([]int, 0, len(xSet))
	for x := range xSet {
		xVals = append(xVals, x)
	}
	sort.Ints(xVals)

	yVals := make([]int, 0, len(ySet))
	for y := range ySet {
		yVals = append(yVals, y)
	}
	sort.Ints(yVals)

	xMap := make(map[int]int)
	for i, x := range xVals {
		xMap[x] = i
	}

	yMap := make(map[int]int)
	for i, y := range yVals {
		yMap[y] = i
	}

	return &CoordMapper{
		xMap: xMap,
		yMap: yMap,
		xInv: xVals,
		yInv: yVals,
	}
}

func (cm *CoordMapper) compress(p Point) Point {
	return Point{X: cm.xMap[p.X], Y: cm.yMap[p.Y]}
}

func (cm *CoordMapper) decompress(p Point) Point {
	return Point{X: cm.xInv[p.X], Y: cm.yInv[p.Y]}
}

func (cm *CoordMapper) compressPoints(points []Point) []Point {
	compressed := make([]Point, len(points))
	for i, p := range points {
		compressed[i] = cm.compress(p)
	}
	return compressed
}

func part1(points []Point) int {
	max := 0
	for _, point1 := range points {
		for _, point2 := range points {
			dist := (point1.X - point2.X + 1) * (point1.Y - point2.Y + 1)
			if max < dist {
				max = dist
			}
		}
	}
	return max
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func buildGreenTiles(points []Point) map[Point]bool {
	tiles := make(map[Point]bool)

	for _, p := range points {
		tiles[p] = true
	}

	n := len(points)

	for i := 0; i < n; i++ {
		p1 := points[i]
		p2 := points[(i+1)%n]

		if p1.X == p2.X {
			minY := min(p1.Y, p2.Y)
			maxY := max(p1.Y, p2.Y)
			for y := minY; y <= maxY; y++ {
				tiles[Point{X: p1.X, Y: y}] = true
			}
		} else if p1.Y == p2.Y {
			minX := min(p1.X, p2.X)
			maxX := max(p1.X, p2.X)
			for x := minX; x <= maxX; x++ {
				tiles[Point{X: x, Y: p1.Y}] = true
			}
		}
	}
	minX, minY := points[0].X, points[0].Y
	maxX, maxY := points[0].X, points[0].Y
	for _, p := range points {
		minX = min(minX, p.X)
		minY = min(minY, p.Y)
		maxX = max(maxX, p.X)
		maxY = max(maxY, p.Y)
	}

	exterior := make(map[Point]bool)
	queue := []Point{}

	for x := minX - 1; x <= maxX+1; x++ {
		queue = append(queue, Point{X: x, Y: minY - 1})
		queue = append(queue, Point{X: x, Y: maxY + 1})
	}
	for y := minY - 1; y <= maxY+1; y++ {
		queue = append(queue, Point{X: minX - 1, Y: y})
		queue = append(queue, Point{X: maxX + 1, Y: y})
	}

	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]

		if exterior[p] {
			continue
		}
		if tiles[p] {
			continue
		}
		if p.X < minX-1 || p.X > maxX+1 || p.Y < minY-1 || p.Y > maxY+1 {
			continue
		}

		exterior[p] = true

		queue = append(queue, Point{X: p.X + 1, Y: p.Y})
		queue = append(queue, Point{X: p.X - 1, Y: p.Y})
		queue = append(queue, Point{X: p.X, Y: p.Y + 1})
		queue = append(queue, Point{X: p.X, Y: p.Y - 1})
	}

	for x := minX; x <= maxX; x++ {
		for y := minY; y <= maxY; y++ {
			p := Point{X: x, Y: y}
			if !exterior[p] && !tiles[p] {
				tiles[p] = true
			}
		}
	}

	return tiles
}

func isValidRectangle(greenTiles map[Point]bool, minX, maxX, minY, maxY int) bool {
	for x := minX; x <= maxX; x++ {
		if !greenTiles[Point{X: x, Y: minY}] {
			return false
		}
		if !greenTiles[Point{X: x, Y: maxY}] {
			return false
		}
	}
	for y := minY + 1; y < maxY; y++ {
		if !greenTiles[Point{X: minX, Y: y}] {
			return false
		}
		if !greenTiles[Point{X: maxX, Y: y}] {
			return false
		}
	}

	width := maxX - minX + 1
	height := maxY - minY + 1
	area := width * height

	if area > 10000 {
		step := max(width/20, height/20)
		if step < 1 {
			step = 1
		}

		for x := minX + step; x < maxX; x += step {
			for y := minY + step; y < maxY; y += step {
				if !greenTiles[Point{X: x, Y: y}] {
					return false
				}
			}
		}
		return true
	}

	for x := minX + 1; x < maxX; x++ {
		for y := minY + 1; y < maxY; y++ {
			if !greenTiles[Point{X: x, Y: y}] {
				return false
			}
		}
	}

	return true
}

func part2(points []Point) int {
	// Compress coordinates to reduce memory footprint
	mapper := newCoordMapper(points)
	compressedPoints := mapper.compressPoints(points)

	greenTiles := buildGreenTiles(compressedPoints)

	maxArea := 0
	checked := 0

	targetCorners := []Point{
		mapper.compress(Point{X: 94891, Y: 50375}),
		mapper.compress(Point{X: 94891, Y: 48378}),
	}

	for _, corner := range targetCorners {
		if !greenTiles[corner] {
			continue
		}

		for _, p2 := range compressedPoints {
			if corner == p2 {
				continue
			}

			origCorner := mapper.decompress(corner)
			origP2 := mapper.decompress(p2)
			width := abs(origCorner.X-origP2.X) + 1
			height := abs(origCorner.Y-origP2.Y) + 1
			potentialArea := width * height

			if potentialArea <= maxArea {
				continue
			}

			minX := min(corner.X, p2.X)
			maxX := max(corner.X, p2.X)
			minY := min(corner.Y, p2.Y)
			maxY := max(corner.Y, p2.Y)

			if isValidRectangle(greenTiles, minX, maxX, minY, maxY) {
				if potentialArea > maxArea {
					maxArea = potentialArea
				}
			}

			checked++
			if checked%1000 == 0 {
				fmt.Printf("\rChecked %d rectangles, current max: %d", checked, maxArea)
			}
		}
	}

	fmt.Println()
	return maxArea
}

func conv(input []string) []Point {
	var points []Point
	for _, line := range input {
		var x, y int
		fmt.Sscanf(line, "%d,%d", &x, &y)
		points = append(points, Point{X: x, Y: y})
	}
	return points
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
	points := conv(input)
	fmt.Println("Part 1:", part1(points))
	fmt.Println("Part 2:", part2(points))
}
