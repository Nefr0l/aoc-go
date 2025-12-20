package main

import (
	days "aoc/2025"
	types "aoc/types"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"
)

var data_path = "./demo.txt"

func main() {
	start := time.Now()
	var lines = get_lines()
	var v = lines_to_vector2(lines, ",")
	days.Day09_part1(v)

	fmt.Println(time.Since(start))
}

// specific days
func run_day05(lines []string) {
	var point = slices.Index(lines, "") // day05
	days.Day05_part2(lines_to_vector2(lines[:point], "-"), lines_to_int_slice(lines[point+1:]))
}

// file handling
func get_lines() []string {
	file, _ := os.ReadFile(data_path)
	return strings.Split(string(file), "\n")
}

// lines handling
func lines_to_vector2(lines []string, char string) []types.Vector2 {
	var vector []types.Vector2

	for _, v := range lines {
		temp := strings.Split(v, char)

		x, _ := strconv.Atoi(temp[0])
		y, _ := strconv.Atoi(temp[1])

		vector = append(vector, types.Vector2{X: x, Y: y})
	}

	return vector
}

func lines_to_vector3(lines []string, char string) []types.Vector3 {
	var vector []types.Vector3

	for _, v := range lines {
		temp := strings.Split(v, char)

		x, _ := strconv.Atoi(temp[0])
		y, _ := strconv.Atoi(temp[1])
		z, _ := strconv.Atoi(temp[2])

		vector = append(vector, types.Vector3{X: x, Y: y, Z: z})
	}

	return vector
}

func lines_to_matrix(lines []string) [][]rune {
	var runes [][]rune

	for _, line := range lines {
		var chars []rune

		for _, char := range line {
			chars = append(chars, char)
		}

		runes = append(runes, chars)
	}
	return runes
}

func lines_to_dict(lines []string, keyLength int) []types.Dict {
	var m []types.Dict

	for _, line := range lines {
		var key strings.Builder
		var value strings.Builder

		for j, char := range line {
			if j < keyLength {
				key.WriteString(string(char))
			} else {
				value.WriteString(string(char))
			}
		}

		valueInt, _ := strconv.Atoi(value.String())
		m = append(m, types.Dict{Key: key.String(), Value: valueInt})
	}

	return m
}

func lines_to_int_slice(lines []string) []int {
	var nums []int

	for _, line := range lines {
		n, _ := strconv.Atoi(line)
		nums = append(nums, n)
	}

	return nums
}

// matrix handling
func expand_matrix(matrix [][]rune, x rune) [][]rune {
	var runes [][]rune

	for i := -1; i <= len(matrix); i++ {
		var chars []rune

		for j := -1; j <= len(matrix[0]); j++ {
			if j == -1 || j == len(matrix[0]) || i == -1 || i == len(matrix) {
				chars = append(chars, x)
			} else {
				chars = append(chars, matrix[i][j])
			}
		}

		runes = append(runes, chars)
	}
	return runes
}
