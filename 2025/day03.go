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
		mxIndex := -1

		fmt.Println(nums)
		fmt.Println(len(nums))
		l := len(nums)

		for j := 0; j < 12; j++ {
			// don't know what the hell i did here but it works
			lTrim := mxIndex + 1
			rTrim := l - 11 + j
			numsTrim := nums[lTrim:rTrim]

			mx := slices.Max(numsTrim)
			mxIndex = slices.Index(numsTrim, mx) + lTrim // index for nums

			//fmt.Printf("Numbers trimmed: %v (%v-%v)\n", numsTrim, lTrim, rTrim)

			n += mx * int(math.Pow10(11-j))

		}

		sum += n

		// this is part 1
		// numsTrim1 := nums[:len(nums)-1]
		// max1 := slices.Max(numsTrim1)
		// i := slices.Index(nums, max1)

		// numsTrim2 := nums[i+1:]
		// max2 := slices.Max(numsTrim2)

		// x := max1*10 + max2

		// fmt.Println(nums)
		// fmt.Println(max2)
	}

	fmt.Println(sum)
}
