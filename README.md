# 🎄 Advent of Code 2025 - Go Solutions

My solutions to [Advent of Code 2025](https://adventofcode.com/2025) written in Go.

## 📁 Project Structure

```
AdventOfCode2025/
├── 01/
│   ├── main.go      # Solution for Day 1
│   └── test.txt     # Test input
├── 02/
│   ├── main.go      # Solution for Day 2
│   └── test.txt     # Test input
├── 03/
│   ├── main.go      # Solution for Day 3
│   └── test.txt     # Test input
├── 04/
│   ├── main.go      # Solution for Day 4
│   └── test.txt     # Test input
├── 05/
│   ├── main.go      # Solution for Day 5
│   └── test.txt     # Test input
├── 06/
│   ├── main.go      # Solution for Day 6
│   └── test.txt     # Test input
├── 07/
│   ├── main.go      # Solution for Day 7
│   └── test.txt     # Test input
├── 08/
│   ├── main.go      # Solution for Day 8
│   └── test.txt     # Test input
├── 09/
│   ├── main.go      # Solution for Day 9
│   └── test.txt     # Test input
├── 10/
│   ├── main.go      # Solution for Day 10
│   └── test.txt     # Test input
├── 11/
│   ├── main.go      # Solution for Day 11
│   └── test.txt     # Test input
└── README.md
```

## 🚀 Running Solutions

Each day's solution is contained in its own directory. To run a specific day's solution:

```bash
cd 01  # or any day number
go run main.go
```

## 📝 Solutions

### Day 1: Dial Puzzle
- **Part 1**: Count how many times the dial passes through position 0
- **Part 2**: Count full rotations (100 positions) including crossings

**Key Concepts**: Modular arithmetic, positive modulo implementation

### Day 2: Invalid ID Detection
- **Part 1**: Find IDs where the first half equals the second half (e.g., 1234**1234**, 12**12**)
- **Part 2**: Find IDs that are repeated patterns of any length (e.g., 1234**1234**, 123**123**123, 1**1**1**1**)

**Key Concepts**: String pattern matching, repeated pattern recognition

### Day 3: Maximum Number Selection
- **Part 1**: Find the largest two-digit number by selecting two digits from a string
- **Part 2**: Select the largest 12-digit number possible from the string

**Key Concepts**: Greedy algorithms, digit selection optimization

### Day 4: Grid Pattern Analysis
- **Part 1**: Count positions marked with "@" that have at least one empty adjacent cell
- **Part 2**: Iteratively remove "@" symbols that meet the criteria until none remain

**Key Concepts**: Grid traversal, neighbor checking, iterative processing

### Day 5: Range Matching and Merging
- **Part 1**: Count how many IDs fall within given ranges
- **Part 2**: Merge overlapping ranges and count total numbers covered

**Key Concepts**: Interval merging, range intersection, overlap detection

### Day 6: Column Operations
- **Part 1**: Apply operations (+, *) column-wise across rows of numbers
- **Part 2**: Parse vertical numbers and apply operations based on operator positions

**Key Concepts**: 2D array processing, vertical number parsing, operator application

### Day 7: Path Counting with Splits
- **Part 1**: Simulate falling objects through a grid, counting collision points
- **Part 2**: Count all possible paths from start to end with branching at split points

**Key Concepts**: Dynamic programming, memoization, path counting, recursive traversal

### Day 8: Junction Box Clustering
- **Part 1**: Connect 1000 closest pairs of junction boxes and find product of 3 largest circuits
- **Part 2**: Connect all boxes into one circuit and multiply X coordinates of final connection

**Key Concepts**: Union-Find (Disjoint Set Union), Minimum Spanning Tree, greedy clustering, 3D distance calculation

### Day 9: Rectangle Area with Constraints
- **Part 1**: Find largest rectangle using any two red tiles as opposite corners
- **Part 2**: Find largest rectangle where red tiles are opposite corners and all tiles inside are red or green (boundary + interior)

**Key Concepts**: Computational geometry, flood fill, polygon interior detection, coordinate compression, rectangle validation

### Day 10: XOR Path Finding
- **Part 1**: Find shortest sequence of XOR operations to reach target using available values
- **Part 2**: TBD

**Key Concepts**: Breadth-First Search (BFS), XOR operations, binary manipulation, state space exploration

### Day 11: Tree Path Counting
- **Part 1**: Count all paths from root node to "out" node in a tree structure
- **Part 2**: Count paths that pass through both "fft" and "dac" nodes before reaching "out"

**Key Concepts**: Tree traversal, recursion, memoization, graph building from input

## 🛠️ Prerequisites

- Go 1.23 or higher

## 📊 Progress

| Day | Part 1 | Part 2 | Solution |
|-----|--------|--------|----------|
| 01  | ⭐     | ⭐     | [main.go](01/main.go) |
| 02  | ⭐     | ⭐     | [main.go](02/main.go) |
| 03  | ⭐     | ⭐     | [main.go](03/main.go) |
| 04  | ⭐     | ⭐     | [main.go](04/main.go) |
| 05  | ⭐     | ⭐     | [main.go](05/main.go) |
| 06  | ⭐     | ⭐     | [main.go](06/main.go) |
| 07  | ⭐     | ⭐     | [main.go](07/main.go) |
| 08  | ⭐     | ⭐     | [main.go](08/main.go) |
| 09  | ⭐     | ⭐     | [main.go](09/main.go) |
| 10  | ⭐     |        | [main.go](10/main.go) |
| 11  | ⭐     | ⭐     | [main.go](11/main.go) |

## 📖 About Advent of Code

[Advent of Code](https://adventofcode.com) is an annual event featuring daily programming puzzles throughout December. Each puzzle consists of two parts, with the second part unlocking after completing the first.

## 📄 License

This project is open source and available for educational purposes.

---

⭐ **Star this repo if you find it helpful!**
