# Functions

## Arguments

Function receives a **copy** of an argument.

**To modify** argument, function needs to receive a **pointer** to the value as an argument

## Variadic functions

`...` allows function to receive any number (variadic) of arguments of the same type

They are captured with slice and become `[]T`

```go
func variadicFunction(args ...int) {
    fmt.println(len(args))
}
```

We can also spread a slice when suffixing `...`

```go
nums := []int{1, 2, 3}

// Basically passing multiple arguments created with slice...
variadicFunction(nums...) // 3
```

## Pointers

```go
package main

import (
	"fmt"
)

func learnPointer(ptr *int) {
	fmt.Println(ptr)  // is passed address
	fmt.Println(*ptr) // is value at address
    fmt.Println(&ptr) // is address of actual pointer

	var ptrToPtr **int = &ptr

	fmt.Println(ptrToPtr)   // is address of the variable 'ptr'
	fmt.Println(&ptrToPtr)  // is address of the variable 'ptrToPtr'
    fmt.Println(*ptrToPtr)  // is same as 'ptr'
    fmt.Println(**ptrToPtr) // is same as '*ptr'
}

func main() {
   var x int = 1
   learnPointer(&x) // passing address
}

// Output
// 0x1040e0f8
// 1
// 0x1040a120
// 0x1040a120
// 0x1040a130
// 0x1040e0f8
// 1
```

## Basic

```go
package main

import "fmt"

// Type comes after the variable name
func add(x, y int) int { // Same as (x int, y int)
	return x + y
}

func main() { // main is special, used as an entry point for the main module
	fmt.Println(add(42, 13))
}

func init() { // will still be run before main func
}
```

## Multiple results

```go
package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
    fmt.Prinln(swap("can even", "handle this"))
}
```

## Named return values

```go
func split(sum int) (x, y int) {
    x = sum * 4 / 9
	y = sum - x
	return
}
```
