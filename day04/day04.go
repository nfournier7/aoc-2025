package main

// Advent of Code: Day 04: Printing Department
//
// https://adventofcode.com/
// https://adventofcode.com/2025/day/4

import (
	"aoc-2025/cond"
	"aoc-2025/stream"
	"aoc-2025/system"
	"fmt"
	"time"
)

const (
    DataPath = "input/data.input"
    NbOfRollOfPaperAroundExcludedLimit = 4
)

type Tile uint8

const (
    RollOfPaper Tile = iota
    Nothing
)

type Direction uint8

const (
    Top Direction = iota
    TopRight
    Right
    BottomRight
    Bottom
    BottomLeft
    Left
    TopLeft
    DirectionCount
)

type Storage struct {
    storage [][]Tile
}

type Position struct {
    x int
    y int
}

func (p *Position) Move(vector DirectionVector) {
    p.x += vector.itX
    p.y += vector.itY
}

type DirectionVector struct {
    itX int
    itY int
}

func (b *Storage) GetTile(x int, y int) Tile {
    if (x < 0 || x >= len(b.storage[0]) || y < 0 || y >= len(b.storage)) {
        return Nothing
    }
    return b.storage[y][x]
}

func (b *Storage) Remove(x int, y int) {
    b.storage[y][x] = Nothing
}

func (b *Storage) GetSurroundingTiles(x int, y int) []Tile {
    tiles := []Tile{}

    directionVectors := []DirectionVector{
        {0, -1}, // top
        {1, -1}, // top-right
        {1, 0}, // right
        {1, 1}, // bottom-right
        {0, 1}, // bottom
        {-1, 1}, // bottom-left
        {-1, 0}, // left
        {-1, -1}, // top-left
    }

    for i := range DirectionCount {
        position := Position{x,y}
        position.Move(directionVectors[i])
        tiles = append(tiles, b.GetTile(position.x, position.y))
    }

    return tiles
}

func calculateRollOfPaperAccessible(data *Storage, canBeRemove bool) uint64 {
    sum := uint64(0)

    for y := 0; y < len(data.storage); y++ {
        for x := 0; x < len(data.storage[y]); x++ {

            if data.GetTile(x, y) != RollOfPaper {
                continue
            }

            nbOfRollOfPaperAround := stream.Count(data.GetSurroundingTiles(x, y), func (tile Tile) bool { return tile == RollOfPaper })

            if nbOfRollOfPaperAround < NbOfRollOfPaperAroundExcludedLimit {

                if (canBeRemove) {
                    data.Remove(x, y)
                    x = 0
                    y = 0
                }

                sum++
            }
        }
    }

    return sum
}

func partTwo(data *Storage) uint64 {
    return calculateRollOfPaperAccessible(data, true)
}

func partOne(data *Storage) uint64 {
    return calculateRollOfPaperAccessible(data, false)
}

func parseInputData(lines []string, data *Storage) {
    for _, line := range lines {

        tiles := []Tile{}

        for j := range line {
            tiles = append(tiles, cond.If(line[j] == byte('@'), RollOfPaper, Nothing))
        }

        data.storage = append(data.storage, tiles)
    }
}


func day04() {
    fmt.Println("Advent of Code 2025: Day04")

    start := time.Now()

    data := Storage{}

    parseInputData(system.ReadLines(DataPath), &data)

    fmt.Printf("Part One: result: %d\n", partOne(&data))
    fmt.Printf("Part Two: result: %d\n", partTwo(&data))

    fmt.Printf("Time elapsed: %d ns\n", time.Since(start).Nanoseconds())

    fmt.Println("End of Day04")
}

func main() {
    day04()
}
