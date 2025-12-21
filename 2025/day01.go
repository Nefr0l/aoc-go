package days

import (
	dict "aoc/types"
	"fmt"
)

var curr = 50
var zeros = 0

func Day01(dict []dict.Dict) error {
	fmt.Println(dict)

	for _, entry := range dict {
		var zerosBefore = zeros

		fmt.Printf("[ENTRY]: Value: %03d rotated by %v%02v. ", curr, entry.Key, entry.Value)

		zeros += entry.Value / 100
		entry.Value = entry.Value % 100

		switch entry.Key {
		case "L":
			if curr == 0 {
				curr = 100 - entry.Value
			} else if curr > entry.Value {
				curr -= entry.Value
			} else if curr == entry.Value {
				curr = 0
				zeros++
			} else if curr < entry.Value {
				var temp = curr
				curr = 100 - entry.Value + temp
				zeros++
			}

		case "R":
			curr += entry.Value
			if curr >= 100 {
				curr -= 100
				zeros++
			}
		}

		fmt.Printf("[END]: Zeros %02d + %02v = %02v, Rotation: %02v, Value: %00d \n", zerosBefore, zeros-zerosBefore, zeros, entry, curr)

	}

	fmt.Println(zeros)

	return nil
}
