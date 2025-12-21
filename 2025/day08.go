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
	for i, a := range points {
		for _, b := range points[i+1:] {
			d := math.Pow(float64(b.X-a.X), 2) + math.Pow(float64(b.Y-a.Y), 2) + math.Pow(float64(b.Z-a.Z), 2) // no need to sqrt, distance is only for reference
			conns = append(conns, connection{a: a, b: b, distance: d})
		}
	}

	sort.Slice(conns, func(i, j int) bool {
		return conns[i].distance < conns[j].distance
	})

	conns = conns[:n] // get n shortest links
	for _, c := range conns {
		links = append(links, []types.Vector3{c.a, c.b})
	}

	for i := 0; i < len(links)-1; i++ { // join links together
		Shorten(i)
	}

	sort.Slice(links, func(i, j int) bool { // sort links
		return len(links[i]) > len(links[j])
	})

	links = links[:3] // get top 3 links
	sum := 1
	for _, v := range links {
		sum *= len(v)
	}

	fmt.Println(sum)
}

func Day08_part2(points []types.Vector3) {
	for i, a := range points {
		for _, b := range points[i+1:] {
			d := math.Pow(float64(b.X-a.X), 2) + math.Pow(float64(b.Y-a.Y), 2) + math.Pow(float64(b.Z-a.Z), 2)
			conns = append(conns, connection{a: a, b: b, distance: d})
		}
	}

	sort.Slice(conns, func(i, j int) bool {
		return conns[i].distance < conns[j].distance
	})

	for _, c := range conns {
		links = append(links, []types.Vector3{c.a, c.b})
	}

	seen := make(map[types.Vector3]bool)
	result := []types.Vector3{}
	lastLink := []types.Vector3{}

out:
	for _, l := range links {
		for _, n := range l {
			if !seen[n] {
				seen[n] = true
				result = append(result, n)
			}

			if len(result) == len(points) {
				lastLink = l
				break out
			}
		}
	}

	sum := lastLink[0].X * lastLink[1].X
	fmt.Println(lastLink)
	fmt.Println(sum)
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
