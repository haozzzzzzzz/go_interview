package main

import (
	"fmt"
)

func findMedianSortedArrays(num1 []int, num2 []int) (result float64) {
	totalNums := append(num1, num2...)
	length := len(totalNums)
	var isOdd bool = true
	if length % 2 != 0 {
		isOdd = true
	} else {
		isOdd = false
	}
	var targetIndex int
	if length > 1 {
		targetIndex = length / 2
	} else {
		targetIndex = 0
	}

	for i := 0; i < length; i ++ {
		num := totalNums[i]
		replaceJ := 0
		for j := i + 1; j < length; j ++ {
			if num > totalNums[j] {
				num = totalNums[j]
				replaceJ = j
			}
		}

		if replaceJ > 0 {
			temp := totalNums[i]
			totalNums[i] = totalNums[replaceJ]
			totalNums[replaceJ] = temp
		}

		if targetIndex == i {
			if isOdd {
				result = float64(totalNums[i])
			} else {
				result = float64(totalNums[i] + totalNums[i - 1]) / 2
			}
		}
	}

	return
}

func main()  {
	fmt.Println(findMedianSortedArrays([]int{1,3}, []int{2}))
	fmt.Println(findMedianSortedArrays([]int{1,3}, []int{2}))
	fmt.Println(findMedianSortedArrays([]int{1}, []int{2}))
	fmt.Println(findMedianSortedArrays([]int{1}, []int{}))
}