package main

import (
	day01 "aoc/2025"
	dict "aoc/types"
	"bufio"
	"os"
	"strings"
	"strconv"
)

var data_path = "./data.txt"

func main() {
	var lines = file_to_string_array()
	var dict = lines_to_dict(lines, 1)
	day01.Day01(dict)
}


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




func lines_to_dict(lines []string, keyLength int) []dict.Dict {
	var m []dict.Dict

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
		m = append(m, dict.Dict{ Key: key.String(), Value: valueInt })
	}

	return m
}
