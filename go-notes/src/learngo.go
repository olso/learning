package main

import (
	"fmt" // std go lib
	//"io/ioutil"
	//m "math" // math lib with m alias
	//"net/http" // networking
	//"os" // os lib, working with file system
	//"strconv" // string conversions lib
)

var packageScopeVariable bool = true

func main() {
    //fmt.Println("MAIN")
    //basicArgument()
    //pointerArgument()
    shorthandVariable()
}

func basicArgument() {
    fmt.Println()
}

func pointerArgument() {
    var x int = 1337
    learnPointer(&x)
    fmt.Println(x)
}

func learnPointer(ptr *int) {
    fmt.Println(ptr)  // is passed address
    fmt.Println(*ptr) // is value at address
    fmt.Println(&ptr) // is address of actual pointer

    *ptr += 1

    var ptrToPtr **int = &ptr

    fmt.Println(ptrToPtr)   // is address of the variable 'ptr'
    fmt.Println(&ptrToPtr)  // is address of the variable 'ptrToPtr'
    fmt.Println(*ptrToPtr)  // is same as 'ptr'
    fmt.Println(**ptrToPtr) // is same as '*ptr'
}

func swap(a, b string) (string, string) {
    return b, a
}

func shorthandVariable() {
    var i, j int = 1, 2
    k := 3
    c, python, java := true, false, "no!"
    fmt.Println(i, j, k, c, python, java)
}
