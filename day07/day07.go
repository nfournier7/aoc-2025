package main

import (
	"aoc-2025/system"
	"fmt"
	"time"
)

// Advent of Code: Day 07: Laboratories
//
// https://adventofcode.com/
// https://adventofcode.com/2025/day/7

const(
    DataPath = "input/data.input"
)

type Tile uint8

const(
    Nothing Tile = iota
    Start
    Splitter 
)

var tileMap = map[byte]Tile{
    byte('.'): Nothing,
    byte('S'): Start,
    byte('^'): Splitter,
}

type Data struct{
    grid [][]byte
}

func (d *Data) isSplitter(p Position) bool {
    return tileMap[d.grid[p.y][p.x]] == Splitter
}

type Node struct{
    left *Node
    right *Node
    value uint64
    position Position
    splitted bool
}

func (n *Node) isLeaf() bool {
    return n.left == nil && n.right == nil
}

func (n *Node) getNbOfPaths() uint64 {
    if (n.isLeaf()) {
        return 1
    } else if (n.value > 0) {
        return n.value
    }

    result := uint64(0)

    if n.left != nil {
        result += n.left.getNbOfPaths()
    }

    if n.right != nil {
        result += n.right.getNbOfPaths()
    }

    n.value = result

    return result
}

func createNode(position Position) *Node {
    return &Node{nil, nil, uint64(0), position, false}
}

type Position struct{
    x int
    y int
}

func (p *Position) move() {
    p.y += 1
}

func partTwo(data *Data) uint64 {
    isStartFound := false
    startLineIndex := 0
    startPosition := Position{}
    beamIndexesMap := make(map[int]*Node)

    for i := 0; i < len(data.grid) && !isStartFound; i++ {
        for indexX := range data.grid[i] {
            if tileMap[data.grid[i][indexX]] == Start {
                isStartFound = true
                startPosition.x = indexX
                startPosition.y = i
            }
        }
    }

    root := createNode(startPosition)
    beamIndexesMap[startPosition.x] = root

    for i := startLineIndex + 1; i < len(data.grid); i++ {
        line := data.grid[i]
        for k := range beamIndexesMap {
            if (tileMap[line[k]] == Splitter) {

                node := beamIndexesMap[k]

                leftIndex := k - 1
                rightIndex := k + 1

                delete(beamIndexesMap, k)

                if rightIndex < len(line) {
                    if _, exists := beamIndexesMap[rightIndex]; exists {
                        node.right = beamIndexesMap[rightIndex]
                    } else {
                        node.right = createNode(Position{rightIndex, i})
                        beamIndexesMap[rightIndex] = node.right
                    }
                }
                if leftIndex >= 0 {
                    if _, exists := beamIndexesMap[leftIndex]; exists {
                        node.left = beamIndexesMap[leftIndex] 
                    } else {
                        node.left = createNode(Position{leftIndex, i})
                        beamIndexesMap[leftIndex] = node.left
                    }
                }
            }
        }
    }

    return root.getNbOfPaths()
}

func partOne(data *Data) uint64 {
    beamSplitCount := uint64(0)
    
    isStartFound := false
    startLineIndex := 0
    beamIndexesMap := make(map[int]struct{})

    for i := 0; i < len(data.grid) && !isStartFound; i++ {
        for indexX := range data.grid[i] {
            if tileMap[data.grid[i][indexX]] == Start {
                isStartFound = true
                startLineIndex = i
                beamIndexesMap[indexX] = struct{}{}
            }
        }
    }

    for i := startLineIndex + 1; i < len(data.grid); i++ {
        line := data.grid[i]
        for k := range beamIndexesMap {
            if (tileMap[line[k]] == Splitter) {

                delete(beamIndexesMap, k)

                leftIndex := k - 1
                rightIndex := k + 1

                beamSplitCount++

                if rightIndex < len(line) {
                    beamIndexesMap[rightIndex] = struct{}{}
                }
                if leftIndex >= 0 {
                    beamIndexesMap[leftIndex] = struct{}{}
                }
            }
        }
    }

    return beamSplitCount
}

func parseInputData(lines []string, data *Data) {
    for _, line := range lines {
        bytes := []byte{}

        for _, b := range line {
            bytes = append(bytes, byte(b))
        }

        data.grid = append(data.grid, bytes)
    }
}

func day07() {
	fmt.Println("Advent of Code 2025: Day07")

	start := time.Now()

	data := Data{}

	parseInputData(system.ReadLines(DataPath), &data)

	fmt.Printf("Part One: result: %d\n", partOne(&data))
	fmt.Printf("Part Two: result: %d\n", partTwo(&data))

	fmt.Printf("Time elapsed: %d ns\n", time.Since(start).Nanoseconds())

	fmt.Println("End of Day07")
}

func main() {
	day07()
}
