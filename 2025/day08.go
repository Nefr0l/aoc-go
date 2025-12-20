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
var n = 1000

func Day08_part1(points []types.Vector3) {
	// get shortest connections
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

	conns = conns[:n]
	for _, c := range conns {
		links = append(links, link{[]types.Vector3{c.a, c.b}})
	}

	// join connections together
	for i := 0; i < len(links)-1; i++ {
		fmt.Println("")
		fmt.Println(i)

		for j := 0; j < 10; j++ { // random ahh value
			ShortenLink(i)
		}
	}

	// get lengths of top 3 links
	sort.Slice(links, func(i, j int) bool {
		return len(links[i].points) > len(links[j].points)
	})

	sum := 1
	for i, v := range links {
		sum *= len(v.points)
		fmt.Println(v)

		if i == 2 {
			break
		}
	}

	fmt.Println(sum)
}

func Day08_part2(points []types.Vector3) {
	// get shortest connections

	for i, a := range points {
		for j := i + 1; j < len(points); j++ {
			b := points[j]

			a2 := math.Pow(float64(b.X-a.X), 2)
			b2 := math.Pow(float64(b.Y-a.Y), 2)
			c2 := math.Pow(float64(b.Z-a.Z), 2)
			d := math.Sqrt(a2 + b2 + c2)

			conns = append(conns, connection{a: a, b: b, distance: d})
			//points = slices.Delete(points, i, i+1)
		}
	}

	fmt.Println(conns[len(conns)-1])

}

func ShortenLink(index int) link {
	var indexesToDelete []int

	// search for links with duplicate values
	for j := index + 1; j < len(links); j++ {
		for jPoint := 0; jPoint < len(links[j].points); jPoint++ {

			n := links[j].points[jPoint]
			idx := slices.Index(links[index].points, n)

			if idx != -1 {
				links[index].points = append(links[index].points, links[j].points...)
				indexesToDelete = append(indexesToDelete, j)

				break
			}
		}
	}

	// join links
	for idx := len(indexesToDelete) - 1; idx >= 0; idx-- {
		foo := indexesToDelete[idx]
		links = slices.Delete(links, foo, foo+1)
	}

	RemoveDupes(index)
	return links[index]
}

func RemoveDupes(index int) {
	var temp link = links[index]

	sort.Slice(temp.points, func(i, j int) bool {
		return temp.points[i].X <= temp.points[j].X
	})

	seen := make(map[types.Vector3]bool)
	result := []types.Vector3{}

	for i := 0; i < len(temp.points); i++ {
		if !seen[temp.points[i]] {
			seen[temp.points[i]] = true
			result = append(result, temp.points[i])
		}
	}

	links[index].points = result
}
