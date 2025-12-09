package days

import (
	"fmt"
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
	indexes := get_indexes(lines)
	fmt.Println(indexes)

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
