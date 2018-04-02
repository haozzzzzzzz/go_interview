package main

import (
	"fmt"
)

func reverse(x int) int {
	var y int64
	for x != 0 {
		remainder := x % 10
		x = x / 10
		y = y * 10 + int64(remainder)
	}

	if y >= -2.147483648e+09 && y <= 2.147483648e+09 {
		return int(y)
	}

	return 0
}

func main()  {
	fmt.Println(reverse(321))
}
