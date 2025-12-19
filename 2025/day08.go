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
	}

	for i := 0; i < len(links)-1; i++ {
		fmt.Println("")
		fmt.Println(i)

		for _, l := range links {
			fmt.Printf("points: %v \n", l.points)
		}
		fmt.Println(len(links))

		ShortenLink(i)
	}
}

func ShortenLink(index int) link { // error is here
	RemoveDupes(index)

	for j := index + 1; j < len(links); j++ {
		for jPoint := 0; jPoint < len(links[j].points); jPoint++ {

			n := links[j].points[jPoint]
			idx := slices.Index(links[index].points, n)

			if idx != -1 {
				links[index].points = append(links[index].points, links[j].points...)
				links = slices.Delete(links, j, j+1)
				ShortenLink(index)
			}
		}
	}

	return links[index]
}

func RemoveDupes(index int) {
	var temp link = links[index]

	sort.Slice(temp.points, func(i, j int) bool {
		return temp.points[i].X <= temp.points[j].X
	})

	for i := 0; i < len(temp.points)-1; i++ {
		if temp.points[i] == temp.points[i+1] {
			temp.points = slices.Delete(temp.points, i, i+1)
		}
	}

	links[index] = temp
}
