package main

// Advent of Code: Day 06: Trash Compactor
//
// https://adventofcode.com/
// https://adventofcode.com/2025/day/6

import (
	"aoc-2025/cond"
	"aoc-2025/number"
	"aoc-2025/stream"
	"aoc-2025/system"
	"fmt"
	"regexp"
	"strings"
	"time"
)

const(
    DataPath = "input/data.input"
)

type Operation uint8

const(
    Addition Operation = iota
    Multiplication 
)

type Data struct{
    numbersGrid [][]uint64
    operationsGrid []Operation
}

func calculateEquation(numbers []uint64, operation Operation) uint64 {
    result := cond.If(operation == Multiplication, uint64(1), uint64(0))

    for _, number := range numbers {
        result = cond.If(operation == Multiplication, result * number, result + number)
    }

    return result
}

func calculateEquations(data *Data) uint64 {
	sum := uint64(0)
    
    for i, numbers := range data.numbersGrid {
        sum += calculateEquation(numbers, data.operationsGrid[i])   
    }

    return sum
}

func partTwo(data *Data) uint64 {
    return calculateEquations(data)
}

func partOne(data *Data) uint64 {
	return calculateEquations(data)
}

func parseInputDataPartOne(lines []string, data *Data) {
    isEquationsArrayInitialized := false

    for _, line := range lines {

        reg := regexp.MustCompile("[^0-9]+")
        
        cleanedString := reg.ReplaceAllString(line, " ")
        matches := strings.Fields(cleanedString)

        if len(matches) == 0 {
            reg = regexp.MustCompile("[^*,+]+")
            cleanedString := reg.ReplaceAllString(line, " ")
            matches = strings.Fields(cleanedString)
            data.operationsGrid = stream.Map(matches, func (operationStr string) Operation { return cond.If(operationStr == "*", Multiplication, Addition) })
            continue
        }

        numbers := stream.Map(matches, func (numberStr string) uint64 { return number.ParseUint[uint64](numberStr) })

        if !isEquationsArrayInitialized {
            for i := 0; i < len(numbers); i++ {
                equationLine := []uint64{}
                data.numbersGrid = append(data.numbersGrid, equationLine)
            }
            isEquationsArrayInitialized = true
        }

        for i, number := range numbers {
            data.numbersGrid[i] = append(data.numbersGrid[i], number)
        }
    }
}

func parseInputDataPartTwo(lines []string, data *Data) {

    equationLine := []uint64{}

    for i := len(lines[0]) - 1; i >= 0; i-- {
        n := []byte{}
        isSeparator := true
        for _, line := range lines {
            if byte('*') == line[i] {
                data.operationsGrid = append(data.operationsGrid, Multiplication)
                continue
            }

            if byte('+') == line[i] {
                data.operationsGrid = append(data.operationsGrid, Addition)
                continue
            }

            if byte(' ') != line[i] {
                isSeparator = false
                n = append(n, line[i])
            }
        }

        if isSeparator {
            data.numbersGrid = append(data.numbersGrid, equationLine)
            equationLine = []uint64{}
        } else {
            equationLine = append(equationLine, number.ParseUint[uint64](string(n)))
        }
    }
    data.numbersGrid = append(data.numbersGrid, equationLine)
}

func day06() {
	fmt.Println("Advent of Code 2025: Day06")

	start := time.Now()

	data := Data{}

	parseInputDataPartOne(system.ReadLines(DataPath), &data)

	fmt.Printf("Part One: result: %d\n", partOne(&data))

	data = Data{}

	parseInputDataPartTwo(system.ReadLines(DataPath), &data)

	fmt.Printf("Part Two: result: %d\n", partTwo(&data))

	fmt.Printf("Time elapsed: %d ns\n", time.Since(start).Nanoseconds())

	fmt.Println("End of Day06")
}

func main() {
	day06()
}
