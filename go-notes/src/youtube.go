package main

import (
    "fmt"
    "math/rand"
)

//someNumbers := []float64 {rand.Float64(), rand.Float64()}

func main() {
    fmt.Println()
    //forLoop0()
    //forLoop1()
    //arrays()

    //someNumbers := []float64 {rand.Float64(), rand.Float64()}
    //fmt.Println(addUpArrayOfNumbers(&someNumbers))
    //fmt.Println(substractNumbers(someNumbers...))
    //fmt.Println(closureSum(&someNumbers)())
    deferExample()
}

func forLoop0() {
    i := 1

    for i <= 10 {
        fmt.Println(i);
        i++
    }
}

func forLoop1() {
    for j := 1; j <= 20; j++ {
        fmt.Println(j)
    }
}

func arrays() {
    var numbers[5] float64

    otherNumbers := [5]float32 {1,2,3,4,5}

    for i := 0; i <= 4; i++ {
        numbers[i] = 312312 * rand.Float64()

        fmt.Println(numbers[i], otherNumbers[i])
    }
}

func addUpArrayOfNumbers(numbers *[]float64) (int, float64) {
    var sum float64 = 0

    for _, number := range *numbers {
        sum += number
    }

    return len(*numbers), sum
}

func substractNumbers(numbers ...float64) float64{
    var result float64 = 0

    for _, number := range numbers {
        result -= number
    }

    return result
}

func closureSum(numbers *[]float64) func() float64  {
    return func () float64 {
        return substractNumbers(*numbers...)
    }
}

func deferExample() {
    defer fmt.Println("first in code order, and will run after parent function is finished")
    fmt.Println("second in code, but will execute first")
}
