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
	// iterate through points
	for i, p1 := range points {
		for _, p2 := range points[i+1:] {

			if !Check(points, p1, p2) {
				continue
			}

			width := math.Abs(float64(p2.X-p1.X)) + 1
			height := math.Abs(float64(p2.Y-p1.Y)) + 1
			area := int(width * height)

			if area > maxArea {
				fmt.Println(p1)
				fmt.Println(p2)
				maxArea = area
			}
		}
	}

	fmt.Println(maxArea)
}

func Check(points []types.Vector2, a types.Vector2, b types.Vector2) bool {
	if a.X == b.X || a.Y == b.Y {
		return false
	}

	minX := min(a.X, b.X)
	maxX := max(a.X, b.X)
	minY := min(a.Y, b.Y)
	maxY := max(a.Y, b.Y)

	// check 1
	dir1 := GetDirection(points, types.Vector2{X: minX, Y: minY})
	dir2 := GetDirection(points, types.Vector2{X: maxX, Y: maxY})

	if dir1*dir2 != 8 {
		return false
	}

	// check 2 - broken
	// for _, p := range points {
	// 	dir := GetDirection(points, p)

	// 	if p.X == minX && dir != 4 {
	// 		return false
	// 	} else if p.X == maxX && dir != 2 {
	// 		return false
	// 	} else if p.Y == minY && dir != 4 {
	// 		return false
	// 	} else if p.Y == maxY && dir != 2 {
	// 		return false
	// 	}

	// 	if p.X > minX && p.X < maxX && p.Y > minY && p.Y < maxY {
	// 		return false
	// 	}
	// }

	return true
}

func GetDirection(points []types.Vector2, a types.Vector2) int {
	var top types.Vector2
	var bottom types.Vector2
	var left types.Vector2
	var right types.Vector2

	for _, b := range points {
		if a.X == b.X && a.Y > b.Y {
			top = b
		} else if a.X == b.X && a.Y < b.Y {
			bottom = b
		} else if a.Y == b.Y && a.X > b.X {
			left = b
		} else if a.Y == b.Y && a.X < b.X {
			right = b
		}
	}

	if left.X == 0 && bottom.X == 0 {
		return 1
	}

	if right.X == 0 && bottom.X == 0 {
		return 2
	}

	if top.X == 0 && right.X == 0 {
		return 3
	}

	if top.X == 0 && left.X == 0 {
		return 4
	}

	// edge cases
	return 0
}
