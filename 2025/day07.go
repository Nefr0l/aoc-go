package days

import (
	"aoc/types"
	"fmt"
	"strings"
)

var sum int = 0
var timelines uint64 = 1
var globalLines []string
var linesLength int16 = 0

func Day07_part1(lines []string) {
	tempLines := lines
	s := strings.Index(tempLines[0], "S")
	StartBeam(types.Vector2{X: s, Y: 0}, tempLines)

	// debug
	for _, line := range tempLines {
		fmt.Println(line)
	}

	fmt.Println(sum)
}

func Day07_part2(lines []string) {
	s := strings.Index(lines[0], "S")

	for _, line := range lines {
		if strings.ContainsAny(line, "^") {
			globalLines = append(globalLines, line)
		}
	}

	linesLength = int16(len(globalLines))
	for _, line := range globalLines {
		fmt.Println(line)
	}

	StartBeam2(int16(s), 0)

	fmt.Println(timelines)
}

func StartBeam(start types.Vector2, lines []string) {
	var current types.Vector2 = start

	for i := start.Y; i >= 0; i++ {
		if current.Y >= len(lines)-1 {
			return
		}

		current.Y++

		if lines[current.Y][current.X] == '|' {
			return
		}

		if lines[current.Y][current.X] == '^' {
			sum++
			break
		}

		temp := []rune(lines[current.Y])
		temp[current.X] = '|'
		lines[current.Y] = string(temp)
	}

	fmt.Println(current)

	if (current.X-1 > 0 && lines[current.Y][current.X-1] != '|') || (current.X-1 == 0) {
		StartBeam(types.Vector2{X: current.X - 1, Y: current.Y - 1}, lines)
	}

	if (current.X+1 < len(lines[0]) && (lines[current.Y][current.X+1] != '|')) || (current.X+1 == len(lines[0])-1) {
		StartBeam(types.Vector2{X: current.X + 1, Y: current.Y - 1}, lines)
	}

}

func StartBeam2(X int16, Y int16) bool {

	if Y == linesLength-1 {
		return false
	}

	if globalLines[Y][X] != '^' {
		return false
	}

	if X > 0 {
		for y := Y; y < linesLength; y++ {
			if StartBeam2(X-1, y) == true {
				break
			}

		}

	}

	if X < linesLength-1 {
		for y := Y; y < linesLength; y++ {
			if StartBeam2(X+1, y) == true {
				break
			}
		}
	}

	fmt.Println(timelines)

	timelines++
	return true
}
