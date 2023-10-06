package main

import "fmt"

type Human struct {
	Age  int
	Name string
}

type Action struct {
	Human // embedded
}

func main() {
	human := &Human{
		Age:  13,
		Name: "Stas",
	}
	human.Hi()
	act := &Action{*human}

	act.Hi()
}

func (h *Human) Hi() {
	fmt.Println("Hello my name is", h.Name)
}

func (h *Human) Bye() {
	fmt.Println("Bye")
}
