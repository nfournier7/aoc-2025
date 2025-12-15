package main

import (
	"aoc-2025/system"
	"testing"
)

func Test_PartOne(t *testing.T) {
    resultExpected := uint64(357)
    
    data := Data{}

    parseInputData(system.ReadLines("input/data.input.test"), &data)

    result := partOne(&data)

    if result != resultExpected {
        t.Fatalf(`result(%d) != resultExpected(%d)`, result, resultExpected)
    }

    resultExpected = uint64(17158)

    data = Data{}

    parseInputData(system.ReadLines("input/data.input"), &data)

    result = partOne(&data)

    if result != resultExpected {
        t.Fatalf(`result(%d) != resultExpected(%d)`, result, resultExpected)
    }
}

func Test_PartTwo(t *testing.T) {
    resultExpected := uint64(3121910778619)
    
    data := Data{}

    parseInputData(system.ReadLines("input/data.input.test"), &data)

    result := partTwo(&data)

    if result != resultExpected {
        t.Fatalf(`result(%d) != resultExpected(%d)`, result, resultExpected)
    }

    resultExpected = uint64(170449335646486)

    data = Data{}

    parseInputData(system.ReadLines("input/data.input"), &data)

    result = partTwo(&data)

    if result != resultExpected {
        t.Fatalf(`result(%d) != resultExpected(%d)`, result, resultExpected)
    }
}