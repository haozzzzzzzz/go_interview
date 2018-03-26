package main

import (
	"fmt"
)

func QuickSort(data []int, low int, high int) {
	var i, pivotKey, j int
	if low < high {
		pivotKey = data[low]
		i = low
		j = high

		for i < j {
			for i < j && data[j] >= pivotKey {
				j --
			}

			if i < j {
				data[i] = data[j]
				i ++
			}

			for i < j && data[i] <= pivotKey {
				i ++
			}

			if i < j {
				data[j] = data[i]
				j --
			}
		}

		// low ... i ... high
		// ....... pivotKey ......
		data[i] = pivotKey
		QuickSort(data, low, i -1)
		QuickSort(data, i + 1, high)
	}
}

func main() {
	var data []int = []int{3, 5, 0, 4, 9 , 8, 10, 41, 2, 14, 6, 1}
	QuickSort(data, 0, len(data) - 1)
	fmt.Println(data)
}
