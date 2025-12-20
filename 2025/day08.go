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
	for i, a := range points { // get shortest connections
		for _, b := range points[i+1:] {
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
	for _, c := range conns { // get n shortest links
		links = append(links, []types.Vector3{c.a, c.b})
	}

	for i := 0; i < len(links)-1; i++ { // join links together
		Shorten(i)
	}

	sort.Slice(links, func(i, j int) bool { // sort links
		return len(links[i]) > len(links[j])
	})

	sum := 1
	links = links[:3]
	for _, v := range links { // get top 3 links
		sum *= len(v)
	}

	fmt.Println(sum)
}

func Day08_part2(points []types.Vector3) {
	// todo
}

func Shorten(i int) {
	var D []int // indexes to delete

	for j := i + 1; j < len(links); j++ {
		for _, point := range links[j] {
			idx := slices.Index(links[i], point)

			if idx == -1 {
				continue
			}

			links[i] = append(links[i], links[j]...)
			D = append(D, j)
			break
		}
	}
	Clear(i)

	for _, d := range slices.Backward(D) {
		links = slices.Delete(links, d, d+1)
	}

	if len(D) > 0 {
		Shorten(i)
	}
}

func Clear(i int) { // clear duplicates of links[i]
	var cleared = links[i]

	sort.Slice(cleared, func(m, n int) bool {
		return cleared[m].X < cleared[n].X
	})

	seen := make(map[types.Vector3]bool)
	result := []types.Vector3{}

	for _, v := range cleared {
		if !seen[v] {
			seen[v] = true
			result = append(result, v)
		}
	}

	links[i] = result
}
