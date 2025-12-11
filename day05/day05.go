package main

// Advent of Code: Day 05: Cafeteria
//
// https://adventofcode.com/
// https://adventofcode.com/2025/day/5

import (
	"aoc-2025/number"
	"aoc-2025/stream"
	"aoc-2025/system"
	"fmt"
	"sort"
	"strings"
	"time"
)

const(
    DataPath = "input/data.input"
)

type Range struct {
    start uint64
    end uint64
}

func (r *Range) isInRange(value uint64) bool {
    return r.start <= value && value <= r.end
}

type Data struct{
    freshRanges []Range
    ingredients []uint64
}

func partTwo(data *Data) uint64 {

    sort.Slice(data.freshRanges, func(i, j int) bool {
        return data.freshRanges[i].start < data.freshRanges[j].start
    })

    coverRanges := []Range{}

    for _, r := range data.freshRanges {
        isCoverRangeUpdated := false

        for i := 0; i < len(coverRanges); i++ {
            coverRange := coverRanges[i]

            if (r.start < coverRange.start && coverRange.isInRange(r.end)) {
                coverRange.start = r.start
                isCoverRangeUpdated = true
            }

            if (r.end > coverRange.end && coverRange.isInRange(r.start)) {
                coverRange.end = r.end
                isCoverRangeUpdated = true
            }

            if (r.start < coverRange.start && r.end > coverRange.end) {
                coverRange.start = r.start
                coverRange.end = r.end
                isCoverRangeUpdated = true
            }

            if (coverRange.isInRange(r.start) && coverRange.isInRange(r.end)) {
                isCoverRangeUpdated = true
            }

            if isCoverRangeUpdated {
                coverRanges[i] = coverRange
            }
        }

        if !isCoverRangeUpdated {
            coverRanges = append(coverRanges, r)
        }
    }

	return stream.SumFn(coverRanges, func (r Range) uint64 { return (r.end - r.start) + 1 })
}

func partOne(data *Data) uint64 {
    count := uint64(0)
    
    for _, ingredient := range data.ingredients {
        for _, r := range data.freshRanges {
            if (r.isInRange(ingredient)) {
                count++
                break
            }
        }
    }

	return count
}

func parseInputData(lines []string, data *Data) {
    isParsingRange := true

    for _, line := range lines {
        if (len(line) == 0) {
            isParsingRange = false
            continue
        }

        if (isParsingRange) {
            values := strings.Split(line, "-")
            r := Range{}
            r.start = number.ParseUint[uint64](values[0])
            r.end = number.ParseUint[uint64](values[1])
            data.freshRanges = append(data.freshRanges, r)
            continue
        }

        data.ingredients = append(data.ingredients, number.ParseUint[uint64](line))
    }
}

func day05() {
	fmt.Println("Advent of Code 2025: Day05")

	start := time.Now()

	data := Data{}

	parseInputData(system.ReadLines(DataPath), &data)

	fmt.Printf("Part One: result: %d\n", partOne(&data))
	fmt.Printf("Part Two: result: %d\n", partTwo(&data))

	fmt.Printf("Time elapsed: %d ns\n", time.Since(start).Nanoseconds())

	fmt.Println("End of Day05")
}

func main() {
	day05()
}
