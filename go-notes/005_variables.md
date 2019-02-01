# Variables

`var` declares a list of variables, variable type is defined after variable name

`var` statement can can be at package scope or at function scope

```go
package main

var packageScopeVariable bool = true

func main() {
   var functionScopeVariable bool = true
}
```

## Variables with initializers

```go
package main

import "fmt"

var i, j int = 1, 2

func main() {
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
}
```

## Short variable declarations

In function, we can use `:=` as a shorthand for `var` declaration with implicit type

```go
func main() {
    var i, j int = 1, 2 // must be same type
    k := 3
    c, python, java := true, false, "no!" // doesn't have to be same type
}
```
