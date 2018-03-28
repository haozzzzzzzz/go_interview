package main

import "fmt"

func main() {
	i := 0

	defer func() {
		i ++
		fmt.Println(i)
	}()


	fmt.Println(i)


	defer func(i int) {
		i ++
		fmt.Println(i)
	}(i)

	fmt.Println(i)
}
