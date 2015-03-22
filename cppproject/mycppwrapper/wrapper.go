package mycppwrapper

/*
#include "cpplib.h"
*/
import "C"

func Bar() {
    C.bar()
}
