package days

import "fmt"

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

func Day04_part2(matrix [][]rune) {
	sum := 0
	height := len(matrix)
	width := len(matrix[0])

	for k := 0; k < 2000; k++ { // just some random ahh number

		for i := 1; i < height-1; i++ {

			for j := 1; j < width-1; j++ {
				char := matrix[i][j]
				nearby := 0

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
					matrix[i][j] = rune('.')
					sum++
				}
			}
		}

	}

	fmt.Println(sum)
}
