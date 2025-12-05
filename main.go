package main

import (
	days "aoc/2025"
	types "aoc/types"
	"bufio"
	"os"
	"strconv"
	"strings"
)

var data_path = "./demo.txt"

func main() {
	// var lines = file_to_string_array()
	// var dict = lines_to_dict(lines, 1)
	// days.Day01(dict)

	var vector = file_to_vector2()
	days.Day02_part2(vector)
}

// methods for handling files
func file_to_string_array() []string {
	file, _ := os.Open(data_path)

	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

func file_to_vector2() []types.Vector2 {
	file, _ := os.ReadFile(data_path)

	var vector []types.Vector2
	var values []string = strings.Split(string(file), ",")

	for _, value := range values {
		temp := strings.Split(value, "-")

		x, _ := strconv.Atoi(temp[0])
		y, _ := strconv.Atoi(temp[1])

		vector = append(vector, types.Vector2{X: x, Y: y})
	}

	return vector
}

// functions for working with lines[]
func lines_to_ascii_matrix(lines []string) [][]rune {
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
