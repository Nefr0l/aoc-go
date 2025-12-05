package days

import (
	t "aoc/types"
	"fmt"
	"strings"
)

func Day02_part1(values []t.Vector2) {
	sum := 0

	for _, value := range values {
		for i := value.X; i <= value.Y; i++ {
			str := fmt.Sprint(i)
			l := len(str)

			if l%2 != 0 {
				continue
			}

			left := ""
			right := ""

			for i := 0; i < l; i++ {
				if i < l/2 {
					left += string(str[i])
				} else {
					right += string(str[i])
				}
			}

			if left != right {
				continue
			}

			sum += i
		}

		fmt.Println(sum)
	}

	fmt.Println(sum)
}

func Day02_part2(values []t.Vector2) {
	sum := 0

	for _, value := range values {
		for i := value.X; i <= value.Y; i++ {
			str := fmt.Sprint(i)
			l := len(str)
			n := l / 2

			//fmt.Printf("string: %v length: %v, n: %v\n", str, l, n)

			for j := 1; j <= n; j++ {
				if l%j == 0 {
					if strings.Repeat(str[:j], l/j) == str {
						sum += i
						break
					}
				}
			}
		}
	}

	fmt.Println(sum)
}
