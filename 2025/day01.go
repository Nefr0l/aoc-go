package day01

import (
	dict "aoc/types"
	"fmt"
)

var pass = 50
var zeros = 0

func Day01(dict []dict.Dict) error {
	fmt.Println(dict)
	yes := false

	for _, entry := range dict {
		var zerosBefore = zeros

		fmt.Printf("[ENTRY]: Value: %03d rotated by %v%02v. ", pass, entry.Key, entry.Value)

		zeros += entry.Value / 100
		entry.Value = entry.Value % 100

		switch entry.Key {
		case "L":
			if pass >= entry.Value {
				fmt.Print("	[TU]	")
				pass -= entry.Value

				if pass == 0 {

					// fix here
					if yes {
						zeros++
						yes = false
					} else {
						yes = true
					}

				}
			} else {
				var temp = pass
				pass = (100 - entry.Value + temp) % 100
				zeros++
			}

		case "R":
			pass += entry.Value
			if pass >= 100 {
				pass -= 100
				zeros++
			}
		}

		fmt.Printf("	[END]: Zeros %02d + %02v = %02v, Rotation: %02v, Value: %00d \n", zerosBefore, zeros-zerosBefore, zeros, entry, pass)

	}

	fmt.Println(zeros)

	return nil
}

func TakeTurn(rotation string) {

}
