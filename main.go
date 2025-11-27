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
	fmt.Println(lines)

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

// func lines_to_char_matrix(lines []string) [][]byte {
// 	return null
// }
