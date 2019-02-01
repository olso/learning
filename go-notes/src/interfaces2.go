package main

import (
    "fmt"
)

type Animal interface {
    Speak() string
}

type Dog struct {
}

type Cat struct {
}

type Human struct {
}

func main() {
    animals:=[]Animal{Dog{},Cat{},Human{}}
    fmt.Println(animals)

    cats:=[]Cat{Cat{},Cat{},Cat{}}
    fmt.Println(cats)

    for _, animal := range animals {
        fmt.Println(animal.Speak())
    }

    for _, cat := range cats {
        fmt.Println(cat.Speak())
    }

    fmt.Println()
}

func (cat Cat) Speak() string {
    return "Meow!"
}

func (dog Dog) Speak() string {
    return "Woof!"
}

func (human Human) Speak() string {
    return "Yolo!"
}
