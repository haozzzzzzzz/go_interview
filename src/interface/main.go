package main

import "fmt"

type People interface {
	Speak(string)string
}

type Student struct {

}

func (m *Student) Speak(think string) (talk string)  {
	if think == "bitch" {
		talk = "You are a good boy."
	} else {
		talk = "hi"
	}
	return
}

func main() {
	var people People = &Student{}
	think := "bitch"
	fmt.Print(people.Speak(think))
}