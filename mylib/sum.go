package mylib

/*
int mysum(int i, int j) {
  return i + j;
}
*/
import "C"

// Sum is proxy func for C.mysum.
func Sum(i, j int) int {
	res, err := C.mysum(C.int(i), C.int(j))
	if err != nil {
		panic(err)
	}
	return int(res)
}
