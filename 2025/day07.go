package days

import (
	"aoc/types"
	"fmt"
	"slices"
	"strings"
)

var sum int = 0
var timelines uint64 = 1
var globalLines []string
var linesLength int = 0

type Beam struct {
	index int
	value int
}

var Beams []Beam

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

	globalLines = lines
	linesLength = len(globalLines)
	Beams = append(Beams, Beam{s, 1})

	for i := 2; i < linesLength; i += 2 {
		//fmt.Printf("line: %v , res: %v \n", i, timelines)

		for j := 0; j < len(Beams); j++ {
			beam := Beams[j]

			if globalLines[i][beam.index] == '^' {
				timelines += uint64(beam.value)

				beam1 := Beam{beam.index - 1, beam.value}
				beam2 := Beam{beam.index + 1, beam.value}

				Beams[j] = beam1
				Beams = append(Beams, beam2)
			}
		}

		Beams = RemoveDuplicates()
	}

	fmt.Println(timelines)
}

func RemoveDuplicates() []Beam {
	var temp []Beam = Beams

	for i := 0; i < len(temp); i++ {
		for j := 0; j < len(temp); j++ {
			if temp[i].index == temp[j].index && i != j {
				temp[i].value += temp[j].value
				temp = slices.Delete(temp, j, j+1)
			}
		}
	}

	return temp
}

// this is for part one
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
