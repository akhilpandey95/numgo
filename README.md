# numgo
A library that is aimed at providing methods and functions that
give scope for performing numerical and scientific calculations
in `go-lang`. The core concept to `numgo` library is that everything
is an array and all operations, calculations, functions and routines
that correspond to are in reference to that `n-dimensional` array.

# Inspiration
Inspired by `numpy` for `python`, `numGo` aims to replicate all the
mathematical concepts that are covered in `numpy`. Therefore, most
of the concepts and ideas are broadly taken from `numpy`.

# Operations to be supported
- Algebra
- Arithmatics
- Trigonometry
- Exponents and logarithms
- Universal Functions
- Logical Functions

# Installation
```shell
go get github.com/akhilpandey95/numgo
```

# Usage
```go
package main

import (
    "fmt"
    "github.com/akhilpandey95/numgo"
)

func main() {
    fmt.Println(numgo.NDArray(3, 3, 3))  // would init a NDArray of shape 3,3,3
    fmt.Println(numgo.Xrange(8, 20))     // would print from 8 to 20
    fmt.Println(numgo.Xrange(8, 9, 0.1)) // would print from 8.1 to 8.9
}
```

# Methods
- [linalg/vectors](https://github.com/akhilpandey95/numGo/blob/master/linalg/README.md)         - Vectors, matrices and eigen values
- [ufunc/math](https://github.com/akhilpandey95/numGo/blob/master/ufunc/README.md)             - Trignometric and mathematical operations on NDArrays
- [ufunc/strings](https://github.com/akhilpandey95/numGo/blob/master/ufunc/README.md)        - Operations supporting string operations
- [ufunc/comparision](https://github.com/akhilpandey95/numGo/blob/master/ufunc/README.md)      - Comparing two NDArrays as to which is Greater, Smaller etc
- [routines/random](https://github.com/akhilpandey95/numGo/blob/master/routines/README.md)        - Generation of Random Numbers on NDArrays and operations on them
- [routines/sampling](https://github.com/akhilpandey95/numGo/blob/master/routines/README.md)      - Sampling array data of type NDArray
- [routines/distributions](https://github.com/akhilpandey95/numGo/blob/master/routines/README.md) - Various Mathematical Distributions on NDArray data

`NOTE:` The package in its entirety is under development `v0.1` will be shipped by the end of December 2017.

# Support
- go version go1.9 linux/amd64

# License
[The MIT License](https://github.com/akhilpandey95/numGo/blob/master/LICENSE)

# Contributors
Would love to see anyone willing to contribute

# Maintainer
[Akhil Pandey](https://github.com/akhilpandey95)
