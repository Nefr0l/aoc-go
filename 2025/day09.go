package days

import (
	"aoc/types"
	"fmt"
	"math"
	"slices"
)

var maxArea = 0
var P []types.Vector2
var Vertexes []types.Vector2
var Circuit []types.Vector2

var totalMinX int
var totalMaxX int
var totalMinY int
var totalMaxY int

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
	// get some data to work with
	P = points

	X := []int{}
	Y := []int{}
	for _, v := range P {
		X = append(X, v.X)
		Y = append(Y, v.Y)
	}

	totalMinX = slices.Min(X)
	totalMaxX = slices.Max(X)

	totalMinY = slices.Min(Y)
	totalMaxY = slices.Max(Y)

	fmt.Println("Data collected. Begin to map polygon...")

	// get circuit -- this is now fast, around 4s
	start := GetStart()
	Raycast(types.Vector2{X: 1, Y: 0}, start, []types.Vector2{start})

	fmt.Println("Polygon mapping completed, starting analyze...")

	// get max area -- this is now fast
	maxArea := 0

	for _, p1 := range P {
	out:
		for _, p2 := range P {
			// get data about 2 points
			minX := min(p1.X, p2.X)
			maxX := max(p1.X, p2.X)

			minY := min(p1.Y, p2.Y)
			maxY := max(p1.Y, p2.Y)

			overlaps := 0
			left := false
			right := false
			top := false
			bottom := false

			// optimizations
			if p1.X == p2.X || p1.Y == p2.Y {
				continue
			}

			area := (maxX - minX + 1) * (maxY - minY + 1)
			if area <= maxArea {
				continue
			}

			// do checks
			for _, c := range Circuit {
				// inner check
				if c.X > minX && c.X < maxX && c.Y > minY && c.Y < maxY {
					break out
				}

				// side check
				if !top && (c.X > minX && c.X < maxX && c.Y == minY) {
					top = true
					overlaps++
				}

				if !bottom && (c.X > minX && c.X < maxX && c.Y == maxY) {
					bottom = true
					overlaps++
				}

				if !right && (c.Y > minY && c.Y < maxY && c.X == minX) {
					right = true
					overlaps++
				}

				if !left && (c.Y > minY && c.Y < maxY && c.X == maxX) {
					left = true
					overlaps++
				}
			}

			if overlaps < 3 {
				continue
			}

			fmt.Printf("a: %v	b: %v	area: %v	overlaps: %v \n", p1, p2, area, overlaps)
			maxArea = area // area is max area at this point
		}

		//fmt.Printf("Point %v checked \n", i)
	}

	fmt.Println(maxArea)
}

func Raycast(velocity types.Vector2, current types.Vector2, vertexes []types.Vector2) {
	line := []types.Vector2{}

	for {
		if polygonFound || (current.X > totalMaxX || current.X < totalMinX || current.Y < totalMinY || current.Y > totalMaxY) {
			return
		}

		current.X += velocity.X
		current.Y += velocity.Y

		// check if current is vertex
		res := Search(current)
		line = append(line, current)

		if res == nil {
			continue
		}

		// add points to circuit
		Circuit = append(Circuit, line...)

		if len(vertexes) >= len(P) {
			polygonFound = true
			return
		}

		if slices.Contains(vertexes, current) {
			return
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
