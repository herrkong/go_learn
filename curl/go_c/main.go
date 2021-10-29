package main

import "C"
import (
	"unsafe"
)
 
// #include <stdio.h>
// #include <stdlib.h>

// void print(char *str) {
//     printf("%s\n", str);
// }


func main() {
    s := "Hello Cgo"
    cs := C.CString(s)
    C.print(cs)
    C.free(unsafe.Pointer(cs))
}