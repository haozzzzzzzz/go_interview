package main

import "fmt"

type Phone interface {
	Call()
}

type IPhone struct {
}

func (m *IPhone) Call()  {
	fmt.Println("Hello")
}

func main()  {
	var phone Phone
	phone = &IPhone{}
	phone.Call()

	iPhone, ok := phone.(Phone)
	fmt.Println(ok, iPhone)
}
