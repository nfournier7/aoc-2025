package main

// Advent of Code: Day 03: Lobby
//
// https://adventofcode.com/
// https://adventofcode.com/2025/day/3

import (
	"aoc-2025/number"
	"aoc-2025/stream"
	"aoc-2025/system"
	"fmt"
	"sort"
	"strings"
	"time"
)

const (
    DataPath = "input/data.input"
)

type MaxMap = map[int]uint8

type Bank struct {
    batteries []uint8
}

type Data struct {
    banks []Bank
}

func (b *Bank) CalculateJoltage(size int) uint64 {
    maxMap := MaxMap{}

    lastIndexMax := -1
    indexMax := 0

    for len(maxMap) < size {
        maxRightValue := uint8(0)

        for i := len(b.batteries) - (size - len(maxMap)); i > lastIndexMax; i-- {
            battery := b.batteries[i]

            _, ok := maxMap[i]

            if (ok) {
                continue
            }

            if (battery >= maxRightValue) {
                maxRightValue = battery
                indexMax = i
            }
        }

        maxMap[indexMax] = maxRightValue
        lastIndexMax = indexMax
    }

	keys := make([]int, 0, len(maxMap))
	for k := range maxMap {
		keys = append(keys, k)
	}

    sort.Ints(keys)

    var builder strings.Builder

    for _, key := range keys {
        builder.WriteString(fmt.Sprintf("%d", maxMap[key]))
    }

    return number.ParseUint[uint64](builder.String())
}

func partOne(data *Data) uint64 {
    return stream.SumFn(data.banks, func(bank Bank) uint64 { return bank.CalculateJoltage(2) })
}

func partTwo(data *Data) uint64 {
    return stream.SumFn(data.banks, func(bank Bank) uint64 { return bank.CalculateJoltage(12) })
}

func parseInputData(lines []string, data *Data) {
    for i := range lines {
        line := lines[i]

        bank := Bank{}

        for charIndex := range line {
            bank.batteries = append(bank.batteries, number.ParseUint[uint8](line[charIndex : charIndex + 1]))
        }

        data.banks = append(data.banks, bank)
    }
}

func day03() {
    fmt.Println("Advent of Code 2025: Day03")

    start := time.Now()

    data := Data{}

    parseInputData(system.ReadLines(DataPath), &data)

    fmt.Printf("Part One: result: %d\n", partOne(&data))
    fmt.Printf("Part Two: result: %d\n", partTwo(&data))

    fmt.Printf("Time elapsed: %d ns\n", time.Since(start).Nanoseconds())

    fmt.Println("End of Day03")
}

func main() {
    day03()
}
