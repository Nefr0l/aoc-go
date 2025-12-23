package days

import (
	"aoc/types"
	"fmt"
	"math"
)

var maxArea = 0
var P []types.Vector2
var Vertexes []types.Vector2
var Circuit []types.Vector2

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
	P = points

	// generally this works on test data, but is too slow for full input

	// get circuit -- this is fairly quick
	start := P[0]

	Raycast(types.Vector2{X: 1, Y: 0}, P[0], []types.Vector2{start})
	fmt.Println(Vertexes)
	fmt.Println(Circuit)

	// get max area -- this is too slow
	maxArea := 0

	for i, p1 := range P {
	out:
		for _, p2 := range P[i+1:] {
			minX := min(p1.X, p2.X)
			maxX := max(p1.X, p2.X)

			minY := min(p1.Y, p2.Y)
			maxY := max(p1.Y, p2.Y)

			for x := minX + 1; x < maxX; x++ {
				for y := minY + 1; y < maxY; y++ {
					for _, c := range Circuit {
						if c.X > maxX || c.X < minX || c.Y < minY || c.Y > maxY {
							continue
						}

						if c.X == x && c.Y == y {
							break out
						}
					}
				}
			}

			area := (maxX - minX + 1) * (maxY - minY + 1)

			if area > maxArea {
				maxArea = area
			}
		}
	}

	fmt.Println(maxArea)
}

func Raycast(velocity types.Vector2, start types.Vector2, vertexes []types.Vector2) {
	current := start
	line := []types.Vector2{}

	for range 99999 {
		current.X += velocity.X
		current.Y += velocity.Y

		res := Search(current.X, current.Y)
		line = append(line, current)

		if res != nil {
			fmt.Println(current)
			//Vertexes = append(Vertexes, current)
			Circuit = append(Circuit, line...)

			for _, v := range vertexes {
				if v.X == current.X && v.Y == current.Y {
					return
				}
			}

			vertexes = append(vertexes, current)

			if velocity.X != 0 {
				Raycast(types.Vector2{X: 0, Y: 1}, current, vertexes)
				Raycast(types.Vector2{X: 0, Y: -1}, current, vertexes)
			} else if velocity.Y != 0 {
				Raycast(types.Vector2{X: -1, Y: 0}, current, vertexes)
				Raycast(types.Vector2{X: 1, Y: 0}, current, vertexes)
			}
		}

	}
}

func Search(x int, y int) *types.Vector2 {
	for _, p := range P {
		if p.X == x && p.Y == y {
			return &p
		}
	}

	return nil
}
