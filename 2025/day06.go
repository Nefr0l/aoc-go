package days

import (
	"fmt"
	"strconv"
	"strings"
)

type LineValue struct {
	Values []int
}

type ColumnValue struct {
	Values []int
	Char   rune
}

func Day06_part1(lines []string) {

	// retrieve data to line values
	var LineValues []LineValue

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
		LineValues = append(LineValues, LineValue{Values: nums})
	}

	for _, a := range LineValues {
		fmt.Println(a)
	}

}
