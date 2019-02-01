package main

import "fmt"

type Rectangle struct {
    leftX float64
    topY float64
    height float64
    width float64
}

func main() {
    rect1:=Rectangle{leftX:0, topY:50, height:10, width:20}

    fmt.Println(rect1, rect1.area(), area(rect1))
}

func (rect *Rectangle) area() float64 {
    return rect.width * rect.height
}

func area(rect Rectangle) float64 {
    return rect.height * rect.width
}
