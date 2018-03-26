package main

import "fmt"

const (
	x = iota
	y
	z = "zz"
	k
	p = iota
	q
	o = "2222"
	a
	b
)

const (
	h = iota
	i
)
func main()  {
	fmt.Println(x, y, z, k, p, q)
	fmt.Println(o, a, b)
	fmt.Println(h, i)
}
