# cgosample

Sample to write cgo.

# 1) How to start "cgo"

[main.go](main.go)

```go
package main

/*
#include <stdio.h>
void foo(char* s) {
    printf("foo received: %s\n", s);
}
*/
import "C"

func main() {
    message := C.CString("Hello, cgo!!")
    C.foo(message)
}
```

and

```
go run main.go
```

# 2) How to make wrapper package and tests

[mylib/sum.go](mylib/sum.go)

```
./mylib
├── sum.go
└── test
    └── sum_test.go
```

```go
package mylibtest

// To avoid import cycle, this package should be named {yourpackage}test.

import (
    "testing"

    "github.com/otiai10/cgosample/mylib"
)

func TestSum(t *testing.T) {
    if mylib.Sum(1, 3) != 4 {
        t.Fail()
    }
    if mylib.Sum(2, 4) != 6 {
        t.Fail()
    }
}
```

run the test

```
go test ./mylib/test/... -v
```
