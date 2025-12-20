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

var conns []connection
var links [][]types.Vector3
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
		links = append(links, []types.Vector3{c.a, c.b})
	}

	// join links together
	for i := 0; i < len(links)-1; i++ {
		Shorten(i)
	}

	// get lengths of top 3 links
	sort.Slice(links, func(i, j int) bool {
		return len(links[i]) > len(links[j])
	})

	sum := 1
	links = links[:3]
	for _, v := range links {
		sum *= len(v)
	}

	fmt.Println(sum)
}

func Day08_part2(points []types.Vector3) {
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
}

func Shorten(index int) {
	var indexesToDelete []int

	// search for links with duplicate values
	for j := index + 1; j < len(links); j++ {
		for jPoint := 0; jPoint < len(links[j]); jPoint++ {
			n := links[j][jPoint]
			idx := slices.Index(links[index], n)

			if idx != -1 {
				links[index] = append(links[index], links[j]...)
				indexesToDelete = append(indexesToDelete, j)
				break
			}
		}
	}
	Clear(index)

	// join links
	for idx := len(indexesToDelete) - 1; idx >= 0; idx-- {
		foo := indexesToDelete[idx]
		links = slices.Delete(links, foo, foo+1)
	}

	if len(indexesToDelete) > 0 {
		Shorten(index)
	}
}

func Clear(index int) {
	var temp = links[index]

	sort.Slice(temp, func(i, j int) bool {
		return temp[i].X < temp[j].X
	})

	seen := make(map[types.Vector3]bool)
	result := []types.Vector3{}

	for i := 0; i < len(temp); i++ {
		if !seen[temp[i]] {
			seen[temp[i]] = true
			result = append(result, temp[i])
		}
	}

	links[index] = result
}
