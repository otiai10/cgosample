package main

/*
#include <iostream>
void baz() {
    std::count << "HELLO, C++!" << std::endl;
}
*/

/*
#include <stdio.h>
void baz() {
    printf("HELLO, C!!!!!!!!!!!!!!!!!!!\n");
}
*/
import "C"

func main() {
    C.baz()
}
