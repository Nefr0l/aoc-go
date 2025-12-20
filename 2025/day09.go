package days

import (
	"aoc/types"
	"fmt"
	"math"
)

func Day09_part1(points []types.Vector2) {
	maxArea := 0

	for i, p1 := range points {
		for _, p2 := range points[i+1:] {
			width := int(math.Abs(float64(p2.X-p1.X)) + 1)
			height := int(math.Abs(float64(p2.Y-p1.Y)) + 1)
			area := width * height
			maxArea = max(maxArea, area)
			//fmt.Printf("p1 is %v	p2 is %v 	area is %v\n", p1, p2, area)
		}
	}

	fmt.Println(maxArea)
}
