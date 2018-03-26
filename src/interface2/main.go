package main

import "fmt"

type People interface {
	Show()
}

type Student struct {

}

func (m *Student) Show()  {
	fmt.Println("Hello, world.")
}

func live() People  {
	var stu *Student
	return stu
}

func main()  {
	var people People = live()
	if people == nil {
		fmt.Println("A")
	} else {
		fmt.Println("B")
	}

	people.Show()
}
