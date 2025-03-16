package main

import (
	"errors"
	"fmt"
)

func main() {

	// fmt.Println("StartHere")
	// ErrBadIdea := errors.New("Something Bad")
	// fmt.Printf("ErrBadData Type %T", ErrBadIdea)

	msg := "goodbyee"
	message(msg)

	c := cat{name: "Oreo", age: 9}
	fmt.Println(c.Speak())
	CatSpeaker(c)
}

func message(s string) {
	if s == "goodbye" {
		panic(errors.New("Accident:" + s))
	}
}

type Speaker interface { // 其他型別只要實現了完全相同的方法，就等於是實作了此介面，不需要宣告implements
	Speak() string // 慣例命名會使用當中的一個方法加上er
}

type cat struct {
	name string
	age  int
}

func (c cat) Speak() string {
	return "meowww"
}
func (c cat) String() string {
	return fmt.Sprintf("%v (%v years old)", c.name, c.age)
}

func CatSpeaker(s Speaker) {
	fmt.Println(s.Speak())
}
