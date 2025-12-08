package days

import (
	"aoc/types"
	"fmt"
	"slices"
)

func Day05_part1(ranges []types.Vector2, values []int) {
	sum := 0

	for _, value := range values {
		for _, rg := range ranges {
			if value >= rg.X && value <= rg.Y {
				sum++
				fmt.Println(value)
				break

			}
		}
	}

	fmt.Println(sum)
}

func Day05_part2(ranges []types.Vector2, values []int) {
	sum := 0
	fmt.Println(len(ranges))

	// part one
	for {
		oldLen := len(ranges)

		for i, r1 := range ranges {
			for j, r2 := range ranges {
				if i == j {
					continue
				}

				if r2.Y <= r1.Y && r2.X >= r1.X {
					ranges = slices.Delete(ranges, j, j+1)
				} else if r2.X < r1.X && r2.Y < r1.Y && r2.Y >= r1.X {
					ranges[i] = types.Vector2{X: r2.X, Y: r1.Y}
					ranges = slices.Delete(ranges, j, j+1)
				} else if r2.Y > r1.Y && r2.X > r1.X && r2.X <= r1.Y {
					ranges[i] = types.Vector2{X: r1.X, Y: r2.Y}
					ranges = slices.Delete(ranges, j, j+1)
				}
			}
		}

		fmt.Println(len(ranges))
		if len(ranges) == oldLen {
			break
		}
	}

	// part two
	for _, r := range ranges {
		sum += r.Y - r.X + 1
	}

	//fmt.Println(ranges)
	fmt.Println(sum)

}
