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
