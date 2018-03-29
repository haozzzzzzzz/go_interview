package main

import "fmt"

func change(m map[string]int, key string, value int) {
	m[key] = value
}

func main() {
	m := make(map[string]int, 0)
	change(m, "pp", 1)
	fmt.Println(m)
}
