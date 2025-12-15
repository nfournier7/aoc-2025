package main

import (
	"aoc-2025/system"
	"strings"
	"testing"
)

func Test_PartOne(t *testing.T) {
    const resultExpected uint64 = 1227775554
    
    data := Data{}

    parseInputData(strings.Split(system.ReadAll("input/data.gift.shop.tests"), ","), &data)

    result := partOne(&data)

    if result != resultExpected {
        t.Fatalf(`result(%d) != resultExpected(%d)`, result, resultExpected)
    }
}

func Test_PartTwo(t *testing.T) {
    const resultExpected uint64 = 4174379265
    
    data := Data{}

    parseInputData(strings.Split(system.ReadAll("input/data.gift.shop.tests"), ","), &data)

    result := partTwo(&data)

    if result != resultExpected {
        t.Fatalf(`result(%d) != resultExpected(%d)`, result, resultExpected)
    }
}