package main

import (
	"aoc-2025/system"
	"testing"
)

func Test_PartOne(t *testing.T) {
    const resultExpected uint64 = 4277556
    
    data := Data{}

    parseInputDataPartOne(system.ReadLines("input/data.input.test"), &data)

    result := partOne(&data)

    if result != resultExpected {
        t.Fatalf(`result(%d) != resultExpected(%d)`, result, resultExpected)
    }
}

func Test_PartTwo(t *testing.T) {
    const resultExpected uint64 = 3263827
    
    data := Data{}

    parseInputDataPartTwo(system.ReadLines("input/data.input.test"), &data)

    result := partTwo(&data)

    if result != resultExpected {
        t.Fatalf(`result(%d) != resultExpected(%d)`, result, resultExpected)
    }
}
