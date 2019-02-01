# Modules

## Package

Package names are lowercase

Source files share package scope, they can see each other's globals, even if they are not exported (are lowercase)

```go
package main // is special entry point for our program

/*
http://stackoverflow.com/a/24790378

is guaranteed to run before main
you can have multiple init functions per package
they will be executed in order they show up in code
*/
func init () {
}

func main () {
}
```

## Import

### Simple import

```go
package main

import "fmt"
import "math/rand"
```

### Shorthand import (preferred)

```go
package main

import (
    "fmt"
    "math/rand"
)
```

### Aliased import

```go
package main

import (
    f "fmt"
    rand "math/rand"
)
```

## Export

In Go, a name is exported if it begins with capital letter

```go
package main

import (
    "fmt"
    "math"
)

func main() {
    fmt.Println(math.Pi); //   capital Pi will work
    fmt.Println(math.pi); // lowercase pi won't work
}
```
