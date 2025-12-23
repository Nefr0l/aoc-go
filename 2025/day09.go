package days

import (
	"aoc/types"
	"fmt"
	"math"
)

var maxArea = 0

func Day09_part1(points []types.Vector2) {
	for i, p1 := range points {
		for _, p2 := range points[i+1:] {
			width := math.Abs(float64(p2.X-p1.X)) + 1
			height := math.Abs(float64(p2.Y-p1.Y)) + 1
			area := width * height

			maxArea = max(maxArea, int(area))
		}
	}

	fmt.Println(maxArea)
}

func Day09_part2(points []types.Vector2) {
	//circuit := []types.Vector2{}
}

func SplitRay(isMovingHorizontally bool, moveY int, moveX int) {

}

func SearchPointOfCoordinates(points []types.Vector2, x int, y int) types.Vector2 {
	for _, p := range points {
		if p.X == x && p.Y == y {
			return p
		}
	}

	return types.Vector2{}
}
