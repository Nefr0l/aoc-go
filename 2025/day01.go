package day01

import (
	dict "aoc/types"
	"fmt"
)

var pass = 50
var zeros = 0

func Day01(dict []dict.Dict) {
	fmt.Println(dict)

	for _, entry := range dict {
		fmt.Print(pass)

		fmt.Print(" , ")
		fmt.Print(entry.Value)
		entry.Value = entry.Value % 100

		switch entry.Key {
		case "L":
			if pass >= entry.Value {
				pass -= entry.Value
			} else {
				var temp = pass
				pass = (100 - entry.Value + temp) % 100
			}

		case "R":
			pass += entry.Value
			if pass >= 100 {
				pass -= 100
			}
		}

		if pass == 0 {
			zeros += 1
		}

		fmt.Print(entry)
		fmt.Println(pass)
	}

	fmt.Println(zeros)
}

func TakeTurn(rotation string) {

}
