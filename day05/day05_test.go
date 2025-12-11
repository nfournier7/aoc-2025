package main

import (
	"aoc-2025/system"
	"testing"
)

func Test_CafeteriaPartOne(t *testing.T) {
    const resultExpected uint64 = 3
    
    data := Data{}

    parseInputData(system.ReadLines("input/data.input.test"), &data)

    result := partOne(&data)

    if result != resultExpected {
        t.Fatalf(`result(%d) != resultExpected(%d)`, result, resultExpected)
    }
}

func Test_PrintingDepartmentPartTwo(t *testing.T) {
    const resultExpected uint64 = 14
    
    data := Data{}

    parseInputData(system.ReadLines("input/data.input.test"), &data)

    result := partTwo(&data)

    if result != resultExpected {
        t.Fatalf(`result(%d) != resultExpected(%d)`, result, resultExpected)
    }
}
