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

func Day08_part1(values []types.Vector3) {
	var points []point
	var links []*link
	var distances []distance

	for _, v := range values {
		points = append(points, point{position: v})
	}

	// calculate closest diistances
	for i, a := range points {
		minimalDistance := 0.0
		closestIndex := 0

		for j, b := range points {
			if i == j {
				continue
			}

			d := math.Sqrt(math.Pow(float64(b.position.X-a.position.X), 2) + math.Pow(float64(b.position.Y-a.position.Y), 2) + math.Pow(float64(b.position.Z-a.position.Z), 2))

			if minimalDistance == 0 || d < minimalDistance {
				minimalDistance = d
				closestIndex = j
			}
		}

		d := distance{aIdx: i, bIdx: closestIndex, length: minimalDistance}

		idx := CanAdd(distances, d)
		if idx == -2 {
			distances = append(distances, d)
		} else if idx == -1 {
			continue
		} else {
			distances[idx] = d
		}
	}

	// print distances
	for _, d := range distances {
		fmt.Printf("point a: %v		point b: %v		distance: %v \n", points[d.aIdx], points[d.bIdx], d.length)
	}

	// work with top 10 minimal distances
	for _, d := range distances {
		a := &points[d.aIdx]
		b := &points[d.bIdx]

		if a == b {
			continue
		}

		if b.link == nil { // if b does not have link
			var l link
			l.points = append(l.points, *a, *b)
			links = append(links, &l)

			a.link = &l
			b.link = a.link
		} else if a.link != b.link { // if b does have link and it's not the same link as a's link
			b.link.points = append(b.link.points, *a)
			a.link = b.link
		}

	}

	// debug prints
	fmt.Println("")
	fmt.Println(len(distances))

	for _, l := range links {
		fmt.Println(l)
	}

}

func CanAdd(distances []distance, d distance) int {

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
