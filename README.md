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
