package days

import (
	"fmt"
)

func Day04_part1(matrix [][]rune) {
	sum := 0
	height := len(matrix)
	width := len(matrix[0])

	for i := 1; i < height-1; i++ {

		for j := 1; j < width-1; j++ {

			char := matrix[i][j]
			nearby := 0

			fmt.Printf("%v ", string(char))

			if string(char) == "." {
				continue
			}
		out:
			for y := -1; y <= 1; y++ {
				for x := -1; x <= 1; x++ {
					if !(x == 0 && y == 0) && matrix[i+y][j+x] == rune('@') {
						nearby++

						if nearby > 3 {
							break out
						}
					}
				}
			}

			if nearby <= 3 {
				sum++
			}
		}
		fmt.Println()
	}

	fmt.Println(sum)

}

func Day04_part2(matrix [][]rune) { // 5-6ms
	height := len(matrix) - 1
	width := len(matrix[0]) - 1
	sum := 0
	prev := 0

	empty := rune('.')
	paper := rune('@')

	for {
		prev = sum
		for i := 1; i < height; i++ {
			for j := 1; j < width; j++ {
				char := matrix[i][j]

				if char == empty {
					continue
				}

				var nearby int8 = -1 // includes 0,0 character

			out:
				for y := -1; y <= 1; y++ {
					for x := -1; x <= 1; x++ {
						if matrix[i+y][j+x] == paper {
							nearby++

							if nearby == 4 {
								break out
							}
						}
					}
				}

				if nearby >= 4 {
					continue
				}

				matrix[i][j] = rune('.')
				sum++
			}
		}

		if prev == sum {
			break
		}
	}

	fmt.Println(sum)
}
