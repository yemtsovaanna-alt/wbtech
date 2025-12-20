package main

import "fmt"

type Human struct {
    name string
    age  int
}

func (h Human) Walk() {
    fmt.Println(h.name, "is walking")
}

func (h Human) Speak() {
    fmt.Println(h.name, "is speaking")
}

func (h Human) GetAge() int {
    return h.age
}

type Action struct {
    Human
    actionType string
}

func main() {
    action := Action{
        Human: Human{
            name: "John",
            age:  30,
        },
        actionType: "running",
    }

    action.Walk()
    action.Speak()
    fmt.Println("Age:", action.GetAge())
}
