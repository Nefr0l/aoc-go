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
	get_raw_lines(lines)
	get_raw_columns()

	sum := 0

	for _, column := range Columns {
		fmt.Println(column)

		//var columnf [][]rune

	}

	fmt.Println(sum)
}

// raw
func get_raw_lines(lines []string) {
	for _, line := range lines {
		var values []string = strings.Split(line, " ")
		var nums []string

		for _, v := range values {
			switch v {
			case "*":
				nums = append(nums, "*")
			case "+":
				nums = append(nums, "+")
			default:
				nums = append(nums, v)
			}
		}
		LinesRaw = append(LinesRaw, LineValueRaw{Values: nums})
	}

	fmt.Println(LinesRaw)
}

func get_raw_columns() { // TODO: FIX THIS
	width := len(LinesRaw[0].Values)

	for x := 0; x < width; x++ {
		var col ColumnValueRaw

		for y := 0; y < len(Lines)-1; y++ {
			col.Values = append(col.Values, LinesRaw[y].Values[x])
		}

		num := LinesRaw[len(Lines)-1].Values[x]

		switch num {
		case "*":
			col.Char = "*"
		case "+":
			col.Char = "+"
		}

		ColumnsRaw = append(ColumnsRaw, col)
	}

	fmt.Println(ColumnsRaw)
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
