package main

// Advent of Code: Day 01: Secret Entrance
//
// https://adventofcode.com/
// https://adventofcode.com/2025/day/1

import (
	"aoc-2025/amath"
	"aoc-2025/cond"
	"aoc-2025/number"
	"aoc-2025/system"
	"fmt"
	"time"
)

const (
    DataPath = "input/data.lock.rotations"
    DialStartingPoint = 50
    NbOfDigits = 100
)

type Direction uint8

const (
    Right Direction = iota
    Left
)

type Rotation struct {
    direction Direction
    distance int32
}

type Rotations struct {
    rotations []Rotation
}

func parseInputData(lines []string, outRotations *Rotations) {
    for i := range lines {
        rotation := Rotation{}
        line := lines[i]

        rotation.distance = number.ParseInt[int32](line[1:])
        rotation.direction = cond.If(line[0:1] == "R", Right, Left)

        outRotations.rotations = append(outRotations.rotations, rotation)
    }
}

func partOne(rotations *Rotations) {
    var value int32 = DialStartingPoint
    var count uint32 = 0

    for i := range rotations.rotations {
        rotation := rotations.rotations[i]

        value += cond.If(rotation.direction == Right, rotation.distance, rotation.distance * -1);

        if (amath.Abs(value) % 100 == 0) {
            value = 0
            count++
        }
    }

    fmt.Printf("Part One: nb of zeros: %d\n", count)
}

func partTwo(rotations *Rotations) {
    var value int32 = DialStartingPoint
    var count uint32 = 0

    for i := range rotations.rotations {
        rotation := rotations.rotations[i]

        count += uint32(rotation.distance) / NbOfDigits
        rotation.distance %= NbOfDigits
        isPositiveBefore := amath.IsPositive(value)
        isZeroBefore := value == 0

        value += cond.If(rotation.direction == Right, rotation.distance, rotation.distance * -1); 

        isPositiveAfter := amath.IsPositive(value)

        valueAbs := amath.Abs(value)
        if (valueAbs % NbOfDigits == 0) {
            value = 0
            count++
        } else if (valueAbs > NbOfDigits) {
            count++
        } else if (!isZeroBefore && isPositiveBefore != isPositiveAfter) {
            count++
        }

        value %= NbOfDigits
    }

    fmt.Printf("Part Two: nb of zeros: %d\n", count)
}

func day01() {
    fmt.Println("Advent of Code 2025: Day01")

    start := time.Now()

    rotations := Rotations{}
    parseInputData(system.ReadLines(DataPath), &rotations)

    partOne(&rotations)
    partTwo(&rotations)

    fmt.Printf("Time elapsed: %d ns\n", time.Since(start).Nanoseconds())

    fmt.Println("End of Day01")
}

func main() {
    day01()
}
