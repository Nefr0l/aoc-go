package days

import (
	"aoc/types"
	"fmt"
	"math"
	"slices"
)

var maxArea = 0
var P []types.Vector2 // points
var C []types.Vector2 // circuit

var minValues types.Vector2
var maxValues types.Vector2
var polygonFound bool = false

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

// higher than 225951086 and 288264150
func Day09_part2(points []types.Vector2) {
	// get circuit
	fmt.Println("Collecting data...")

	P = points
	X, Y := []int{}, []int{}

	for _, p := range P {
		X = append(X, p.X)
		Y = append(Y, p.Y)
	}

	minValues = types.Vector2{X: slices.Min(X), Y: slices.Min(Y)}
	maxValues = types.Vector2{X: slices.Max(X), Y: slices.Max(Y)}

	s := GetStart()
	Raycast(types.Vector2{X: 1, Y: 0}, s, []types.Vector2{s})

	// get max area
	fmt.Println("Getting max area...")
	maxArea := 0

	for _, p1 := range P {
	out:
		for _, p2 := range P {
			minX, maxX := min(p1.X, p2.X), max(p1.X, p2.X)
			minY, maxY := min(p1.Y, p2.Y), max(p1.Y, p2.Y)
			left, right, top, down := 0, 0, 0, 0

			// get small rects out of the way
			area := (maxX - minX + 1) * (maxY - minY + 1)
			if area <= maxArea || p1.X == p2.X || p1.Y == p2.Y {
				continue
			}

			// checks
			for _, c := range C {
				if c.X > minX && c.X < maxX && c.Y > minY && c.Y < maxY { // inner check
					break out
				}

				if c.X > minX && c.X < maxX && c.Y == minY { // side check part 1
					top = 1
				} else if c.X > minX && c.X < maxX && c.Y == maxY {
					down = 1
				} else if c.Y > minY && c.Y < maxY && c.X == minX {
					left = 1
				} else if c.Y > minY && c.Y < maxY && c.X == maxX {
					right = 1
				}
			}

			// side check part 2
			if (left + right + top + down) < 3 {
				continue
			}

			fmt.Printf("a: %v	b: %v	area: %v \n", p1, p2, area)
			maxArea = area // area is max area at this point
		}
	}

	fmt.Println(maxArea)
}

func Raycast(velocity types.Vector2, curr types.Vector2, vertexes []types.Vector2) {
	side := []types.Vector2{}

	for {
		if polygonFound || curr.X > maxValues.X || curr.X < minValues.X || curr.Y < minValues.Y || curr.Y > maxValues.Y {
			return
		}

		curr.X += velocity.X
		curr.Y += velocity.Y

		// check if current is vertex
		res := Search(curr)

		if res == nil {
			side = append(side, curr) // prevent adding vertexes to circuit
			continue
		}

		// add points to circuit
		C = append(C, side...)

		if len(vertexes) == len(P) {
			polygonFound = true
			return
		}

		if slices.Contains(vertexes, curr) {
			return
		}

		vertexes = append(vertexes, curr)

		// split ray
		if velocity.X != 0 {
			Raycast(types.Vector2{X: 0, Y: 1}, curr, vertexes)
			Raycast(types.Vector2{X: 0, Y: -1}, curr, vertexes)
		} else if velocity.Y != 0 {
			Raycast(types.Vector2{X: -1, Y: 0}, curr, vertexes)
			Raycast(types.Vector2{X: 1, Y: 0}, curr, vertexes)
		}
	}
}

func Search(point types.Vector2) *types.Vector2 {
	if slices.Contains(P, point) {
		return &point
	}

	return nil
}

func GetStart() types.Vector2 {
	point := P[0]

	for _, p := range P {
		if p.X < point.X || p.Y < point.Y {
			point = p
		}
	}

	return point
}
