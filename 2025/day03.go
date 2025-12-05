package days

import (
	"fmt"
	"math"
	"slices"
)

func Day03_part1(lines []string) {
	sum := 0

	for _, line := range lines {
		nums := []int{}

		for _, i := range line {
			nums = append(nums, int(i)-48)
		}

		n := 0
		in := -1

		fmt.Println(nums)
		fmt.Println(len(nums))
		l := len(nums)

		for j := 0; j < 12; j++ {

			nTrim := nums[in+1 : (l - (11 - j))] // here
			fmt.Println(nTrim)
			fmt.Println(l - (11 - j))

			mx := slices.Max(nTrim)
			in = slices.Index(nTrim, mx)

			fmt.Println(j)

			n += mx * int(math.Pow10(l-j))

		}

		// numsTrim1 := nums[:len(nums)-1]
		// max1 := slices.Max(numsTrim1)
		// i := slices.Index(nums, max1)

		// numsTrim2 := nums[i+1:]
		// max2 := slices.Max(numsTrim2)

		// x := max1*10 + max2
		sum += n
		fmt.Println(n)

		// fmt.Println(nums)
		// fmt.Println(max2)
	}

	fmt.Println(sum)
}
