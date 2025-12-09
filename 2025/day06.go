package days

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type LineValue struct {
	Values []int
}

type LineValueRaw struct {
	Values []string
}

type ColumnValue struct {
	Values []int
	Char   rune
}

type ColumnValueRaw struct {
	Values []string
	Char   string
}

var Lines []LineValue
var LinesRaw []LineValueRaw
var Columns []ColumnValue
var ColumnsRaw []ColumnValueRaw

func Day06_part1(lines []string) {
	get_lines_values(lines)
	get_columns()
	sum := 0

	for _, column := range Columns {
		res := 0

		if column.Char == '*' {
			res = 1
		}

		for _, n := range column.Values {
			switch column.Char {
			case '*':
				res *= n
			case '+':
				res += n
			}
		}
		sum += res
	}
	fmt.Println(sum)
}

func Day06_part2(lines []string) {
	// higher than 10227054229927

	indexes := get_indexes(lines)
	indexes = append([]int{0}, indexes...)
	fmt.Println(indexes)

	totalSum := 0

	for i := 0; i < len(indexes)-1; i++ {
		// get start and end index
		startIndex := indexes[i] + 1
		endIndex := indexes[i+1]

		if i == 0 {
			startIndex = 0
		} else if i == len(indexes)-2 {
			endIndex = len(lines[0])
		}

		fmt.Printf("Calculating numbers from index %v to %v \n", startIndex, endIndex)

		// work with indexes

		var numbers []int
		for x := startIndex; x < endIndex; x++ {
			var number int = 0

			for y := 0; y < len(lines)-1; y++ {
				temp := int(lines[y][x]) - 48

				if temp == -16 {
					number /= 10
					continue

				}

				pow := len(lines) - 2 - y
				number += int(math.Pow10(pow)) * temp
			}

			numbers = append(numbers, number)
		}

		fmt.Println(numbers)

		// add or multiply
		char := lines[len(lines)-1][startIndex]

		var sum int
		switch char {
		case 42: // *
			sum = 1
		case 43: // +
			sum = 0
		}

		for _, n := range numbers {
			switch char {
			case 42:
				if n != 0 {
					sum *= n
				}

			case 43:
				sum += n
			}
		}

		totalSum += sum
		fmt.Println(sum)
	}

	fmt.Println(totalSum)

}

// raw
func get_indexes(lines []string) []int {
	spaceIndex := 0
	width := len(lines[0])
	var indexes []int

	for i := spaceIndex; i < width; i++ {

		emptyRow := true

		for j, line := range lines {
			if j == len(lines)-2 {
				break
			}

			if line[i] != ' ' {
				emptyRow = false
				break
			}
		}

		if emptyRow {
			spaceIndex = i
			indexes = append(indexes, spaceIndex)
		}
	}

	return indexes
}

// non-raw
func get_lines_values(lines []string) {
	for _, line := range lines {
		var values []string = strings.Fields(line)
		var nums []int

		for _, v := range values {
			switch v {
			case "*":
				nums = append(nums, 0)
			case "+":
				nums = append(nums, 1)
			default:
				num, _ := strconv.Atoi(v)
				nums = append(nums, num)
			}
		}
		Lines = append(Lines, LineValue{Values: nums})
	}

	fmt.Println(Lines)
}

func get_columns() {
	width := len(Lines[0].Values)

	for x := 0; x < width; x++ {
		var col ColumnValue

		for y := 0; y < len(Lines)-1; y++ {
			col.Values = append(col.Values, Lines[y].Values[x])
		}

		num := Lines[len(Lines)-1].Values[x]

		switch num {
		case 0:
			col.Char = rune('*')
		case 1:
			col.Char = rune('+')
		}

		Columns = append(Columns, col)
	}
}
