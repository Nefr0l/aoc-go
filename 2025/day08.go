package days

import (
	"aoc/types"
	"fmt"
	"math"
	"slices"
)

type point struct {
	position types.Vector3
	link     *link
}

type link struct {
	points []point
}

type distance struct {
	aIdx   int
	bIdx   int
	length float64
}

var distances []*distance
var points []*point
var links []*link

func Day08_part1(values []types.Vector3) {

	for _, v := range values {
		points = append(points, &point{position: v, link: &link{}})
	}

	// calculate closest distances - sth wrong here
	for i, a := range points {
		minDistance := 0.0
		closestIdx := 0

		for j, b := range points {
			if i == j {
				continue
			}

			a2 := math.Pow(float64(b.position.X-a.position.X), 2)
			b2 := math.Pow(float64(b.position.Y-a.position.Y), 2)
			c2 := math.Pow(float64(b.position.Z-a.position.Z), 2)
			d := math.Sqrt(a2 + b2 + c2)

			if minDistance == 0 || d < minDistance {
				minDistance = d
				closestIdx = j
			}
		}

		d := distance{aIdx: i, bIdx: closestIdx, length: minDistance}
		idx := CanAdd(d)

		switch idx {
		case -2:
			distances = append(distances, &d)
		case -1:
			continue
		default:
			distances[idx] = &d
		}
	}

	// print distances
	// for _, d := range distances {
	// 	fmt.Printf("a: %v \t b: %v \t d: %v \n", points[d.aIdx], points[d.bIdx], d.length)
	// }

	// work with top 10 minimal distances
	for _, d := range distances {
		a := points[d.aIdx]
		b := points[d.bIdx]

		if len(b.link.points) == 0 { // here is error, i can feel it
			var l link = link{points: []point{*a, *b}}
			a.link = &l

			b.link = a.link

			links = append(links, a.link)
		} else if a.link != b.link {
			b.link.points = append(b.link.points, *a)
			a.link = b.link
		}

	}

	// debug prints
	fmt.Println("")
	fmt.Println(len(distances))

	for _, l := range links {
		fmt.Println(l.points)
	}
}

func CanAdd(d distance) int {
	// check length
	if len(distances) < 10 {
		return -2
	}

	// check duplicates
	for _, d2 := range distances {
		if d.aIdx == d2.bIdx && d.bIdx == d2.aIdx {
			return -1
		}
	}

	// check if length < max
	var lengths []float64
	for _, d := range distances {
		lengths = append(lengths, d.length)
	}

	mx := slices.Max(lengths)
	idx := slices.Index(lengths, mx)

	if d.length < mx {
		return idx
	}

	return -1
}
