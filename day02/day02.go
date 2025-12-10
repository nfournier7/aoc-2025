package main

// Advent of Code: Day 02: Gift Shop
//
// https://adventofcode.com/
// https://adventofcode.com/2025/day/2

import (
	"aoc-2025/number"
	"aoc-2025/stream"
	"aoc-2025/system"
	"fmt"
	"strings"
	"time"
)

const (
    DataPath = "input/data.gift.shop"
)

type Direction uint8

const (
    Right Direction = iota
    Left
)

type IdRange struct {
    start uint64
    end uint64
}

type Data struct {
    ranges []IdRange
}

func isCodeInvalid(value uint64) bool {
    valueStr := fmt.Sprintf("%d", value)
    valueStrLength := len(valueStr)

    for sequenceLength := 1; sequenceLength <= valueStrLength / 2; sequenceLength++ {
        
        if valueStrLength % sequenceLength != 0 {
            continue
        }

        isInvalid := true
        sequence := valueStr[0: sequenceLength]

        for index := range valueStrLength {
            if valueStr[index] != sequence[index % sequenceLength] {
                isInvalid = false
                break
            }
        }

        if isInvalid {
            return true
        }
    }

    return false
}

func analyzeRangePartTwo(idRange *IdRange, invalidIds *[]uint64) {
    for value := idRange.start; value <= idRange.end; value++ {
        if (isCodeInvalid(value)) {
            *invalidIds = append(*invalidIds, value)
        }
    }
}

func analyzeRangePartOne(idRange *IdRange, invalidIds *[]uint64) {
    for value := idRange.start; value <= idRange.end; value++ {
        valueStr := fmt.Sprintf("%d", value)
        
        valueLength := len(valueStr)

        if (valueLength == 1) {
            continue
        }

        sequenceMaxSize := valueLength / 2

        firstStr := valueStr[:sequenceMaxSize]
        secondStr := valueStr[sequenceMaxSize:]

        if (len(firstStr) != len(secondStr)) {
            continue
        }

        first := number.ParseUint[uint64](firstStr)
        second := number.ParseUint[uint64](secondStr)
        
        if (first == second) {
            *invalidIds = append(*invalidIds, value)
        }
    }
}

func partTwo(data *Data) uint64 {
    invalidIds := []uint64{}

    for i := range data.ranges {
        analyzeRangePartTwo(&data.ranges[i], &invalidIds)
    }

    return stream.SumFn(invalidIds, func(value uint64) uint64 { return value })
}

func partOne(data *Data) uint64 {
    invalidIds := []uint64{}

    for i := range data.ranges {
        analyzeRangePartOne(&data.ranges[i], &invalidIds)
    }

    return stream.SumFn(invalidIds, func(value uint64) uint64 { return value })
}

func parseInputData(ranges []string, data *Data) {
    for i := range ranges {
        rangeEntry := ranges[i]

        entries := strings.Split(rangeEntry, "-")

        idRange := IdRange{}
        
        idRange.start = number.ParseUint[uint64](entries[0])
        idRange.end = number.ParseUint[uint64](entries[1])

        data.ranges = append(data.ranges, idRange)
    }
}


func day02() {
    fmt.Println("Advent of Code 2025: Day02")

    start := time.Now()

    data := Data{}

    parseInputData(strings.Split(system.ReadAll(DataPath), ","), &data)

    fmt.Printf("Part One: result: %d\n", partOne(&data))
    fmt.Printf("Part Two: result: %d\n", partTwo(&data))

    fmt.Printf("Time elapsed: %d ns\n", time.Since(start).Nanoseconds())

    fmt.Println("End of Day02")
}

func main() {
    day02()
}
