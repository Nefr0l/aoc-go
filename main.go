package main

import (
	day01 "aoc/2025"
	"bufio"
	"fmt"
	"os"
)

var data_path = "./data.txt"

func main() {
	fmt.Println("Hello World!")

	var lines = file_to_string_array()
	var runes = lines_to_ascii_matrix(lines)
	fmt.Println(runes)

	day01.Day01()
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
