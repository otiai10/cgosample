#include <iostream>
extern "C" {
void bar() {
    std::cout << "Hello, cgo with C++!" << std::endl;
}
}/* extern "C" */
