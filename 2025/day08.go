package days

import (
	"aoc/types"
	"fmt"
	"math"
	"slices"
	"sort"
)

type connection struct {
	a        types.Vector3
	b        types.Vector3
	distance float64
}

type link struct {
	// parent types.Vector3
	points []types.Vector3
}

var conns []connection
var links []link

func Day08_part1(points []types.Vector3) {

	for i, a := range points {
		for j := i + 1; j < len(points); j++ {
			b := points[j]

			a2 := math.Pow(float64(b.X-a.X), 2)
			b2 := math.Pow(float64(b.Y-a.Y), 2)
			c2 := math.Pow(float64(b.Z-a.Z), 2)
			d := math.Sqrt(a2 + b2 + c2)

			conns = append(conns, connection{a: a, b: b, distance: d})
		}
	}

	sort.Slice(conns, func(i, j int) bool {
		return conns[i].distance < conns[j].distance
	})

	conns = conns[:10]

	for _, c := range conns {
		links = append(links, link{[]types.Vector3{c.a, c.b}})
		fmt.Println(c)
	}

	// work with top 10 connections

	for i := 0; i < len(links); i++ {

		for j := i + 1; j < len(links)-1; j++ {

			for _, iPoint := range links[i].points {
				idx := slices.Index(links[j].points, iPoint)

				if idx != -1 {
					for _, jPoint := range links[j].points {
						if jPoint != links[j].points[idx] {
							links[i].points = append(links[i].points, jPoint)
						}
					}

					links = slices.Delete(links, j, j+1)
				}
			}

			// // parents match
			// if links[i].parent == links[j].parent {
			// 	links[i].points = append(links[i].points, links[j].points...)
			// 	links = slices.Delete(links, j, j+1)
			// }

			// // a point is b parent
			// idx := slices.Index(links[j].points, links[i].parent)

			// if idx != -1 {
			// 	links[i].points = append(links[i].points, links[j].parent)

			// 	for fooIdx, foo := range links[j].points {
			// 		if fooIdx != idx {
			// 			links[i].points = append(links[i].points, foo)
			// 		}
			// 	}

			// 	links = slices.Delete(links, j, j+1)
			// }

			// // a parent is b point
			// idx = slices.Index(links[i].points, links[j].parent)

			// if idx != -1 {
			// 	links[i].points = append(links[i].points, links[j].points...)
			// 	links = slices.Delete(links, j, j+1)
			// }
		}
	}

	// debug
	fmt.Println("")
	for _, l := range links {
		fmt.Printf("points: %v \n", l.points)
	}
	fmt.Println(len(links))
}
