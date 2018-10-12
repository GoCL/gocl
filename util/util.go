package util

/*
#include "util.h"
*/
import "C"
import "fmt"

func GoSum(a int, b int) {
	s := C.sum(C.int(a), C.int(b))
	fmt.Print(s)
}
