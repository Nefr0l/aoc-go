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
				break
			}
		}
	}

	fmt.Println(sum)
}

func Day05_part2(ranges []types.Vector2, values []int) {
	sum := 0

	// remove duplicates
	for {
		oldLen := len(ranges)

		for i := 0; i < len(ranges); i++ {
			for j := 0; j < len(ranges); j++ {
				if i == j {
					continue
				}

				if ranges[j].Y <= ranges[i].Y && ranges[j].X >= ranges[i].X {
					ranges = slices.Delete(ranges, j, j+1)
				} else if ranges[j].X < ranges[i].X && ranges[j].Y < ranges[i].Y && ranges[j].Y >= ranges[i].X {
					ranges[i] = types.Vector2{X: ranges[j].X, Y: ranges[i].Y}
					ranges = slices.Delete(ranges, j, j+1)
				} else if ranges[j].Y > ranges[i].Y && ranges[j].X > ranges[i].X && ranges[j].X <= ranges[i].Y {
					ranges[i] = types.Vector2{X: ranges[i].X, Y: ranges[j].Y}
					ranges = slices.Delete(ranges, j, j+1)
				}
			}

			fmt.Printf("Length of ranges[]: %v \n", len(ranges))
		}

		if len(ranges) == oldLen {
			break
		}
	}

	// add sums
	for _, r := range ranges {
		sum += r.Y - r.X + 1
	}

	fmt.Println(sum)
}
