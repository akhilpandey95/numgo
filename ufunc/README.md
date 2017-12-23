# numgo/ufunc
This segment of the library deals with having methods in place
that would be helpful to perform various mathematical and string
operations on the N-D Array. Math, String and Comparision operations
could be performed on the N-D Array using the below mentioned methods.

# Methods
Math, String and comparision operations are part of this segment.
Below are the list of methods available:

#### Math:
   - numgo.**Sin()**
   - numgo.**Cos()**
   - numgo.**Tan()**
   - numgo.**Cosec()**
   - numgo.**Sec()**
   - numgo.**Cot()**
   - numgo.**Asin()**
   - numgo.**Acos()**
   - numgo.**Atan()**
   - numgo.**Sinh()**
   - numgo.**Cosh()**
   - numgo.**Tanh()**
   - numgo.**Asinh()**
   - numgo.**Acosh()**
   - numgo.**Atanh()**
   - numgo.**Sqrt()**

#### Strings
   - numgo.**Strcount()**
   - numgo.**Strjoin()**
   - numgo.**Strtoupper()**
   - numgo.**Strtolower()**
   - numgo.**Strcapitalize()**
   - numgo.**Strcontains()**

#### Comparision of two N-D Arrays
   - numgo.**Equal()**
   - numgo.**Not_Equal()**
   - numgo.**Greater()**
   - numgo.**Greater_B()**
   - numgo.**Lesser()**
   - numgo.**Lesser_B()**

# Usage
numgo.**Sin()**

The `Sin()` function takes two arguments
`@param1 - string {measure}     either "deg" or "rad"`
`@param2 - NArray {n-d array}`
`*returns* - NArray {n-d array}`

```go
package main

import (
    "fmt"
    "github.com/akhilpandey95/numgo"
)

func main() {
    /*
     * d or deg implies the output is required in degrees
     * r or rad implies the output is required in radians
     */
    n := numgo.Xrange(1, 10)
    fmt.Println(numgo.Sin("d", n))
    fmt.Println(numgo.Sin("r", n))
    fmt.Println(numgo.Sin("deg", n))
    fmt.Println(numgo.Sin("rad", n))
}
```

numgo.**Cos()**

The `Cos()` function takes two arguments
`@param1 - string {measure}     either "deg" or "rad"`
`@param2 - NArray {n-d array}`
`returns - NArray {n-d array}`

```go
package main

import (
    "fmt"
    "github.com/akhilpandey95/numgo"
)

func main() {
    /*
     * d or deg implies the output is required in degrees
     * r or rad implies the output is required in radians
     */
    n := numgo.Xrange(1, 10)
    fmt.Println(numgo.Cos("d", n))
    fmt.Println(numgo.Cos("r", n))
    fmt.Println(numgo.Cos("deg", n))
    fmt.Println(numgo.Cos("rad", n))
}
```

numgo.**Tan()**

The `Tan()` function takes two arguments
`@param1 - string {measure}     either "deg" or "rad"`
`@param2 - NArray {n-d array}`
`returns - NArray {n-d array}`

```go
package main

import (
    "fmt"
    "github.com/akhilpandey95/numgo"
)

func main() {
    /*
     * d or deg implies the output is required in degrees
     * r or rad implies the output is required in radians
     */
    n := numgo.Xrange(1, 10)
    fmt.Println(numgo.Tan("d", n))
    fmt.Println(numgo.Tan("r", n))
    fmt.Println(numgo.Tan("deg", n))
    fmt.Println(numgo.Tan("rad", n))
}
```

numgo.**Cosec()**

The `Cosec()` function takes two arguments
`@param1 - string {measure}     either "deg" or "rad"`
`@param2 - NArray {n-d array}`
`returns - NArray {n-d array}`

```go
package main

import (
    "fmt"
    "github.com/akhilpandey95/numgo"
)

func main() {
    /*
     * d or deg implies the output is required in degrees
     * r or rad implies the output is required in radians
     */
    n := numgo.Xrange(1, 10)
    fmt.Println(numgo.Cosec("d", n))
    fmt.Println(numgo.Cosec("r", n))
    fmt.Println(numgo.Cosec("deg", n))
    fmt.Println(numgo.Cosec("rad", n))
}
```

numgo.**Sec()**

The `Sec()` function takes two arguments
`@param1 - string {measure}     either "deg" or "rad"`
`@param2 - NArray {n-d array}`
`returns - NArray {n-d array}`

```go
package main

import (
    "fmt"
    "github.com/akhilpandey95/numgo"
)

func main() {
    /*
     * d or deg implies the output is required in degrees
     * r or rad implies the output is required in radians
     */
    n := numgo.Xrange(1, 10)
    fmt.Println(numgo.Sec("d", n))
    fmt.Println(numgo.Sec("r", n))
    fmt.Println(numgo.Sec("deg", n))
    fmt.Println(numgo.Sec("rad", n))
}
```

numgo.**Cot()**

The `Cot()` function takes two arguments
`@param1 - string {measure}     either "deg" or "rad"`
`@param2 - NArray {n-d array}`
`returns - NArray {n-d array}`

```go
package main

import (
    "fmt"
    "github.com/akhilpandey95/numgo"
)

func main() {
    /*
     * d or deg implies the output is required in degrees
     * r or rad implies the output is required in radians
     */
    n := numgo.Xrange(1, 10)
    fmt.Println(numgo.Cot("d", n))
    fmt.Println(numgo.Cot("r", n))
    fmt.Println(numgo.Cot("deg", n))
    fmt.Println(numgo.Cot("rad", n))
}
```

numgo.**Asin()**

The `Asin()` function takes one arguments
`@param1 - NArray {n-d array}`
`returns - NArray {n-d array}`

```go
package main

import (
    "fmt"
    "github.com/akhilpandey95/numgo"
)

func main() {
    n := numgo.Xrange(1, 10)
    fmt.Println(numgo.Asin(n))
}
```

numgo.**Acos()**

The `Acos()` function takes one arguments
`@param1 - NArray {n-d array}`
`returns - NArray {n-d array}`

```go
package main

import (
    "fmt"
    "github.com/akhilpandey95/numgo"
)

func main() {
    n := numgo.Xrange(1, 10)
    fmt.Println(numgo.Acos(n))
}
```

numgo.**Atan()**

The `Atan()` function takes one arguments
`@param1 - NArray {n-d array}`
`returns - NArray {n-d array}`

```go
package main

import (
    "fmt"
    "github.com/akhilpandey95/numgo"
)

func main() {
    n := numgo.Xrange(1, 10)
    fmt.Println(numgo.Atan(n))
}
```

numgo.**Sinh()**

The `Sinh()` function takes one arguments
`@param1 - NArray {n-d array}`
`returns - NArray {n-d array}`

```go
package main

import (
    "fmt"
    "github.com/akhilpandey95/numgo"
)

func main() {
    n := numgo.Xrange(1, 10)
    fmt.Println(numgo.Sinh(n))
}
```

numgo.**Cosh()**

The `Cosh()` function takes one arguments
`@param1 - NArray {n-d array}`
`returns - NArray {n-d array}`

```go
package main

import (
    "fmt"
    "github.com/akhilpandey95/numgo"
)

func main() {
    n := numgo.Xrange(1, 10)
    fmt.Println(numgo.Cosh(n))
}
```

numgo.**Tanh()**

The `Tanh()` function takes one arguments
`@param1 - NArray {n-d array}`
`returns - NArray {n-d array}`

```go
package main

import (
    "fmt"
    "github.com/akhilpandey95/numgo"
)

func main() {
    n := numgo.Xrange(1, 10)
    fmt.Println(numgo.Tanh(n))
}
```

numgo.**Asinh()**

The `Asinh()` function takes one arguments
`@param1 - NArray {n-d array}`
`returns - NArray {n-d array}`

```go
package main

import (
    "fmt"
    "github.com/akhilpandey95/numgo"
)

func main() {
    n := numgo.Xrange(1, 10)
    fmt.Println(numgo.Asinh(n))
}
```

numgo.**Acosh()**

The `Acosh()` function takes one arguments
`@param1 - NArray {n-d array}`
`returns - NArray {n-d array}`

```go
package main

import (
    "fmt"
    "github.com/akhilpandey95/numgo"
)

func main() {
    n := numgo.Xrange(1, 10)
    fmt.Println(numgo.Acosh(n))
}
```

numgo.**Atanh()**

The `Atanh()` function takes one arguments
`@param1 - NArray {n-d array}`
`returns - NArray {n-d array}`

```go
package main

import (
    "fmt"
    "github.com/akhilpandey95/numgo"
)

func main() {
    n := numgo.Xrange(1, 10)
    fmt.Println(numgo.Atanh(n))
}
```

numgo.**Sqrt()**

The `Sqrt()` function takes one arguments
`@param1 - NArray {n-d array}`
`returns - NArray {n-d array}`

```go
package main

import (
    "fmt"
    "github.com/akhilpandey95/numgo"
)

func main() {
    n := numgo.Xrange(10, 100, 10)
    fmt.Println(numgo.Sqrt(n))
}
```

numgo.**Strcount()**
```go
package main

import (
    "fmt"
    "github.com/akhilpandey95/numgo"
)

func main() {
    fmt.Println("This is just the start")
}
```

numgo.**Strjoin()**
```go
package main

import (
    "fmt"
    "github.com/akhilpandey95/numgo"
)

func main() {
    fmt.Println("This is just the start")
}
```

numgo.**Strtoupper()**
```go
package main

import (
    "fmt"
    "github.com/akhilpandey95/numgo"
)

func main() {
    fmt.Println("This is just the start")
}
```

numgo.**Strtolower()**
```go
package main

import (
    "fmt"
    "github.com/akhilpandey95/numgo"
)

func main() {
    fmt.Println("This is just the start")
}
```

numgo.**Strcapitalize()**
```go
package main

import (
    "fmt"
    "github.com/akhilpandey95/numgo"
)

func main() {
    fmt.Println("This is just the start")
}
```

numgo.**Strcontains()**
```go
package main

import (
    "fmt"
    "github.com/akhilpandey95/numgo"
)

func main() {
    fmt.Println("This is just the start")
}
```

numgo.**Equal()**
```go
package main

import (
    "fmt"
    "github.com/akhilpandey95/numgo"
)

func main() {
    fmt.Println("This is just the start")
}
```

numgo.**Not_Equal()**
```go
package main

import (
    "fmt"
    "github.com/akhilpandey95/numgo"
)

func main() {
    fmt.Println("This is just the start")
}
```

numgo.**Greater()**
```go
package main

import (
    "fmt"
    "github.com/akhilpandey95/numgo"
)

func main() {
    fmt.Println("This is just the start")
}
```

numgo.**Greater_B()**
```go
package main

import (
    "fmt"
    "github.com/akhilpandey95/numgo"
)

func main() {
    fmt.Println("This is just the start")
}
```

numgo.**Lesser()**
```go
package main

import (
    "fmt"
    "github.com/akhilpandey95/numgo"
)

func main() {
    fmt.Println("This is just the start")
}
```

numgo.**Lesser_B()**
```go
package main

import (
    "fmt"
    "github.com/akhilpandey95/numgo"
)

func main() {
    fmt.Println("This is just the start")
}
```

# Support
- go version go1.9 linux/amd64

# License
[The MIT License](https://github.com/akhilpandey95/numGo/blob/master/LICENSE)

# Contributors
Would love to see anyone willing to contribute

# Maintainer
[Akhil Pandey](https://github.com/akhilpandey95)
