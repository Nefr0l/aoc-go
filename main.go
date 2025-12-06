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
	var lines = file_to_lines()
	var matrix = lines_to_matrix(lines)
	var matrix_expanded = expand_matrix(matrix, rune('.'))
	days.Day04_part2(matrix_expanded)

}

// methods for handling files
func file_to_lines() []string {
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

// functions for working with lines
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

// functions for matrix
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
