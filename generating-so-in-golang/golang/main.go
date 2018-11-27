package main

 

import "C"

//export Add
func Add(a int,b int) int {
       return a+b
}

//export Minus
func Minus(a int,b int) int {
	return a-b
}

func main() {}

// func main() {
// 	result1 := Add(77,88)
// 	result2 := Minus(100,22)
//     fmt.Printf("result1:%d,result2:%d\n",result1,result2)
// }