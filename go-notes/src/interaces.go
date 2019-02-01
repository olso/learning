package main

import (
    "fmt"
    "math"
)

type Shape interface {
    area() float64
}

type Something interface {
    area() float64
}

type Rectangle struct {
    height float64
    width float64
}

type Circle struct {
    radius float64
}

func main() {
    fmt.Println()

    rectangle:=Rectangle{height:20, width:50}
    circle:=Circle{radius:25}

    fmt.Println(rectangle, circle)
    fmt.Println(rectangle.area(), circle.area())
}

func (rectangle *Rectangle) area() float64 {
    return rectangle.height * rectangle.width
}

func (circle *Circle) area() float64 {
    return math.Pi * math.Pow(circle.radius, 2)
}

func getArea(shape Shape) float64 {
    return shape.area()
}

func getAre(something Something) float64 {
    return something.area()
}
