package days

import (
	"aoc/types"
	"fmt"
	"math"
)

type point struct {
	position types.Vector3
	link     *link
}

type link struct {
	points []point
}

func Day08_part1(values []types.Vector3) {
	var points []point
	var links []*link

	for _, v := range values {
		points = append(points, point{position: v})
	}

	for i, a := range points {
		// calculate closest point
		minimalDistance := 0.0
		closestIndex := 0

		for j, b := range points {
			d := math.Sqrt(math.Pow(float64(b.position.X-a.position.X), 2) + math.Pow(float64(b.position.Y-a.position.Y), 2) + math.Pow(float64(b.position.Z-a.position.Z), 2))

			if i == j {
				continue
			}

			if minimalDistance == 0 {
				minimalDistance = d
				closestIndex = j
			} else if d < minimalDistance {
				minimalDistance = d
				closestIndex = j
			}
		}

		// TODO: 10 or 1000 pairs max
		// do something with the closest point

		if points[i] == points[closestIndex] {
			continue
		}

		if points[closestIndex].link == nil { // if b does not have link
			var l link
			l.points = append(l.points, a, points[closestIndex]) // maybe check for link's length ?
			links = append(links, &l)

			points[i].link = &l
			points[closestIndex].link = points[i].link
		} else if points[closestIndex].link != points[i].link { // if b does have link and it's not the same link as a's link
			points[closestIndex].link.points = append(points[closestIndex].link.points, points[i])
			points[i].link = points[closestIndex].link
		}

		// print links

	}

	fmt.Println("")

	for _, l := range links {
		fmt.Println(l)
	}

}
