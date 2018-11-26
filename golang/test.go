package main
 
/*
#cgo CFLAGS: -I./
#cgo LDFLAGS: -L./ -ltest
#include "calculator.h" //非标准c头文件，所以用引号
*/
import "C"
 
/*
CFLAGS中的-I（大写的i） 参数表示.h头文件所在的路径
LDFLAGS中的-L(大写) 表示.so文件所在的路径 -l(小写的L) 
表示指定该路径下的库名称，比如要使用libtest.so，则只需用-ltest （省略了libtest.so中的lib和.so字符，关于这些字符所代表的具体含义请自行google）表示。
*/
import (
    "fmt"
)
 
func main() {
	result1 := C.add(77,88)
	result2 := C.minus(100,22)
    fmt.Printf("result1:%d,result2:%d\n",result1,result2)
}